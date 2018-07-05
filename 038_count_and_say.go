package main

import (
	"flag"
	"fmt"
	"strconv"
)

func countAndSay(n int) string {
	if n == 0 {
		return ""
	}

	dp := make([]string, n)
	dp[0] = "1"
	for i := 1; i < n; i++ {
		nstr := ""
		num := 0
		cd := '.'
		for _, d := range dp[i-1] {
			if d == cd {
				num++
				continue
			}

			if num > 0 {
				nstr = nstr + strconv.Itoa(num) + string(cd)
			}
			cd = d
			num = 1
		}
		nstr = nstr + strconv.Itoa(num) + string(cd)
		dp[i] = nstr
	}

	return dp[n-1]
}

var n = flag.Int("n", 1, "specify the number")

func main() {
	flag.Parse()
	fmt.Printf("'count and say' of %d : %s\n", *n, countAndSay(*n))
}
