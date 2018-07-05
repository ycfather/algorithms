package main

import (
	"fmt"
)

func searchInsert(nums []int, target int) int {
	size := len(nums)
	if size == 0 {
		return 0
	}

	low := 0
	high := size - 1
	for low <= high {
		mid := low + (high-low)/2
		if nums[mid] == target {
			return mid
		}

		if target < nums[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	for i, v := range nums {
		if v >= target {
			return i
		}
	}

	return size
}

func main() {
	nums := []int{1, 3, 5, 6}
	fmt.Printf("index : %d\n", searchInsert(nums, 5))
}
