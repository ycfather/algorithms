package main

import (
	"fmt"
)

func lengthOfLastWord(s string) int {
	length := 0
	flag := false
	for idx := len(s) - 1; idx >= 0; idx-- {
		if s[idx] == ' ' {
			if !flag {
				continue
			} else {
				break
			}
		} else if !flag {
			flag = true
		}
		length++
	}

	return length
}

func main() {
	s := ""
	fmt.Printf("Length Of Last Word (%s): %d\n", s, lengthOfLastWord(s))
}
