package main

import (
	"flag"
	"fmt"
)

func isMatch(s string, p string) bool {
	m := len(s)
	n := len(p)
	dp := make([][]bool, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true

	for i := 0; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == '*' {
				dp[i][j] = dp[i][j-2] || (i > 0 && (s[i-1] == p[j-2] || p[j-2] == '.') && dp[i-1][j])
			} else {
				dp[i][j] = i > 0 && dp[i-1][j-1] && (s[i-1] == p[j-1] || p[j-1] == '.')
			}
		}
	}

	return dp[m][n]
}

var s = flag.String("s", "", "source string")
var p = flag.String("p", "", "pattern")

func main() {
	flag.Parse()
	fmt.Printf("[source string : %s, pattern : %s] is matched? %v\n", *s, *p, isMatch(*s, *p))
}
