package main

import (
	"fmt"
)

func combinationSum(candidates []int, target int) [][]int {
	var path []int
	var result [][]int
	backtracking(candidates, target, 0, path, &result)

	return result
}

func backtracking(candidates []int, target int, start int, path []int, result *[][]int) {
	if target == 0 {
		tmp := make([]int, len(path))
		copy(tmp, path)
		*result = append(*result, tmp)

		return
	}

	for i := start; i < len(candidates); i++ {
		if target-candidates[i] < 0 {
			continue
		}
		path = append(path, candidates[i])
		backtracking(candidates, target-candidates[i], i, path, result)
		path = path[:len(path)-1]
	}
}

func main() {
	result := combinationSum([]int{2, 7, 6, 3, 5, 1}, 9)
	fmt.Printf("result: %+v\n", result)
}
