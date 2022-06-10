package main

import (
	"fmt"
)

func combine(n int, k int) [][]int {
	res := [][]int{}
	if k <= 0 || n < k {
		return res
	}

	path := []int{}
	dfs(n, k, 1, path, &res)

	return res
}

func dfs(n int, k int, index int, path []int, res *[][]int) {
	if len(path) == k {
		tmp := make([]int, k)
		copy(tmp, path)
		*res = append(*res, tmp)
		return
	}

	for i := index; i <= n-(k-len(path))+1; i++ {
		path = append(path, i)
		dfs(n, k, i+1, path, res)
		path = path[:(len(path) - 1)]
	}
}

func main() {
	res := combine(4, 2)
	fmt.Printf("result: %+v\n", res)
}
