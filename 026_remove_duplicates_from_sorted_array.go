package main

import (
	"fmt"
)

func removeDuplicates(nums []int) int {
	idx := 0
	for i := 0; i < len(nums); {
		j := i + 1
		for ; j < len(nums); j++ {
			if nums[j] == nums[i] {
				continue
			} else {
				break
			}
		}
		nums[idx] = nums[i]
		idx++
		i = j
	}

	return idx
}

func main() {
	nums := []int{1, 1, 1, 1, 1, 1}
	sz := removeDuplicates(nums)
	fmt.Printf("after : %+v, size : %d\n", nums, sz)
}
