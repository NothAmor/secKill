package secgin

// Response the unified json structure
type response struct {
	Code    int         `json:"code"`
	Stat    int         `json:"stat"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}

// NilErr 传入nil err, data返回空结构体
type NilErr struct {
}

func Success(v interface{}) interface{} {
	ret := response{Stat: 1, Code: 0, Message: "ok", Data: v}
	return ret
}

func Error(e SecError, data interface{}) interface{} {
	if data == nil {
		data = NilErr{}
	}
	ret := response{Stat: 0, Code: e.Code, Message: e.Msg, Data: data}
	return ret
}
