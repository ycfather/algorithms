package main

import (
	"fmt"
)

func removeElement(nums []int, val int) int {
	idx := 0
	for _, num := range nums {
		if num == val {
			continue
		} else {
			nums[idx] = num
			idx++
		}
	}

	return idx
}

func main() {
	nums := []int{3, 2, 2, 3}
	sz := removeElement(nums, 3)
	fmt.Printf("after : %+v, size : %d\n", nums, sz)
}
