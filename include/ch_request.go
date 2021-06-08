package include

type ChRequest struct {
	RequestMethod   int
	RequestResource []byte
	RequestVersion  []byte
	RequestArgs     []byte
	RequestHeader   map[string][]byte
	RequestBody     []byte
}

func NewRequest() *ChRequest {
	cr := &ChRequest{}
	cr.RequestHeader = make(map[string][]byte)
	return cr
}

type ChRouters struct {
	Routers map[string]struct {
		Method  []int
		Handler func(req *ChRequest, res *ChResponse) *ChResponse
	}
}

func (cr *ChRouters) Add(router string, methods []int, handler func(req *ChRequest, res *ChResponse) *ChResponse) {
	cr.Routers[router] = struct {
		Method  []int
		Handler func(req *ChRequest, res *ChResponse) *ChResponse
	}{Method: methods, Handler: handler}
}
