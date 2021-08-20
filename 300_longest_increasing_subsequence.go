package main

import (
	"fmt"
)

func lengthOfLIS(nums []int) int {
	sz := len(nums)
	dp := make([]int, sz)
	for i := 0; i < sz; i++ {
		dp[i] = 1
	}

	res := 1
	for i := 1; i < sz; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] && dp[i] < dp[j]+1 {
				dp[i] = dp[j] + 1
			}
		}
		if dp[i] > res {
			res = dp[i]
		}
	}

	return res
}

func main() {
	nums := []int{7, 7, 7, 7, 7, 7, 7}
	fmt.Printf("res: %d\n", lengthOfLIS(nums))
}
