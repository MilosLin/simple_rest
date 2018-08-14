// Package protocol : api相關協定，包含可以回傳的error code訊息.
package protocol

import "fmt"

// Response : api 回應格式
type Response struct {
	Code    int         `json:"Code"`
	Message string      `json:"Message"`
	Result  interface{} `json:"Result"`
}

var (
	SomethingWrongRes = func(err error) *Response {
		return &Response{
			Code:    1,
			Message: fmt.Sprintf("Something Wrong. Err:%s", err.Error()),
		}
	}
)
