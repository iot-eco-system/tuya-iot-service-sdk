package tuyasdk

//
// Auth API
//

const (
	// 获取身份token
	EndpointGetAccessToken = "/v1.0/token?grant_type=1"
	// 刷新用户令牌
	EndpointRefreshToken = "/v1.0/token/"
)

//
// Device API
//

const (
	// 获取设备列表
	EndpointDeviceList = "/v1.2/iot-03/devices"
	// 获取设备详情
	EndpointDeviceGetInfo = "/v1.1/iot-03/devices/%s"
	// 下发指令
	EndpointDeviceSendCommand = "/v1.0/iot-03/devices/%s/commands"
	// 批量获取设备最新状态
	EndpointDeviceGetStatusMulti = "/v1.0/iot-03/devices/status"
	// 获取单个设备的状态
	EndpointDeviceGetStatusSingle = "/v1.0/iot-03/devices/%s/status"
	// 修改设备信息
	EndpointDeviceUpdateInfo = "/v1.0/iot-03/devices/%s"
	// 移除设备
	EndpointDeviceRemoveSingle = "/v1.0/iot-03/devices/%s"
	// 移除设备
	EndpointDeviceRemoveMulti = "/v1.0/iot-03/devices?device_ids=%s"
	// 恢复出厂设置
	EndpointDeviceRestoreFactoryDefault = "/v1.0/iot-03/devices/%s/actions/reset"
)

//
// Asset API
//

const (
	// 创建资产
	EndpointAssetCreate = "/v1.0/iot-02/assets"
	// 删除资产
	EndpointAssetDelete = "/v1.0/iot-02/assets/%s"
	// 批量查询资产信息
	EndpointAssetList = "/v1.0/iot-02/assets?asset_ids=%s"
	// 修改资产
	EndpointAssetUpdate = "/v1.0/iot-02/assets/%s"
	// 批量资产授权用户（支持选择是否授权子节点）
	EndpointAssetAuthorize = "/v1.0/iot-03/users/%s/actions/batch-assets-authorized"
)

//
// User API
//

const (
	// 同步用户
	EndpointUserSync = "/v1.0/apps/%s/user"
	// 获取用户信息
	EndpointUserInfoGet = "/v1.0/users/%s/infos"
	// 用户列表
	EndpointUserList = "/v2.0/apps/%s/users"
	// 关联用户授权登录
	EndpointUserAuthorizedLogin = "/v1.0/iot-01/associated-users/actions/authorized-login"
)
