package model

type TokenResult struct {
	AccessToken  string `json:"access_token"`
	ExpireTime   int    `json:"expire_time"`
	RefreshToken string `json:"refresh_token"`
	UID          string `json:"uid"`
}

type TokenResponse struct {
	BaseResponse
	Result TokenResult `json:"result"`
}
