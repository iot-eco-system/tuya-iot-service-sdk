package model

// 下发指令

type DeviceProperty struct {
	Code  string      `json:"code"`
	Value interface{} `json:"value"`
}

type DeviceSendCommandRequest struct {
	DeviceID string           `json:"-"`
	Commands []DeviceProperty `json:"commands"`
}

type DeviceSendCommandResponse struct {
	BaseResponse
	Result bool `json:"result"`
}

// 批量获取设备最新状态

type DeviceGetStatusMultiRequest struct {
	DeviceIDs []string `json:"-"`
}

type DeviceGetStatusMultiResult struct {
	DeviceID string           `json:"id"`
	Status   []DeviceProperty `json:"status"`
}

type DeviceGetStatusMultiResponse struct {
	BaseResponse
	Result DeviceGetStatusMultiResult `json:"result"`
}

// 获取单个设备的状态

type DeviceGetStatusSingleRequest struct {
	DeviceID string `json:"-"`
}

type DeviceGetStatusSingleResponse struct {
	BaseResponse
	Result []DeviceProperty `json:"result"`
}

// 获取设备信息

type DeviceInfo struct {
	// 设备 ID
	ID string `json:"id"`
	// 设备 uuid
	UUID string `json:"uuid"`
	// 产品品类
	Category string `json:"category"`
	// 设备名称
	Name string `json:"name"`
	// 产品 ID
	ProductID string `json:"product_id"`
	// 产品名称
	ProductName string `json:"product_name"`
	// 密钥
	LocalKey string `json:"local_key"`
	// 是否为子设备
	Sub bool `json:"sub"`
	// 资产 ID
	AssetID string `json:"asset_id"`
	// 设备 IP
	IP string `json:"ip"`
	// 经度
	Lon string `json:"lon"`
	// 纬度
	Lat string `json:"lat"`
	// 产品型号
	Model string `json:"model"`
	// 时区
	TimeZone string `json:"time_zone"`
	// 激活时间
	ActiveTime int64 `json:"active_time"`
	// 更新时间
	UpdateTime int64 `json:"update_time"`
	// 初次配网时间
	CreateTime int64 `json:"create_time"`
	// 是否在线
	Online bool `json:"online"`
}

type DeviceGetInfoRequest struct {
	DeviceID string `json:"-"`
}

type DeviceGetInfoResponse struct {
	BaseResponse
	Result DeviceInfo `json:"result"`
}

// 修改设备信息

type DeviceUpdateInfoRequest struct {
	DeviceID string `json:"-"`
	Name     string `json:"name"`
}

type DeviceUpdateInfoResponse struct {
	BaseResponse
	Result bool `json:"result"`
}

// 批量移除设备

type DeviceRemoveMultiRequest struct {
	DeviceIDs []string `json:"-"`
}

type DeviceRemoveMultiResponse struct {
	BaseResponse
	Result bool `json:"result"`
}

// 移除设备

type DeviceRemoveSingleRequest struct {
	DeviceID string `json:"-"`
}

type DeviceRemoveSingleResponse struct {
	BaseResponse
	Result bool `json:"result"`
}

// 恢复出厂设置

type DeviceRestoreFactoryDefaultRequest struct {
	DeviceID string `json:"-"`
}

type DeviceRestoreFactoryDefaultResponse struct {
	BaseResponse
	Result bool `json:"result"`
}

// 设备列表

type DeviceListRequest struct {
	// 包含不同维度的设备
	// * asset 行业资产维度的设备
	// * homeApp 家庭体系下 App 小程序维度关联的设备
	// * tuyaUser 涂鸦智能账号维度设备。
	// 默认值为 asset。
	SourceType string `query:"source_type"`
	// 在 source_type 不同时传对应的业务 ID。
	// * asset维度时 传资产 ID
	// * homeApp维度时传递 Schema
	// * tuyaUser维度时传递 UID
	SourceID string `query:"source_id"`
	// 设备 ID，最多 20 个
	DeviceIDs string `query:"device_ids"`
	// 设备名称
	Name string `query:"name"`
	// 产品品类
	Category string `query:"category"`
	// 产品 ID
	ProductID string `query:"product_id"`
	// 上一页返回的分页标识
	LastRowKey string `query:"last_row_key"`
	// 分页大小，默认 20，最大 200
	PageSize int `query:"page_size"`
}

type DeviceListResult struct {
	// 是否有下一页
	HasMore bool `json:"has_more"`
	// 返回数据
	List []DeviceListInfo `json:"list"`
	// 上一页的主键标识
	LastRowKey string `json:"last_row_key"`
	// 符合条件的数据总条数
	Total int64 `json:"total"`
}

type DeviceListInfo struct {
	// 设备ID
	ID string `json:"id"`
	// 网关ID, 非网关子设备时为空
	GatewayID string `json:"gateway_id"`
	// 节点ID, 非网关子设备时为空
	NodeID string `json:"node_id"`
	// 设备uuid
	UUID string `json:"uuid"`
	// 产品品类
	Category string `json:"category"`
	// 产品品类名称
	CategoryName string `json:"category_name"`
	// 设备名称
	Name string `json:"name"`
	// 产品ID
	ProductID string `json:"product_id"`
	// 产品名称
	ProductName string `json:"product_name"`
	// 密钥
	LocalKey string `json:"local_key"`
	// 是否为子设备
	Sub bool
	// 资产ID
	AssetID string `json:"asset_id"`
	// 家庭ID
	OwnerID string `json:"owner_id"`
	// 设备 IP
	IP string `json:"ip"`
	// 经度
	Lon string `json:"lon"`
	// 纬度
	Lat string `json:"lat"`
	// 产品型号
	Model string `json:"model"`
	// 时区
	TimeZone string `json:"time_zone"`
	// 激活时间
	ActiveTime int64 `json:"active_time"`
	// 更新时间
	UpdateTime int64 `json:"update_time"`
	// 初次配网时间
	CreateTime int64 `json:"create_time"`
	// 在线状态
	Online bool `json:"online"`
	// 设备图标
	Icon string `json:"icon"`
}

type DeviceListResponse struct {
	BaseResponse
	Result DeviceListResult `json:"result"`
}
