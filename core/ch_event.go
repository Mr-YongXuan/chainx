package core

import (
	"fmt"
	"github.com/Mr-YongXuan/chainx/include"
	"github.com/panjf2000/gnet"
	"github.com/panjf2000/gnet/pool/goroutine"
	"log"
	"strings"
	"sync"
	"time"
)

var cr = &include.ChRouters{}
var opt sync.Mutex

type chainxServer struct {
	*gnet.EventServer
	pool      *goroutine.Pool
	conn      map[gnet.Conn]int
	blackList map[string]struct {
		HowLong int
		Active  string
		Reason  string
	}
}

/* OnOpened append timer for timeout disconnect */
func (cs *chainxServer) OnOpened(c gnet.Conn) (out []byte, action gnet.Action) {
	/* black list */
	addr := c.RemoteAddr()
	if banInfo, ok := cs.blackList[strings.Split(addr.String(), ":")[0]]; ok {
		action = gnet.Close
		fmt.Println("client:", addr.String(), "trying to connect, but has been banned, reason:", banInfo.Reason, "Active:", banInfo.Active)
		return
	}

	/* keep-alive */
	opt.Lock()
	if _, ok := cs.conn[c]; !ok {
		cs.conn[c] = 30
	}
	opt.Unlock()
	return
}

/* OnClosed client or server active closed, clear timer */
func (cs *chainxServer) OnClosed(c gnet.Conn, _ error) (action gnet.Action) {
	opt.Lock()
	if _, ok := cs.conn[c]; ok {
		delete(cs.conn, c)
	}
	opt.Unlock()
	return
}

/* Tick keep-alive timeout check */
func (cs *chainxServer) Tick() (delay time.Duration, action gnet.Action) {
	go func() {
		for c, t := range cs.conn {
			if t <= 0 {
				/* keep-alive timeout */
				_ = c.Close()
			} else {
				cs.conn[c] = cs.conn[c] - 1
			}
		}
	}()

	delay = time.Second
	return
}

/* React events trans flow here */
func (cs *chainxServer) React(frame []byte, c gnet.Conn) (out []byte, action gnet.Action) {
	/* step:1 -> parse request struct */
	startTime := time.Now()
	_ = cs.pool.Submit(func() {
		buf, ctl := ChainxTRS(frame)
		_ = c.AsyncWrite(buf)

		/* close connection */
		if v, ok := ctl["conn"]; ok && v == "close" {
			_ = c.Close()
		}
	})
	fmt.Printf("usage: %dus\n", time.Since(startTime).Microseconds())
	return
}

func InitialRouters() *include.ChRouters {
	/* initial routers */
	cr.Routers = make(map[string]struct {
		Method  []int
		Handler func(req *include.ChRequest, res *include.ChResponse) *include.ChResponse
	})
	return cr
}

/* EventStartup start chainx */
func EventStartup() {
	/* initial goroutine work pool */
	wp := goroutine.Default()
	defer wp.Release()

	/* initial chainx */
	var chainx = &chainxServer{pool: wp}
	chainx.conn = make(map[gnet.Conn]int)
	chainx.blackList = make(map[string]struct {
		HowLong int
		Active  string
		Reason  string
	})
	//chainx.blackList["127.0.0.1"] = struct {
	//	HowLong int
	//	Active  string
	//	Reason  string
	//}{HowLong: -1, Active: "SYSTEM-chainx", Reason: "out of access"}

	/* initialed, enjoy! TODO REMOVE PORT REUSE BEFORE PRODUCT */
	log.Fatal(gnet.Serve(chainx, "tcp://0.0.0.0:9000", gnet.WithMulticore(true), gnet.WithTicker(true), gnet.WithReusePort(true)))
}
