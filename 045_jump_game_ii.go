package main

import (
	"fmt"
)

func max(a, b int) int {
	if a >= b {
		return a
	}

	return b
}

func jump(nums []int) int {
	reach := 0
	last := 0
	step := 0

	for i, num := range nums {
		if i > reach {
			break
		}

		if i > last {
			step++
			last = reach
		}
		reach = max(reach, num+i)
	}

	return step
}

func main() {
	nums := []int{2, 3, 1, 1, 4}
	fmt.Printf("minimum jump steps of %+v : %d\n", nums, jump(nums))
}
