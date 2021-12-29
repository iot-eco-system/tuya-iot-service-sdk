package model

type Response interface {
	IsSuccess() bool
	GetErrorCode() int
	GetErrorMessage() string
}

type BaseResponse struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
	Success bool   `json:"success"`
	T       int64  `json:"t"`
}

func (b BaseResponse) IsSuccess() bool {
	return b.Success
}

func (b BaseResponse) GetErrorCode() int {
	return b.Code
}

func (b BaseResponse) GetErrorMessage() string {
	return b.Message
}
