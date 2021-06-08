# chainx
enjoy your web server~

# Installation
```sh
$ go get -u github.com/Mr-YongXuan/chainx
```

# Import
```go
import (
	"github.com/Mr-YongXuan/chainx"
	"github.com/Mr-YongXuan/chainx/include"
)
```

# Example
```go
package main

import (
	"github.com/Mr-YongXuan/chainx"
	"github.com/Mr-YongXuan/chainx/include"
)

func hello(req *include.ChRequest, res *include.ChResponse) *include.ChResponse {
	res.Jsonify(include.JSON{
		"message": "Hello,Chainx!",
		"user-Agent": req.GetHeader("User-Agent"),
	})

	return res
}

func main() {
	c := chainx.New("127.0.0.1", 8080, true)
	c.Add("/hello", []int{include.ChHttpGet}, hello)
	c.StartService()
}
```
