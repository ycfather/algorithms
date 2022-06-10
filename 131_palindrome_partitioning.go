package main

import (
	"fmt"
)

func partition(s string) [][]string {
	var path []string
	var result [][]string
	// backtrace(s, 0, path, &result)
	var backtracking func(start int)
	backtracking = func(start int) {
		if start >= len(s) {
			tmp := make([]string, len(path))
			copy(tmp, path)
			result = append(result, tmp)

			return
		}

		for i := start; i < len(s); i++ {
			if isPalindrome(s, start, i) {
				path = append(path, s[start:i+1])
			} else {
				continue
			}
			backtracking(i + 1)
			path = path[:len(path)-1]
		}
	}
	backtracking(0)

	return result
}

func backtrace(s string, start int, path []string, result *[][]string) {
	if start >= len(s) {
		tmp := make([]string, len(path))
		copy(tmp, path)

		*result = append(*result, tmp)

		return
	}

	for i := start; i < len(s); i++ {
		if isPalindrome(s, start, i) {
			path = append(path, s[start:i+1])
		} else {
			continue
		}
		backtrace(s, i+1, path, result)
		path = path[:len(path)-1]
	}
}

func isPalindrome(s string, start, end int) bool {
	for i, j := start, end; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

func main() {
	result := partition("aab")
	fmt.Printf("result: %+v\n", result)
}
