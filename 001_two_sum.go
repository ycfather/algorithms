package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	size := len(nums)
	for i := 0; i < size-1; i++ {
		for j := i + 1; j < size; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return []int{-1, -1}
}

func twoSum2(nums []int, target int) []int {
	size := len(nums)
	table := make(map[int]int, size) // key : number, value : index
	for i, num := range nums {
		if di, ok := table[target-num]; ok {
			return []int{i, di}
		} else {
			table[num] = i
		}
	}

	return []int{-1, -1}
}

func main() {
	nums := []int{2, 7, 11, 5}
	fmt.Printf("%+v\n", twoSum(nums, 14))
}
