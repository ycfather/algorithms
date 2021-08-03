package main

func canJump(nums []int) bool {
	k := 0
	for i, num := range nums {
		if i > k {
			return false
		}
		k = max(k, i+num)
	}

	return true
}

func max(a, b int) int {
	if a >= b {
		return a
	}

	return b
}
