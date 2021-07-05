package core

import (
	"bytes"
	"github.com/Mr-YongXuan/chainx/include"
	"github.com/Mr-YongXuan/chainx/lib"
	"time"
)

var logOut = lib.NewLogsOut()

// ChainxRSM chainx resource manager
func ChainxRSM(chr *include.ChRequest) (out []byte) {
	startTime := time.Now()
	var res = include.NewChResponse()
	var buffer bytes.Buffer
	var reStCode = include.St200
	/* router handler */
	if st, ok := cr.Routers[string(chr.Resource)]; ok {
		/* method allow check */
		if lib.ChMethodIsApprove(chr.Method, cr.Routers[string(chr.Resource)].Method) {
			res = st.Handler(chr, include.NewChResponse())
			lib.BasicResponseHeaders(&buffer, include.HttpVer11, include.St200, len(res.Body))
			/* options approve */
			if chr.Method == include.ChHttpOption {
				var tmpOptions string
				for _, method := range cr.Routers[string(chr.Resource)].Method {
					switch method {
					case include.ChHttpGet:
						tmpOptions = tmpOptions + "GET, "
						break

					case include.ChHttpPost:
						tmpOptions = tmpOptions + "POST, "
						break

					case include.ChHttpPut:
						tmpOptions = tmpOptions + "PUT, "
						break
					}
				}
				res.Headers["Allow"] = tmpOptions + "HEAD, OPTIONS"
			}

		} else {
			reStCode = include.St405
			res = include.ChMethodNotAllowed()
			lib.BasicResponseHeaders(&buffer, include.HttpVer11, include.St405, len(res.Body))
		}

	} else {
		/* static resource or direct 404 */
		reStCode = include.St404
		res = include.ChNotFound()
		lib.BasicResponseHeaders(&buffer, include.HttpVer11, include.St404, len(res.Body))
	}

	if len(res.Headers) != 0 {
		for headerKey, headerValue := range res.Headers {
			buffer.WriteString(headerKey + ": " + headerValue + "\r\n")
		}
	}

	buffer.WriteString("\r\n")
	if chr.Method != include.ChHttpHead {
		buffer.Write(res.Body)
	}

	logOut.AccessInfo(reStCode, chr.GetMethod(), chr.Resource, time.Since(startTime).Microseconds())
	return buffer.Bytes()
}
