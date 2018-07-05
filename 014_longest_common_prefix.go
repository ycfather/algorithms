package main

import (
	"fmt"
)

func longestCommonPrefix(strs []string) string {
	size := len(strs)
	if size == 0 {
		return ""
	}

	if size == 1 {
		return strs[0]
	}

	minLen := 1<<31 - 1 // MaxInt32
	for _, str := range strs {
		if minLen > len(str) {
			minLen = len(str)
		}
	}

	pos := 0
	for i := 0; i < size-1; i++ {
		for j := 0; j < minLen; j++ {
			if strs[i+1][j] != strs[i][j] {
				if j == 0 {
					return ""
				}
				if pos == 0 || pos > j {
					pos = j
				}
				break
			}

			if j == minLen-1 {
				if pos == 0 || pos > minLen {
					pos = minLen
				}
				break
			}
		}
	}

	return strs[0][0:pos]
}

func main() {
	strs := []string{"abd", "abc", "abcd"}
	fmt.Printf("LCP of %v : %s\n", strs, longestCommonPrefix(strs))
}
