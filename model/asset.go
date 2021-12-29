package model

// 删除资产

type AssetDeleteRequest struct {
	AssetID string `json:"-"`
}

type AssetDeleteResponse struct {
	BaseResponse
	Result bool `json:"result"`
}

// 批量查询资产信息

type AssetInfoListRequest struct {
	AssetIDs []string `json:"-"`
}

type AssetInfo struct {
	// 资产 ID
	AssetID string `json:"asset_id"`
	// 父级 ID，最顶级为 -1
	ParentAssetID string `json:"parent_asset_id"`
	// 资产名称
	AssetName string `json:"asset_name"`
	// 资产完整名称
	AssetFullName string `json:"asset_full_name"`
	// 资产层级，根节点为 0
	Level int `json:"level"`
}

type AssetInfoListResponse struct {
	BaseResponse
	Result []AssetInfo `json:"result"`
}

// 修改资产

type AssetInfoUpdateRequest struct {
	// 资产 ID
	AssetID string `json:"-"`
	// 资产名称
	Name string `json:"name"`
	// 元数据 ID（节点存储说明 ID）
	MetaID string `json:"meta_id"`
}

type AssetInfoUpdateResponse struct {
	BaseResponse
	Result bool `json:"result"`
}

// 批量资产授权用户（支持选择是否授权子节点）

type AssetAuthorizeRequest struct {
	UID                string   `json:"uid"`
	AssetIDs           []string `json:"asset_ids"`
	AuthorizedChildren bool     `json:"authorizedChildren"`
}

type AssetAuthorizeResponse struct {
	BaseResponse
	Result bool `json:"result"`
}

// 新增资产

type AssetCreateRequest struct {
	// true	资产名称
	Name string `json:"name"`
	// false	元数据 ID（节点存储说明 ID）
	MetaID string `json:"meta_id"`
	// false	父资产 ID
	ParentAssetID string `json:"parent_asset_id"`
}

type AssetCreateResponse struct {
	BaseResponse
	Result bool `json:"result"`
}
