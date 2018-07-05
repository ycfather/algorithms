package main

import (
	"fmt"
)

func searchRange(nums []int, target int) []int {
	low := 0
	high := len(nums) - 1

	idx := -1
	for low <= high {
		mid := low + (high-low)/2
		if nums[mid] == target {
			idx = mid
			break
		}

		if target < nums[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	fmt.Printf("target index : %d\n", idx)

	size := len(nums)
	floor, ceiling := idx, idx
	for ceiling < size {
		if ceiling+1 < size && nums[ceiling+1] == target {
			ceiling++
		} else {
			break
		}
	}

	for floor >= 0 {
		if floor-1 >= 0 && nums[floor-1] == target {
			floor--
		} else {
			break
		}
	}

	return []int{floor, ceiling}
}

func main() {
	nums := []int{5, 7, 7, 8, 8, 10}
	fmt.Printf("range : %+v\n", searchRange(nums, 6))
}
