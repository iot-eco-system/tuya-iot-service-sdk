package tuyasdk

import (
	"net/http"

	"github.com/iot-eco-system/tuya-iot-service-sdk/model"
	"go.uber.org/zap"
)

// 获取token
func (t *TuyaAPIClient) GetToken() error {
	var response model.TokenResponse
	if err := t.InvokeAPI(http.MethodGet, EndpointGetAccessToken, nil, &response); err != nil {
		t.logger.Error("failed to get access token", zap.Error(err))
		return err
	}

	t.tokenMU.Lock()
	defer t.tokenMU.Unlock()

	t.token = response.Result.AccessToken
	t.refreshToken = response.Result.RefreshToken
	t.tokenExpires = response.Result.ExpireTime
	return nil
}

// 刷新token
func (t *TuyaAPIClient) RefreshToken() error {
	t.tokenMU.RLock()
	if t.refreshToken == "" {
		return ErrRefreshTokenEmpty
	}
	var response model.TokenResponse
	if err := t.InvokeAPI(http.MethodGet, EndpointRefreshToken, nil, &response); err != nil {
		t.logger.Error("failed to refresh access token", zap.Error(err))
		return ErrGetAccessTokenFailed
	}

	t.tokenMU.Lock()
	defer t.tokenMU.Unlock()

	t.token = response.Result.AccessToken
	t.refreshToken = response.Result.RefreshToken
	t.tokenExpires = response.Result.ExpireTime
	return nil
}
