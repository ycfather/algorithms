package main

import (
	"fmt"
)

func combinationSum3(k int, n int) [][]int {
	var path []int
	var result [][]int
	backtrace(n, k, 1, path, &result)

	return result
}

func backtrace(n, k, index int, path []int, result *[][]int) {
	if len(path) == k {
		sum := 0
		for _, v := range path {
			sum += v
		}

		if sum == n {
			tmp := make([]int, k)
			copy(tmp, path)
			*result = append(*result, tmp)
		}

		return
	}

	for i := index; i <= 9-(k-len(path))+1; i++ {
		path = append(path, i)
		backtrace(n, k, i+1, path, result)
		path = path[:len(path)-1]
	}
}

func main() {
	result := combinationSum3(3, 9)
	fmt.Printf("result: %+v\n", result)
}
