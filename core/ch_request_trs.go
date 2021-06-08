package core

import (
	"github.com/Mr-YongXuan/chainx/include"
	"github.com/Mr-YongXuan/chainx/lib"
)

// ChainxTRS pipeline step->1 transform: uint8_array->go struct
func ChainxTRS(reqBdy []byte) (out []byte, ctl map[string]string) {
	var curr, prevPos = 0, 0
	var tmpHeaderKey = "ch_un"
	chr := include.NewRequest()
	chr.RequestHeader = make(map[string][]byte)
	ctl = make(map[string]string)

	for pos, char := range reqBdy {
		if curr < 10 {
			/* parse request line */
			if char == include.ChStrSpace && curr == 0 {
				/* try to fetch request method */
				curr++
				prevPos = pos
				switch pos {
				case 3:
					if lib.ChStr3Cmp(reqBdy, 'G', 'E', 'T', ' ') {
						chr.RequestMethod = include.ChHttpGet

					} else if lib.ChStr3Cmp(reqBdy, 'P', 'U', 'T', ' ') {
						chr.RequestMethod = include.ChHttpPut

					}

				case 4:
					if lib.ChStr4Cmp(reqBdy, 'P', 'O', 'S', 'T', ' ') {
						chr.RequestMethod = include.ChHttpPost

					} else if lib.ChStr4Cmp(reqBdy, 'H', 'E', 'A', 'D', ' ') {
						chr.RequestMethod = include.ChHttpHead

					}

				case 7:
					if lib.ChStr7Cmp(reqBdy, 'O', 'P', 'T', 'I', 'O', 'N', 'S', ' ') {
						chr.RequestMethod = include.ChHttpOption

					}
				}

			} else if (char == include.ChStrSpace || char == include.ChStrQry) && curr == 1 {
				/* try fetch request resource*/
				chr.RequestResource = reqBdy[prevPos+1 : pos]
				prevPos = pos
				if char == include.ChStrQry {
					curr++
				}

			} else if char == include.ChStrSpace && curr == 2 {
				/* try fetch query args */
				chr.RequestArgs = reqBdy[prevPos+1 : pos]
				curr++
				prevPos = pos

			} else if char == include.ChStrCr && reqBdy[pos+1] == include.ChStrLf {
				/* try fetch request http version */
				var ver = reqBdy[prevPos+1 : pos]
				if len(ver) == 8 {
					chr.RequestVersion = ver
				} else {
					chr.RequestVersion = include.HttpVerUn
				}
				curr = 10
				prevPos = pos + 2
			}

		} else if curr < 20 {
			/* parse request headers */
			if char == include.ChStrColon && reqBdy[pos+1] == include.ChStrSpace {
				/* header key */
				tmpHeaderKey = string(reqBdy[prevPos:pos])
				prevPos = pos + 2

			} else if char == include.ChStrCr && reqBdy[pos+1] == include.ChStrLf {
				/* header value */
				chr.RequestHeader[tmpHeaderKey] = reqBdy[prevPos:pos]

				prevPos = pos + 2
				if reqBdy[pos+2] == include.ChStrCr && reqBdy[pos+3] == include.ChStrLf {
					/* request body */
					chr.RequestBody = reqBdy[pos+4:]
					break
				}
			}
		}
	}

	out = ChainxRSM(chr)
	/* gen server control */
	if v, ok := chr.RequestHeader["Connection"]; ok && lib.ChStrCmp(v, include.HttpClose) {
		ctl["conn"] = "close"
	}
	return
}
