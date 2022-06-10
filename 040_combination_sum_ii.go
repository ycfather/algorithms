package main

import (
	"fmt"
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	var path []int
	var result [][]int

	sort.Ints(candidates)
	backtrace(candidates, target, 0, path, &result)

	return result
}

func backtrace(candidates []int, target int, start int, path []int, result *[][]int) {
	if target == 0 {
		tmp := make([]int, len(path))
		copy(tmp, path)
		*result = append(*result, tmp)

		return
	}

	for i := start; i < len(candidates) && target-candidates[i] >= 0; i++ {
		if i > start && candidates[i] == candidates[i-1] {
			continue
		}

		path = append(path, candidates[i])
		backtrace(candidates, target-candidates[i], i+1, path, result)
		path = path[:len(path)-1]
	}
}

func main() {
	result := combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8)
	fmt.Printf("result: %+v\n", result)
}
