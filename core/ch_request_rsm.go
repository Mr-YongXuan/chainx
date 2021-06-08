package core

import (
	"bytes"
	"github.com/Mr-YongXuan/chainx/include"
	"github.com/Mr-YongXuan/chainx/lib"
)

// ChainxRSM chainx resource manager
func ChainxRSM(chr *include.ChRequest) (out []byte) {
	var res = include.NewChResponse()
	var buffer bytes.Buffer
	/* router handler */
	if st, ok := cr.Routers[string(chr.RequestResource)]; ok {
		/* method allow check */
		if lib.ChMethodIsApprove(chr.RequestMethod, cr.Routers[string(chr.RequestResource)].Method) {
			res = st.Handler(chr, include.NewChResponse())
			lib.BasicResponseHeaders(&buffer, include.HttpVer11, include.St200, len(res.ResponseBody))
			/* options approve */
			if chr.RequestMethod == include.ChHttpOption {
				var tmpOptions string
				for _, method := range cr.Routers[string(chr.RequestResource)].Method {
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
				res.ResponseHeaders["Allow"] = tmpOptions + "HEAD, OPTIONS"
			}

		} else {
			res = include.ChMethodNotAllowed()
			lib.BasicResponseHeaders(&buffer, include.HttpVer11, include.St405, len(res.ResponseBody))
		}

	} else {
		/* static resource or direct 404 */
		res = include.ChNotFound()
		lib.BasicResponseHeaders(&buffer, include.HttpVer11, include.St404, len(res.ResponseBody))
	}

	if len(res.ResponseHeaders) != 0 {
		for headerKey, headerValue := range res.ResponseHeaders {
			buffer.WriteString(headerKey + ": " + headerValue + "\r\n")
		}
	}

	buffer.WriteString("\r\n")
	if chr.RequestMethod != include.ChHttpHead {
		buffer.Write(res.ResponseBody)
	}

	return buffer.Bytes()
}
