package main

import (
	"fmt"
	"strconv"
	"strings"
)

func restoreIpAddresses(s string) []string {
	var path, result []string
	var backtracking func(start int)
	backtracking = func(start int) {
		if len(path) == 4 {
			result = append(result, strings.Join(path, "."))
			return
		}

		for i := start; i < len(s); i++ {
			if len(s)-1-i > 3*(3-len(path)) {
				continue
			}

			if isValidIpSec(s, start, i) {
				path = append(path, s[start:i+1])
				backtracking(i + 1)
				path = path[:len(path)-1]
			}
		}
	}
	backtracking(0)

	return result
}

func isValidIpSec(s string, start, end int) bool {
	if end > start && s[start] == '0' {
		return false
	}

	sec, _ := strconv.Atoi(s[start : end+1])
	if sec > 255 {
		return false
	}

	return true
}

func main() {
	result := restoreIpAddresses("101023")
	fmt.Printf("result: %+v\n", result)
}
