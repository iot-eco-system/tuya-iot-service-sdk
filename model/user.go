package model

// 同步用户

type UserSyncRequest struct {
	// App schema
	Schema string `json:"-"`

	// body	true国家码
	CountryCode string `json:"country_code"`
	// body	true用户名
	Username string `json:"username"`
	// body	true用户密码，建议使用 MD5 Hash 原密码
	Password string `json:"password"`
	// body	true用户名类型，默认值为 3，取值说明如下：1：手机号, 2：邮箱号, 3：其他
	UsernameType int `json:"username_type"`
	// body	false昵称
	Nickname string `json:"nick_name"`
	// body	false时区 ID
	TimeZoneID string `json:"time_zone_id"`
}

type UserSyncResult struct {
	UID string `json:"uid"`
}

type UserSyncResponse struct {
	BaseResponse
	Result UserSyncResult `json:"result"`
}

// 用户信息

type UserProperty struct {
	// 用户属性代码
	Code string `json:"code"`
	// 用户属性代码取值
	Value string `json:"value"`
}

type UserInfo struct {
	// 用户 UID
	UID string `json:"uid"`
	// 用户名
	Username string `json:"username"`
	// 国家码
	CountryCode string `json:"country_code"`
	// 手机号码 (仅手机号注册用户支持）
	Mobile string `json:"mobile"`
	// 邮箱（仅邮箱注册用户支持）
	Email string `json:"email"`
	// 昵称
	Nickname string `json:"nick_name"`
	// 头像
	Avatar string `json:"avatar"`
	// 创建时间
	CreateTime int64 `json:"create_time"`
	// 更新时间
	UpdateTime int64 `json:"update_time"`
	// 用户属性
	UserProperties []UserProperty `json:"user_properties"`
	// 时区ID
	TimeZoneID string `json:"time_zone_id"`
	// 温度单位
	TempUnit int `json:"temp_unit"`
}

type UserInfoGetRequest struct {
	UID string `json:"uid"`
}

type UserInfoGetResponse struct {
	BaseResponse
	Result UserInfo `json:"result"`
}

// 用户列表

type UserInfoListRequest struct {
	// uri	true	应用名
	Schema string `query:"-"`
	// query	true	当前页数
	PageNo int `query:"page_no"`
	// query	true	每页的行数。取值范围：0~100
	PageSize int `query:"page_size"`
	// query	false	用户名
	Username string `query:"username"`
	// query	false	开始时间，以 10 位时间戳表示，默认应用创建时间
	StartTime int64 `query:"start_time"`
	// query	false	截止时间，以 10 位时间戳表示。当 start_time 为 Null 时，结束时间默认为当前时间，否则默认为开始时间一个月后
	EndTime int64 `query:"end_time"`
}

type UserInfoListResult struct {
	// 返回的值对象集合
	UserList []UserInfo `json:"list"`
	// 判断是否存在更多数据。true ：有,false ：无
	HasMore bool `json:"has_more"`
	// 用户总数
	Total int64 `json:"total"`
}

type UserInfoListResponse struct {
	BaseResponse
	Result UserInfoListResult `json:"result"`
}

// 关联用户授权登录

type UserAuthorizedLoginRequest struct {
	// true	国家码
	CountryCode string `json:"country_code"`
	// true	用户名
	Username string `json:"username"`
	// true	密码
	Password string `json:"password"`
	// true	App标识
	Schema string `json:"schema"`
}

type UserAuthorizedLoginResult struct {
	// 访问令牌
	AccessToken string `json:"access_token"`
	// 有效时间，单位：秒
	ExpireTime int `json:"expire_time"`
	// 刷新令牌
	RefreshToken string `json:"refresh_token"`
	// 涂鸦用户 ID
	UID string `json:"uid"`
}

type UserAuthorizedLoginResponse struct {
	BaseResponse
	Result UserAuthorizedLoginResult `json:"result"`
}
