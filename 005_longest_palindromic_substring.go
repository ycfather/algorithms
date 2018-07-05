package main

import (
	"fmt"
	"strings"
)

func min(a, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}

func longestPalindrome(s string) string {
	b := []byte(s)
	nb := make([]byte, 2*len(b)+1)
	for i := range nb {
		if i%2 == 0 {
			nb[i] = '#'
		} else {
			nb[i] = b[i/2]
		}
	}

	RL := make([]int, len(nb))
	maxLen := 0
	maxLenC := 0

	maxRight := 0
	pos := 0

	for i := range nb {
		if i < maxRight {
			RL[i] = min(RL[2*pos-i], maxRight-i)
		} else {
			RL[i] = 1
		}
		//尝试扩展，注意处理边界
		for i-RL[i] >= 0 && i+RL[i] < len(nb) && nb[i-RL[i]] == nb[i+RL[i]] {
			RL[i] += 1
		}
		//更新maxRight,pos
		if RL[i]+i-1 > maxRight {
			maxRight = RL[i] + i - 1
			pos = i
		}
		// 更新最长回文串的长度
		if maxLen < RL[i] {
			maxLen = RL[i]
			maxLenC = i
		}
		//maxLen = max(maxLen, RL[i])
	}

	return strings.Replace(string(nb[maxLenC-maxLen+1:maxLenC+maxLen]), "#", "", -1)
}

func main() {
	// fmt.Printf("longest palindromic substring of 'babad' is '%s'\n", longestPalindrome("babad"))
	fmt.Printf("longest palindromic substring of 'forgeeksskeegfor' is '%s'\n", longestPalindrome("forgeeksskeegfor"))
}
