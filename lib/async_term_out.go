package lib

import (
	"fmt"
	"time"
)

var TermLogger = true

type LogsOut struct {
	ChainxRunning    bool
	AccessInfoBuffer chan string
}

func NewLogsOut() *LogsOut {
	to := &LogsOut{
		ChainxRunning: true,
	}
	to.AccessInfoBuffer = make(chan string)
	go to.programLogsOutput()
	return to
}

/* SignalClose close program logs output before exit */
func (to *LogsOut) SignalClose() {
	to.ChainxRunning = false
	time.Sleep(time.Second)
	to.AccessInfoBuffer <- "Closed"
}

/* programLogsOutput loop print logs to terminal (unblock) */
func (to *LogsOut) programLogsOutput() {
	for to.ChainxRunning {
		currentTime := time.Now().Format("2006-01-02 03:04:05")
		fmt.Printf("%s %s\n", currentTime, <-to.AccessInfoBuffer)
	}
}

/* AccessInfo append a new log into output queue */
func (to *LogsOut) AccessInfo(stCode []byte, method string, resource []byte, usages int64) {
	if TermLogger {
		to.AccessInfoBuffer <- fmt.Sprintf("-> %s -> %s -> %s -> %dμs", stCode, method, resource, usages)
	}
}
