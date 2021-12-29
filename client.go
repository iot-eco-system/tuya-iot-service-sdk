package tuyasdk

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"runtime/debug"
	"sync"
	"time"

	"github.com/iot-eco-system/tuya-iot-service-sdk/model"
	"go.uber.org/zap"
)

const (
	DefaultTuyaAPIHost = "https://openapi.tuyacn.com"
	RetryWait          = 5 * time.Second
)

type TuyaAPIClient struct {
	Host     string
	ClientID string
	Secret   string

	// API token
	tokenMU      sync.RWMutex
	token        string
	refreshToken string
	tokenExpires int

	logger *zap.Logger

	httpClient *http.Client
	stop       chan struct{}
}

type NewTuyaAPIClientOptions struct {
	Debug    bool
	Host     string
	ClientID string
	Secret   string
}

func NewTuyaAPIClient(opts NewTuyaAPIClientOptions) *TuyaAPIClient {
	if len(opts.ClientID) == 0 {
		panic("ClientID should not be empty")
	}
	if len(opts.Secret) == 0 {
		panic("Secret should not be empty")
	}

	t := &TuyaAPIClient{
		Host:       opts.Host,
		ClientID:   opts.ClientID,
		Secret:     opts.Secret,
		httpClient: http.DefaultClient,
		stop:       make(chan struct{}),
	}

	var loggerConfig zap.Config
	if opts.Debug {
		loggerConfig = zap.NewDevelopmentConfig()
	} else {
		loggerConfig = zap.NewProductionConfig()
	}

	loggerConfig.EncoderConfig.TimeKey = "timestamp"
	t.logger, _ = loggerConfig.Build()

	if len(t.Host) == 0 {
		t.Host = DefaultTuyaAPIHost
	}

	return t
}

// Check if client is ready for accepting request or not
func (t *TuyaAPIClient) Ready() bool {
	t.tokenMU.RLock()
	defer t.tokenMU.RUnlock()

	return len(t.token) > 0 && len(t.refreshToken) > 0
}

// Invoke API
func (t *TuyaAPIClient) InvokeAPI(method, endpoint string, request interface{}, response model.Response) error {
	body := []byte(``)
	if request != nil {
		body, _ = json.Marshal(request)
	}

	t.logger.Debug(
		"invoke api:",
		zap.ByteString("body", body),
		zap.String("endpoint", endpoint),
	)

	req, _ := http.NewRequest(method, t.Host+endpoint, bytes.NewReader(body))
	t.buildHeader(req, body)
	resp, err := t.httpClient.Do(req)
	if err != nil {
		t.logger.Error("failed to invoke api",
			zap.String("endpoint", endpoint),
			zap.Error(err),
		)
		return err
	}

	defer resp.Body.Close()
	bs, _ := ioutil.ReadAll(resp.Body)
	t.logger.Debug("api response:", zap.ByteString("body", bs))
	err = json.Unmarshal(bs, response)
	if err != nil {
		t.logger.Error("failed to parse response",
			zap.String("endpoint", endpoint),
			zap.Error(err),
		)
		return err
	}

	t.logger.Debug(
		"invoke api complete",
		zap.String("endpoint", endpoint),
		zap.ByteString("payload", bs),
	)

	if !response.IsSuccess() {
		return errors.New(response.GetErrorMessage())
	}

	return nil
}

func (t *TuyaAPIClient) Start() {
	defer func() {
		if err := recover(); err != nil {
			t.logger.Debug("error occurred")
			debug.PrintStack()
		}
	}()
	go t.start()
}

func (t *TuyaAPIClient) Stop() {
	t.logger.Sync()
	t.stop <- struct{}{}
}

func (t *TuyaAPIClient) start() {
	tick := time.NewTicker(RetryWait)

gettoken:
	for {
		select {
		case <-t.stop:
			tick.Stop()
			return
		case <-tick.C:
			if err := t.GetToken(); err == nil {
				tick.Stop()
				break gettoken
			}
		}
	}

	for {
		tick = time.NewTicker(time.Duration(t.tokenExpires) * time.Second)
		select {
		case <-tick.C:
			t.RefreshToken()
		case <-t.stop:
			tick.Stop()
			return
		}
	}
}

func BuildQueryURL(endpoint string, request interface{}) string {
	val := reflect.ValueOf(request)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		panic("request should be struct or pointer to struct")
	}

	query := make(url.Values)
	for i := 0; i < val.NumField(); i++ {
		fieldTyp := val.Type().Field(i)
		key := fieldTyp.Tag.Get("query")
		if key == "-" {
			continue
		}

		fieldVal := val.Field(i)
		query.Add(key, fmt.Sprintf("%v", fieldVal.Interface()))
	}

	return endpoint + "?" + query.Encode()
}
