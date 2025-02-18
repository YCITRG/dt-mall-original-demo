package entity

type R struct {
	Code  string         `json:"code"`
	Msg   string         `json:"msg"`
	Data  map[string]any `json:"data"`
	Error string         `json:"error"`
}

func NewR(Code string, Msg string) *R {
	return &R{Code: Code, Msg: Msg, Data: nil, Error: ""}
}

func (receiver *R) WithMsg(error string) *R {
	receiver.Error = error
	return receiver
}

func (receiver *R) WithData(data map[string]any) *R {
	receiver.Data = data
	return receiver
}

func (receiver *R) WithError(error string) *R {
	receiver.Error = error
	return receiver
}

var (
	ResOk          = NewR("0", "成功")
	ResUnKnowError = NewR("9999", "未知错误")
)
