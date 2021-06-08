package include

import "encoding/json"

type ChResponse struct {
	Code    []byte
	Headers map[string]string
	Body    []byte
	Ctl     map[string]string
}

func NewChResponse() *ChResponse {
	cr := &ChResponse{}
	cr.Headers = make(map[string]string)
	cr.Ctl = make(map[string]string)
	cr.Headers["Content-Type"] = "text/plain"
	cr.Code = St200
	return cr
}

func (cr *ChResponse) SetReturnJson() {
	cr.Headers["Content-Type"] = "application/json"
}

func (cr *ChResponse) Jsonify(v map[string]interface{}) (ok bool) {
	cr.Headers["Content-Type"] = "application/json"
	out, err := json.Marshal(v)
	if err != nil {
		cr.Body = []byte{0x7b, 0x7d} // {}
		return false
	}
	cr.Body = out
	return true
}