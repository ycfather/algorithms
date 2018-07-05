package main

import (
	"fmt"
)

func search(nums []int, target int) int {
	low := 0
	high := len(nums) - 1

	for low <= high {
		mid := low + (high-low)/2
		if nums[mid] == target {
			return mid
		}

		if nums[low] <= nums[mid] {
			if nums[low] <= target && target < nums[mid] {
				high = mid - 1
			} else {
				low = mid + 1
			}
		} else {
			if target > nums[mid] && target <= nums[high] {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	}

	return -1
}

func main() {
	nums := []int{3, 1}
	fmt.Printf("target index : %d\n", search(nums, 1))
}
