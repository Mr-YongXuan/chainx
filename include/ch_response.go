package include

type ChResponse struct {
	ResponseCode    []byte
	ResponseHeaders map[string]string
	ResponseBody    []byte
	ResponseCtl     map[string]string
}

func NewChResponse() *ChResponse {
	cr := &ChResponse{}
	cr.ResponseHeaders = make(map[string]string)
	cr.ResponseCtl = make(map[string]string)
	cr.ResponseCode = St200
	return cr
}

func (cr *ChResponse) SetReturnJson() {
	cr.ResponseHeaders["Content-Type"] = "application/json"
}
