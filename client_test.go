package tuyasdk_test

import (
	"testing"
	"time"

	tuyasdk "github.com/iot-eco-system/tuya-iot-service-sdk"
	"github.com/iot-eco-system/tuya-iot-service-sdk/model"
)

const (
	clientID = "<CLIENT_ID>"
	secret   = "<SECRET>"
)

func assert(t *testing.T, cond bool, msg string, args ...interface{}) {
	if !cond {
		t.Fatalf(msg, args...)
	}
}

func TestClientStart(t *testing.T) {
	cli := tuyasdk.NewTuyaAPIClient(tuyasdk.NewTuyaAPIClientOptions{
		ClientID: clientID,
		Secret:   secret,
	})

	cli.Start()

	timer := time.NewTicker(15 * time.Second)
	<-timer.C

	assert(t, cli.Ready(), "get access token failed")
}

func TestClientAPI(t *testing.T) {
	cli := tuyasdk.NewTuyaAPIClient(tuyasdk.NewTuyaAPIClientOptions{
		ClientID: clientID,
		Secret:   secret,
	})

	cli.Start()

	timer := time.NewTicker(15 * time.Second)
	<-timer.C

	assert(t, cli.Ready(), "get access token failed")

	t.Run("call user list api", func(t *testing.T) {
		resp, err := cli.UserInfoList(&model.UserInfoListRequest{
			Schema:   "iotsdktestapp",
			PageNo:   1,
			PageSize: 10,
		})
		assert(t, err == nil, "failed to get user list")
		assert(t, resp.Result.Total == 0, "total invalid")
	})

	t.Run("call device list api", func(t *testing.T) {
		did := "<DEVICE_ID>"
		resp, err := cli.DeviceGetInfo(&model.DeviceGetInfoRequest{
			DeviceID: did,
		})
		assert(t, err == nil, "failed to get device info")
		assert(t, resp.Result.ID == did, "failed to get device info")
	})
}

func TestBuildQueryURL(t *testing.T) {
	q := model.DeviceListRequest{
		SourceType: "1",
		SourceID:   "2",
		DeviceIDs:  "3",
		Name:       "4",
		Category:   "5",
		ProductID:  "6",
		LastRowKey: "7",
		PageSize:   1,
	}

	expect := "/v1/devices?category=5&device_ids=3&last_row_key=7&name=4&page_size=1&product_id=6&source_id=2&source_type=1"

	endpoint := tuyasdk.BuildQueryURL("/v1/devices", q)
	assert(t, endpoint == expect, endpoint)

	endpoint = tuyasdk.BuildQueryURL("/v1/devices", &q)
	assert(t, endpoint == expect, endpoint)
}
