package main

import (
	"fmt"
	"sort"
)

func permuteUnique(nums []int) [][]int {
	var res [][]int
	used := make([]bool, len(nums))
	path := make([]int, 0)
	fmt.Printf("original array: %v\n", nums)
	sort.Ints(nums)
	fmt.Printf("sorted array: %v\n", nums)
	dfs(nums, len(nums), 0, &path, used, &res)

	return res
}

func dfs(nums []int, length int, depth int, path *[]int, used []bool, res *[][]int) {
	if depth == length {
		fmt.Printf("depth: %d, %v\n", depth, *path)
		buf := make([]int, len(*path))
		copy(buf, *path)
		*res = append(*res, buf)
		return
	}

	for i := 0; i < length; i++ {
		if used[i] || i > 0 && nums[i] == nums[i-1] && !used[i-1] {
			continue
		}

		*path = append(*path, nums[i])
		used[i] = true
		dfs(nums, length, depth+1, path, used, res)
		used[i] = false
		*path = append((*path)[:len(*path)-1])
	}
}

func main() {
	nums := []int{1, 2, 1, 1}
	res := permuteUnique(nums)
	fmt.Printf("permutations: %+v\n", res)
}
