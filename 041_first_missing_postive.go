package main

import (
	"fmt"
)

func firstMissingPositive(nums []int) int {
	size := len(nums)
	idx := 0
	for idx < size {
		if nums[idx] == idx+1 || nums[idx] <= 0 || nums[idx] > size || nums[idx] == nums[nums[idx]-1] {
			idx++
			continue
		}

		nums[idx], nums[nums[idx]-1] = nums[nums[idx]-1], nums[idx]
	}

	for i, num := range nums {
		if num != i+1 {
			return i + 1
		}
	}

	return size + 1
}

func main() {
	nums := []int{1, 2, 2, 1, 3, 1, 0, 4, 0}
	fmt.Printf("original : %+v\n", nums)
	missing := firstMissingPositive(nums)
	fmt.Printf("after : %+v\n", nums)
	fmt.Printf("the first missing postive of %+v : %d\n", nums, missing)
}
