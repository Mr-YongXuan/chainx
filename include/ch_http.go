package include

const (
	/* request methods */
	ChHttpGet    int = 0x01
	ChHttpPut    int = 0x02
	ChHttpPost   int = 0x03
	ChHttpHead   int = 0x04
	ChHttpOption int = 0x05

	/* ASCII */
	ChStrSpace = 0x20
	ChStrLf    = 0x0a
	ChStrCr    = 0x0d
	ChStrQry   = 0x3f
	ChStrColon = 0x3a
)

type JSON map[string]interface{}

var (
	/* Http versions */
	HttpVerUn = []byte{0x75, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e}       //unknown
	HttpVer10 = []byte{0x48, 0x54, 0x54, 0x50, 0x2f, 0x31, 0x2e, 0x30} // HTTP/1.0
	HttpVer11 = []byte{0x48, 0x54, 0x54, 0x50, 0x2f, 0x31, 0x2e, 0x31} // HTTP/1.1
	HttpVer20 = []byte{0x48, 0x54, 0x54, 0x50, 0x2f, 0x32, 0x2e, 0x30} // HTTP/2.0

	/* http connection */
	HttpClose = []byte{0x63, 0x6c, 0x6f, 0x73, 0x65} // close

	/* Http Status Code */
	St200 = []byte(" 200 OK")
	St403 = []byte(" 403 Forbidden")
	St404 = []byte(" 404 Not Found")
	St405 = []byte(" 405 Method Not Allowed")
)

// ChForbidden return 403 html content
func ChForbidden() *ChResponse {
	res := NewChResponse()
	res.Code = St403
	res.Headers["Content-Type"] = "text/html"
	res.Body = []byte(`
<div style="text-align: center;">
<h1>403 Forbidden</h1>
</div>
<hr />
<div style="text-align: center;">
chainx
</div>
`)
	return res
}

// ChNotFound return 404 html content
func ChNotFound() *ChResponse {
	res := NewChResponse()
	res.Code = St404
	res.Headers["Content-Type"] = "text/html"
	res.Body = []byte(`
<div style="text-align: center;">
<h1>404 Not Found</h1>
</div>
<hr />
<div style="text-align: center;">
chainx
</div>
`)
	return res
}

// ChMethodNotAllowed return 405 html content
func ChMethodNotAllowed() *ChResponse {
	res := NewChResponse()
	res.Code = St405
	res.Headers["Content-Type"] = "text/html"
	res.Body = []byte(`
<div style="text-align: center;">
<h1>405 Method Not Allowed</h1>
</div>
<hr />
<div style="text-align: center;">
chainx
</div>
`)
	return res
}
