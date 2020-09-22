package model

// Return Codes
const (
	CodeOK = iota
	CodeErr
)

func init() {
}

// RespOK returns when no error happen
type RespOK struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

// RespErr returns when error happen
type RespErr struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
