package main

import (
	"flag"
	"fmt"
)

func longestValidParentheses(s string) int {
	max := 0
	size := len(s)
	if size == 0 {
		return max
	}

	dp := make([]int, size)
	dp[size-1] = 0
	for i := size - 2; i >= 0; i-- {
		if s[i] == '(' {
			j := i + 1 + dp[i+1]
			if j < size && s[j] == ')' {
				dp[i] = dp[i+1] + 2
				if j+1 < size {
					dp[i] += dp[j+1]
				}
			}

			if max < dp[i] {
				max = dp[i]
			}
		}
	}

	return max
}

var s = flag.String("s", "", "specify parentheses string")

func main() {
	flag.Parse()
	fmt.Printf("longest valid parentheses of '%s' : %d\n", *s, longestValidParentheses(*s))
}
