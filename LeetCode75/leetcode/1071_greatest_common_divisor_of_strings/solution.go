package main

import "fmt"

func main() {
	fmt.Println(gcdOfStrings("LEET", "CODE"))
}

func gcdOfStrings(str1 string, str2 string) string {
	var s, t, strBuild string
	var slen, tlen, exp int

	if len(str1) > len(str2) {
		s = str1
		t = str2
		slen = len(str1)
		tlen = len(str2)
	} else {
		s = str2
		t = str1
		slen = len(str2)
		tlen = len(str1)
	}
	if slen == tlen && s == t {
		return t
	}

	for {
		fmt.Printf("tlen %v", tlen)
		strBuild = ""

		if slen%tlen == 0 {
			exp = slen / tlen
			for i := 0; i < exp; i++ {
				strBuild += t
			}
			if strBuild == s {
				return t
			} else if tlen == 1 {
				return ""
			} else {
				tlen = tlen / 2
				t = t[:tlen]

			}
		} else if tlen == 1 {
			return ""
		} else {
			tlen = tlen / 2
			t = t[:tlen]
		}
	}
}
