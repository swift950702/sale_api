package utils

type JsonStruct struct {
	Code  int         `json:"code"`
	Msg   string      `json:"msg,omitempty"`
	Data  interface{} `json:"data"`
	Count int64       `json:"count,omitempty"`
}

func ReturnSuccess(msg string, data interface{}) (json *JsonStruct) {
	json = &JsonStruct{Code: 1, Msg: msg, Data: data, Count: 0}
	return
}

func ReturnError(code int, msg string) *JsonStruct {
	json := &JsonStruct{Code: code, Msg: msg}
	return json
}
