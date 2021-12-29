package tuyasdk

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/iot-eco-system/tuya-iot-service-sdk/model"
)

// 下发指令
// https://developer.tuya.com/cn/docs/cloud/e2512fb901?id=Kag2yag3tiqn5
func (t *TuyaAPIClient) DeviceSendCommands(request *model.DeviceSendCommandRequest) (*model.DeviceSendCommandResponse, error) {
	var response model.DeviceSendCommandResponse
	endpoint := fmt.Sprintf(EndpointDeviceSendCommand, request.DeviceID)
	if err := t.InvokeAPI(http.MethodPost, endpoint, request, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// 获取设备列表
// https://developer.tuya.com/cn/docs/cloud/0f866b1299?id=Kb3d6972ym5np
func (t *TuyaAPIClient) DeviceList(request *model.DeviceListRequest) (*model.DeviceListResponse, error) {
	var response model.DeviceListResponse
	if err := t.InvokeAPI(http.MethodGet, EndpointDeviceList, nil, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// 批量获取设备最新状态
// https://developer.tuya.com/cn/docs/cloud/8919fe076c?id=Kag2yc7dle32x
func (t *TuyaAPIClient) DeviceGetStatusMulti(request *model.DeviceGetStatusMultiRequest) (*model.DeviceGetStatusMultiResponse, error) {
	var response model.DeviceGetStatusMultiResponse
	endpoint := EndpointDeviceGetStatusMulti + "?device_ids=" + strings.Join(request.DeviceIDs, ",")
	if err := t.InvokeAPI(http.MethodGet, endpoint, nil, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// 获取单个设备的状态
// https://developer.tuya.com/cn/docs/cloud/f76865b055?id=Kag2ycn1lvwpt
func (t *TuyaAPIClient) DeviceGetStatusSingle(request *model.DeviceGetStatusSingleRequest) (*model.DeviceGetStatusSingleResponse, error) {
	var response model.DeviceGetStatusSingleResponse
	endpoint := fmt.Sprintf(EndpointDeviceGetStatusSingle, request.DeviceID)
	if err := t.InvokeAPI(http.MethodGet, endpoint, nil, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// 获取设备信息
// https://developer.tuya.com/cn/docs/cloud/d00d20c097?id=Kag2xtiyewd3r
func (t *TuyaAPIClient) DeviceGetInfo(request *model.DeviceGetInfoRequest) (*model.DeviceGetInfoResponse, error) {
	var response model.DeviceGetInfoResponse
	endpoint := fmt.Sprintf(EndpointDeviceGetInfo, request.DeviceID)
	if err := t.InvokeAPI(http.MethodGet, endpoint, nil, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// 修改设备信息
// https://developer.tuya.com/cn/docs/cloud/a5930a2eb8?id=Kag2y0jkpziab
func (t *TuyaAPIClient) DeviceUpdateInfo(request *model.DeviceUpdateInfoRequest) (*model.DeviceUpdateInfoResponse, error) {
	var response model.DeviceUpdateInfoResponse
	endpoint := fmt.Sprintf(EndpointDeviceUpdateInfo, request.DeviceID)
	if err := t.InvokeAPI(http.MethodPut, endpoint, request, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// 批量移除设备
// https://developer.tuya.com/cn/docs/cloud/9d95322343?id=Kag2y03rkutdv
func (t *TuyaAPIClient) DeviceRemoveMulti(request *model.DeviceRemoveMultiRequest) (*model.DeviceRemoveMultiResponse, error) {
	var response model.DeviceRemoveMultiResponse
	endpoint := fmt.Sprintf(EndpointDeviceRemoveMulti, strings.Join(request.DeviceIDs, ","))
	if err := t.InvokeAPI(http.MethodDelete, endpoint, nil, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// 移除设备
// https://developer.tuya.com/cn/docs/cloud/c06cc49a20?id=Kaj5805s9jsiu
func (t *TuyaAPIClient) DeviceRemoveSingle(request *model.DeviceRemoveSingleRequest) (*model.DeviceRemoveSingleResponse, error) {
	var response model.DeviceRemoveSingleResponse
	endpoint := fmt.Sprintf(EndpointDeviceRemoveSingle, request.DeviceID)
	if err := t.InvokeAPI(http.MethodDelete, endpoint, nil, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// 恢复出厂设置
// https://developer.tuya.com/cn/docs/cloud/4a9c660db3?id=Kag2y0zpy3e4g
func (t *TuyaAPIClient) DeviceRestoreFactoryDefault(request *model.DeviceRestoreFactoryDefaultRequest) (*model.DeviceRestoreFactoryDefaultResponse, error) {
	var response model.DeviceRestoreFactoryDefaultResponse
	endpoint := fmt.Sprintf(EndpointDeviceRestoreFactoryDefault, request.DeviceID)
	if err := t.InvokeAPI(http.MethodPost, endpoint, nil, &response); err != nil {
		return nil, err
	}
	return &response, nil
}
