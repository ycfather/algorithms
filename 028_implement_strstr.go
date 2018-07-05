package main

import (
	"flag"
	"fmt"
)

func strStr(haystack string, needle string) int {
	hSize := len(haystack)
	nSize := len(needle)

	if nSize == 0 {
		return 0
	}

	if hSize < nSize {
		return -1
	}

	idx := -1
	for i := 0; i < hSize; i++ {
		if haystack[i] != needle[0] {
			continue
		}

		if hSize-i < nSize {
			break
		}

		matched := true
		for j := 1; j < nSize; j++ {
			if haystack[i+j] != needle[j] {
				matched = false
				break
			}
		}

		if matched {
			idx = i
			break
		}
	}

	return idx
}

var haystack = flag.String("h", "", "specify haystack")
var needle = flag.String("n", "", "specify needle")

func main() {
	flag.Parse()
	fmt.Printf("strstr('%s', '%s') : %d\n", *haystack, *needle, strStr(*haystack, *needle))
}
