package main

import (
	"flag"
	"fmt"
)

func isMatch(s string, p string) bool {
	slen, plen := len(s), len(p)
	scur, pcur, sstar, pstar := 0, 0, -1, -1

	for scur < slen {
		if pcur < plen && (s[scur] == p[pcur] || p[pcur] == '?') {
			scur++
			pcur++
		} else if pcur < plen && p[pcur] == '*' {
			pstar = pcur
			pcur++
			sstar = scur
		} else if pstar >= 0 {
			pcur = pstar + 1
			sstar++
			scur = sstar
		} else {
			return false
		}
	}

	for pcur < plen && p[pcur] == '*' {
		pcur++
	}

	return pcur == plen
}

func isMatch2(s, p string) bool {
	slen, plen := len(s), len(p)
	dp := make([][]bool, slen+1)
	for i := 0; i < slen+1; i++ {
		dp[i] = make([]bool, plen+1)
	}
	dp[0][0] = true
	for i := 1; i <= plen; i++ {
		if p[i-1] == '*' {
			dp[0][i] = dp[0][i-1]
		}
	}

	for i := 1; i <= slen; i++ {
		for j := 1; j <= plen; j++ {
			if p[j-1] == '*' {
				dp[i][j] = dp[i-1][j] || dp[i][j-1]
			} else {
				dp[i][j] = (s[i-1] == p[j-1] || p[j-1] == '?') && dp[i-1][j-1]
			}
		}
	}

	return dp[slen][plen]
}

var s = flag.String("s", "", "specify source string")
var p = flag.String("p", "", "specify pattern string")

func main() {
	flag.Parse()
	fmt.Printf("'%s' matches pattern '%s' : %v\n", *s, *p, isMatch(*s, *p))
}
