package main

import (
	"fmt"
	"sort"
)

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)

	res := 0
	for i, j := len(s)-1, len(g)-1; i >= 0 && j >= 0; {
		if s[i] >= g[j] {
			res++
			i--
		}
		j--
	}

	return res
}

func main() {
	fmt.Printf("result: %d\n", findContentChildren([]int{1, 2}, []int{1, 2, 3}))
}
