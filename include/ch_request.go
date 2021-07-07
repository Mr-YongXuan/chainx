package include

import (
	"bytes"
	"github.com/tidwall/gjson"
)

type ChRequest struct {
	Method   int
	Resource []byte
	Version  []byte
	Args     []byte
	ArgsMap  map[string]string
	Header   map[string][]byte
	Body     []byte
}

/* NewRequest Request struct build func */
func NewRequest() *ChRequest {
	cr := &ChRequest{}
	cr.Header = make(map[string][]byte)
	return cr
}

/* GetHeader using the keyword fetch header from headers */
func (cr *ChRequest) GetHeader(key string) string {
	if v, ok := cr.Header[key]; ok {
		return string(v)
	}
	return ""
}

/* GetData get client request payload */
func (cr *ChRequest) GetData() string {
	return string(cr.Body)
}

/* GetHttpVersion get the client http request version */
func (cr *ChRequest) GetHttpVersion() string {
	return string(cr.Version)
}

/* GetMethod get the client http request method */
func (cr *ChRequest) GetMethod() string {
	switch cr.Method {
	case ChHttpGet:
		return "GET"

	case ChHttpPost:
		return "POST"

	case ChHttpPut:
		return "PUT"

	case ChHttpOption:
		return "OPTIONS"

	case ChHttpHead:
		return "HEAD"
	}
	return "UNKNOWN"
}

/* GetQueryArgument get the client http request argument */
func (cr *ChRequest) GetQueryArgument(key string) string {
	if cr.ArgsMap == nil {
		/* initial query argument */
		cr.ArgsMap = make(map[string]string)

		/* parse query argument */
		var buffer bytes.Buffer
		var argumentKey string
		for _, char := range cr.Args {
			switch char {
			case 0x3d: // =
			    argumentKey = buffer.String()
			    buffer.Reset()

			case 0x26: // &
			    cr.ArgsMap[argumentKey] = buffer.String()
			    argumentKey = ""
			    buffer.Reset()

			default:
				buffer.WriteByte(char)
			}
		}

		if argumentKey != "" {
			cr.ArgsMap[argumentKey] = buffer.String()
		}
	}

	if v, ok := cr.ArgsMap[key]; ok {
		return v
	}
	return ""
}

/* GetJson try to fetch json data from request body */
func (cr *ChRequest) GetJson(path string) (res gjson.Result) {
	res = gjson.GetBytes(cr.Body, path)
	return
}

type ChRouters struct {
	Routers map[string]struct {
		Method  []int
		Handler func(req *ChRequest, res *ChResponse) *ChResponse
	}
}

/* Add add a router into chainx routers map */
func (cr *ChRouters) Add(router string, methods []int, handler func(req *ChRequest, res *ChResponse) *ChResponse) {
	cr.Routers[router] = struct {
		Method  []int
		Handler func(req *ChRequest, res *ChResponse) *ChResponse
	}{Method: methods, Handler: handler}
}
