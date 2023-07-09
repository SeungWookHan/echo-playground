package protocol

import "strings"

// 모든 응답의 헤더
type RespHeader struct {
	Result       ResultCode `json:"result"`
	ResultString string     `json:"result_string"`
	Message      string     `json:"message"`
}

func NewRespHeader(resultCode ResultCode, message ...string) *RespHeader {
	return &RespHeader{
		Result:       resultCode,
		ResultString: resultCode.toString(),
		Message:      strings.Join(message, ","),
	}
}
