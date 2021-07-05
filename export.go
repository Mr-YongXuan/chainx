package chainx

import (
	"fmt"
	"github.com/Mr-YongXuan/chainx/core"
	"github.com/Mr-YongXuan/chainx/include"
	"github.com/Mr-YongXuan/chainx/lib"
)

type Chainx struct{
	Addr string
	Port int
	PortReuse bool
	chainxRouter *include.ChRouters
}

func New(addr string, port int, portReuse bool) *Chainx {
	return &Chainx{
		Addr: addr,
		Port: port,
		PortReuse: portReuse,
		chainxRouter: core.ChainxRouters(),
	}
}

func (c *Chainx)Add(router string, methods []int, handler func(req *include.ChRequest, res *include.ChResponse) *include.ChResponse) {
	c.chainxRouter.Add(router, methods, handler)
}

func (c *Chainx) StartService() {
	fmt.Printf("ready to listen: http://%s:%d\n", c.Addr, c.Port)
	core.EventStartup(c.Addr, c.Port, c.PortReuse)
}

/* Sessions */
func NewSessions(defaultExpire int64) *lib.Sessions {
	return lib.NewSessions(defaultExpire)
}