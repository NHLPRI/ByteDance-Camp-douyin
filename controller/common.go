package controller

//公共的响应结构体，每个响应结构体都包含该结构体
type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}
