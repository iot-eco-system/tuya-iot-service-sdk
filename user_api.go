package tuyasdk

import (
	"fmt"
	"net/http"

	"github.com/iot-eco-system/tuya-iot-service-sdk/model"
)

// 同步用户
// https://developer.tuya.com/cn/docs/cloud/21707ff1ba?id=Kawfjd7120rb0
func (t *TuyaAPIClient) UserSync(request *model.UserSyncRequest) (*model.UserSyncResponse, error) {
	var response model.UserSyncResponse
	endpoint := fmt.Sprintf(EndpointUserSync, request.Schema)
	if err := t.InvokeAPI(http.MethodPost, endpoint, request, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// 获取用户信息
// https://developer.tuya.com/cn/docs/cloud/cfebf22ad3?id=Kawfjdgic5c0w
func (t *TuyaAPIClient) UserInfoGet(request *model.UserInfoGetRequest) (*model.UserInfoGetResponse, error) {
	var response model.UserInfoGetResponse
	endpoint := fmt.Sprintf(EndpointUserInfoGet, request.UID)
	if err := t.InvokeAPI(http.MethodGet, endpoint, nil, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// 获取用户列表
// https://developer.tuya.com/cn/docs/cloud/76f3e0885f?id=Kawfji9n0sdmq
func (t *TuyaAPIClient) UserInfoList(request *model.UserInfoListRequest) (*model.UserInfoListResponse, error) {
	var response model.UserInfoListResponse
	endpoint := fmt.Sprintf(EndpointUserList, request.Schema)
	endpoint = BuildQueryURL(endpoint, request)

	if err := t.InvokeAPI(http.MethodGet, endpoint, nil, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// 关联用户授权登录
// https://developer.tuya.com/cn/docs/cloud/c40fc05907?id=Kawfjj0r2m82l
func (t *TuyaAPIClient) UserAuthorizedLogin(request *model.UserAuthorizedLoginRequest) (*model.UserAuthorizedLoginResponse, error) {
	var response model.UserAuthorizedLoginResponse
	if err := t.InvokeAPI(http.MethodPost, EndpointUserAuthorizedLogin, request, &response); err != nil {
		return nil, err
	}
	return &response, nil
}
