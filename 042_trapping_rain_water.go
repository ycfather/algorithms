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

func trap(height []int) int {
	if len(height) < 3 {
		return 0
	}

	area := 0
	left := 0
	right := len(height) - 1
	curMaxHeight := 0
	for left <= right {
		if height[left] < height[right] {
			curMaxHeight = max(height[left], curMaxHeight)
			area += curMaxHeight - height[left]
			left++
		} else {
			curMaxHeight = max(height[right], curMaxHeight)
			area += curMaxHeight - height[right]
			right--
		}
	}

	return area
}

func main() {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Printf("max rain water of %+v : %d\n", height, trap(height))
}
