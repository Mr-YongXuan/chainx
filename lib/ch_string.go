package lib

import "github.com/Mr-YongXuan/chainx/include"

/* nginx */
func ChStr3Cmp(ua []byte, c0, c1, c2, c3 byte) (res bool) {
	return ua[0] == c0 && ua[1] == c1 && ua[2] == c2 && ua[3] == c3
}

func ChStr4Cmp(ua []byte, c0, c1, c2, c3, c4 byte) (res bool) {
	return ua[0] == c0 && ua[1] == c1 && ua[2] == c2 && ua[3] == c3 && ua[4] == c4
}

func ChStr7Cmp(ua []byte, c0, c1, c2, c3, c4, c5, c6, c7 byte) (res bool) {
	return ua[0] == c0 && ua[1] == c1 && ua[2] == c2 && ua[3] == c3 && ua[4] == c4 &&
		ua[5] == c5 && ua[6] == c6 && ua[7] == c7
}

func ChHttpVerCmp(req []byte, ver []byte) (res bool) {
	return req[0] == ver[0] && req[1] == ver[1] && req[2] == ver[2] && req[3] == ver[3] &&
		req[4] == ver[4] && req[5] == ver[5] && req[6] == ver[6] && req[7] == ver[7]
}

func ChMethodIsApprove(method int, routerMethods []int) (approve bool) {
	if method == include.ChHttpHead || method == include.ChHttpOption {
		return true
	}
	for _, routerMethod := range routerMethods {
		if method == routerMethod {
			return true
		}
	}
	return false
}

func ChStrCmp(s1 []byte, s2 []byte) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}
