package main

import (
	"fmt"
)

func max(a, b int) int {
	if a < b {
		return b
	}

	return a
}

func min(a, b int) int {
	if a > b {
		return b
	}

	return a
}

func maxArea(height []int) int {
	maxArea := 0
	left := 0
	right := len(height) - 1

	for left < right {
		maxArea = max(maxArea, min(height[left], height[right])*(right-left))

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return maxArea
}

func main() {
	height := []int{2, 3, 2, 4, 1}
	fmt.Printf("max area : %d\n", maxArea(height))
}
