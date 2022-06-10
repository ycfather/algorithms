package main

func minCostClimbingStairs(cost []int) int {
	sz := len(cost)
	dp := make([]int, sz)
	dp[0] = cost[0]
	dp[1] = cost[1]

	for i := 2; i < sz; i++ {
		if dp[i-2] < dp[i-1] {
			dp[i] = dp[i-2] + cost[i]
		} else {
			dp[i] = dp[i-1] + cost[i]
		}
	}

	if dp[sz-2] < dp[sz-1] {
		return dp[sz-2]
	}

	return dp[sz-1]
}
