package main

import (
	"fmt"
)

func largestRectangleArea(heights []int) int {
	heights = append([]int{0}, heights...)
	heights = append(heights, 0)
	fmt.Printf("heights: %v\n", heights)

	maxArea := 0
	stack := make([]int, 0)

	for i, h := range heights {
		for len(stack) > 0 && heights[stack[len(stack)-1]] > h {
			cur := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			left := stack[len(stack)-1] + 1
			right := i - 1
			if area := (right - left + 1) * heights[cur]; area > maxArea {
				fmt.Printf("area: %d\n", area)
				maxArea = area
			}
		}
		stack = append(stack, i)
	}

	return maxArea
}

func main() {
	heights := []int{2, 1, 5, 6, 2, 3}
	fmt.Printf("largest rectangle area: %d\n", largestRectangleArea(heights))
}
