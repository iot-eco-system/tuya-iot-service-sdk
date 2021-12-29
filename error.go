package tuyasdk

import "errors"

var (
	ErrGetAccessTokenFailed = errors.New("failed to get access token")
	ErrRefreshTokenEmpty    = errors.New("refresh token empty")
)
