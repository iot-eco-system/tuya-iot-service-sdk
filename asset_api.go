package tuyasdk

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/iot-eco-system/tuya-iot-service-sdk/model"
)

// 新增资产
// https://developer.tuya.com/cn/docs/cloud/5ed35de173?id=Kag2yq4x8xtlj
func (t *TuyaAPIClient) AssetCreate(request *model.AssetCreateRequest) (*model.AssetCreateResponse, error) {
	var response model.AssetCreateResponse
	if err := t.InvokeAPI(http.MethodPost, EndpointAssetCreate, nil, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// 删除资产
// https://developer.tuya.com/cn/docs/cloud/d02323ac23?id=Kag2yp44mpfff
func (t *TuyaAPIClient) AssetDelete(request *model.AssetDeleteRequest) (*model.AssetDeleteResponse, error) {
	var response model.AssetDeleteResponse
	endpoint := fmt.Sprintf(EndpointAssetDelete, request.AssetID)
	if err := t.InvokeAPI(http.MethodDelete, endpoint, nil, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// 批量查询资产信息
// https://developer.tuya.com/cn/docs/cloud/22f9b5f5de?id=Kag2yprgv4nqa
func (t *TuyaAPIClient) AssetInfoList(request *model.AssetInfoListRequest) (*model.AssetInfoListResponse, error) {
	var response model.AssetInfoListResponse
	endpoint := fmt.Sprintf(EndpointAssetList, strings.Join(request.AssetIDs, ","))
	if err := t.InvokeAPI(http.MethodGet, endpoint, nil, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// 修改资产
// https://developer.tuya.com/cn/docs/cloud/2c4cd4ecb4?id=Kag2yozk5yww6
func (t *TuyaAPIClient) AssetInfoUpdate(request *model.AssetInfoUpdateRequest) (*model.AssetInfoUpdateResponse, error) {
	var response model.AssetInfoUpdateResponse
	endpoint := fmt.Sprintf(EndpointAssetUpdate, request.AssetID)
	if err := t.InvokeAPI(http.MethodPut, endpoint, request, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// 批量资产授权用户（支持选择是否授权子节点）
// https://developer.tuya.com/cn/docs/cloud/18aa86cb99?id=Kaopmpy6lx7al
func (t *TuyaAPIClient) AssetAuthorize(request *model.AssetAuthorizeRequest) (*model.AssetAuthorizeResponse, error) {
	var response model.AssetAuthorizeResponse
	endpoint := fmt.Sprintf(EndpointAssetAuthorize, request.UID)
	if err := t.InvokeAPI(http.MethodPost, endpoint, request, &response); err != nil {
		return nil, err
	}
	return &response, nil
}
