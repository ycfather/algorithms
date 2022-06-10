package main

import (
	"fmt"
)

func subsets(nums []int) [][]int {
	var path []int
	var result [][]int

	var backtracking func(start int)
	backtracking = func(start int) {
		tmp := make([]int, len(path))
		copy(tmp, path)
		result = append(result, tmp)

		for i := start; i < len(nums); i++ {
			path = append(path, nums[i])
			backtracking(i + 1)
			path = path[:len(path)-1]
		}
	}

	backtracking(0)

	return result
}

func main() {
	result := subsets([]int{1, 2, 3})
	fmt.Printf("result: %+v\n", result)
}
