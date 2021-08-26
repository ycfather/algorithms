package main

import (
	"fmt"
)

func lengthOfLISByDP(nums []int) int {
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

func lengthOfLIS(nums []int) int {
	piles := 0
	n := len(nums)
	top := make([]int, n)

	for i := 0; i < n; i++ {
		left, right := 0, piles
		for left < right {
			mid := (left + right) / 2
			if nums[i] <= top[mid] {
				right = mid
			} else {
				left = mid + 1
			}
		}

		if left == piles {
			piles++
		}

		top[left] = nums[i]
	}

	return piles
}

func main() {
	nums := []int{7, 7, 7, 7, 7, 7, 7}
	fmt.Printf("res: %d\n", lengthOfLISByDP(nums))
}
