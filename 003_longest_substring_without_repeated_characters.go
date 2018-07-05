package main

import (
	"flag"
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
	b := []byte(s)
	m := make(map[byte]int)

	maxLen := 0
	prev := 0
	for i, v := range b {
		if idx, ok := m[v]; ok {
			if maxLen < i-prev {
				maxLen = i - prev
			}

			for j := prev; j <= idx; j++ {
				delete(m, b[j])
			}

			// shift the left ending right
			prev = idx + 1
		}
		m[v] = i
	}

	size := len(b)
	if maxLen < size-prev {
		maxLen = size - prev
	}

	return maxLen
}

func lengthOfLongestSubstring2(s string) int {
	table := [128]int{} // index : character, value : the character's last index
	maxLen := 0
	prev := 0
	for i, v := range s {
		if prev < table[v] {
			prev = table[v]
		}

		if maxLen < i-prev+1 {
			maxLen = i - prev + 1
		}
		table[v] = i + 1
	}

	return maxLen
}

var s = flag.String("s", "", "specify the string")

func main() {
	flag.Parse()
	fmt.Printf("longest substring without repeated characters of '%s' : %d\n", *s, lengthOfLongestSubstring2(*s))
}
