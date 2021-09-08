package main

import (
	"fmt"
)

func maxSubArray(nums []int) int {
	sz := len(nums)

	pre := nums[0]
	res := pre
	for i := 1; i < sz; i++ {
		cur := max(pre+nums[i], nums[i])
		res = max(res, cur)
		pre = cur
	}

	return res
}

func max(a, b int) int {
	res := a
	if a < b {
		res = b
	}

	return res
}

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Printf("max: %d\n", maxSubArray(nums))
}
