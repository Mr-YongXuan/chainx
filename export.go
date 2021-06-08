package chainx

import (
	"fmt"
	"github.com/Mr-YongXuan/chainx/core"
	"github.com/Mr-YongXuan/chainx/include"
)

func GetRoutersMap() *include.ChRouters {
	return core.InitialRouters()
}

func StartService(addr string, port int) {
	fmt.Printf("ready to listen: http://%s:%d\n", addr, port)
	core.EventStartup(addr, port)
}