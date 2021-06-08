package chainx

import (
	"github.com/Mr-YongXuan/chainx/core"
	"github.com/Mr-YongXuan/chainx/include"
)

func GetRoutersMap() *include.ChRouters {
	return core.InitialRouters()
}

func StartService() {
	core.EventStartup()
}