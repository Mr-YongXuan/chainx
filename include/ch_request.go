package include

import "bytes"

type ChRequest struct {
	RequestMethod   int
	RequestResource []byte
	RequestVersion  []byte
	RequestArgs     []byte
	RequestArgsMap  map[string]string
	RequestHeader   map[string][]byte
	RequestBody     []byte
}

/* NewRequest Request struct build func */
func NewRequest() *ChRequest {
	cr := &ChRequest{}
	cr.RequestHeader = make(map[string][]byte)
	return cr
}

/* GetHeader using the keyword fetch header from headers */
func (cr *ChRequest) GetHeader(key string) string {
	if v, ok := cr.RequestHeader[key]; ok {
		return string(v)
	}
	return ""
}

/* GetData get client request payload */
func (cr *ChRequest) GetData() string {
	return string(cr.RequestBody)
}

/* GetHttpVersion get the client http request version */
func (cr *ChRequest) GetHttpVersion() string {
	return string(cr.RequestVersion)
}

/* GetMethod get the client http request method */
func (cr *ChRequest) GetMethod() string {
	switch cr.RequestMethod {
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
	if cr.RequestArgsMap == nil {
		/* initial query argument */
		cr.RequestArgsMap = make(map[string]string)

		/* parse query argument */
		var buffer bytes.Buffer
		var argumentKey string
		for _, char := range cr.RequestArgs {
			switch char {
			case 0x3d: // =
			    argumentKey = buffer.String()
			    buffer.Reset()

			case 0x26: // &
			    cr.RequestArgsMap[argumentKey] = buffer.String()
			    buffer.Reset()

			default:
				buffer.WriteByte(char)
			}
		}
	}

	if v, ok := cr.RequestArgsMap[key]; ok {
		return v
	}
	return ""
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
