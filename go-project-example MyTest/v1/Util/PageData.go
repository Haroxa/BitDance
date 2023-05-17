package Util

type PageData struct {
	Code int64       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func NewPageData(code int64, msg string, data interface{}) *PageData {
	return &PageData{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}
