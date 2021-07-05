package lib

import (
	"bytes"
	"fmt"
	"github.com/Mr-YongXuan/chainx/include"
	"time"
)

func BasicResponseHeaders(buffer *bytes.Buffer, httpVer []byte, stCode []byte, bodyLen int) {
	/* gen: response line */
	buffer.Write(httpVer)
	buffer.Write(stCode)
	buffer.Write([]byte("\r\n"))

	/* gen: GMT web server time format */
	buffer.WriteString("Date: " + time.Now().Format("Mon, 02 Jan 2006 03:04:05 GMT") + "\r\n")

	/* gen: const headers */
	buffer.WriteString("Server: Chainx/0.0.1\r\n")

	/* gen: http/1.0 => close | http/1.1 || http/2.0 => keep-alive */
	if ChHttpVerCmp(httpVer, include.HttpVer11) || ChHttpVerCmp(httpVer, include.HttpVer20) {
		buffer.WriteString("Connection: keep-alive\r\n")
	}

	/* gen: response body length */
	buffer.WriteString(fmt.Sprintf("Content-Length: %d\r\n", bodyLen))
}
