package main

import (
	"fmt"
)

func longestCommonSubsequence(text1 string, text2 string) int {
	m := len(text1)
	n := len(text2)

	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = 1 + dp[i-1][j-1]
			} else {
				dp[i][j] = max(dp[i][j-1], dp[i-1][j])
			}
		}
	}

	return dp[m][n]
}

func max(a, b int) int {
	res := a
	if res < b {
		res = b
	}

	return res
}

func main() {
	text1 := "abcde"
	text2 := "ace"

	fmt.Printf("lcs: %d\n", longestCommonSubsequence(text1, text2))
}
