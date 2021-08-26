package main

import (
	"fmt"
	"math"
)

const INVALID = 10001

var memo [][]int

func minFallingPathSum(matrix [][]int) int {
	n := len(matrix)
	memo = make([][]int, n)
	for i := 0; i < n; i++ {
		memo[i] = make([]int, n)
		for j := 0; j < n; j++ {
			memo[i][j] = INVALID
		}
	}

	res := math.MaxInt
	for i := 0; i < n; i++ {
		t := dp(matrix, n-1, i)
		if res > t {
			res = t
		}
	}

	return res
}

func dp(matrix [][]int, i, j int) int {
	if i < 0 || i >= len(matrix) || j < 0 || j >= len(matrix[0]) {
		return INVALID
	}

	if i == 0 {
		return matrix[i][j]
	}

	if memo[i][j] != INVALID {
		return memo[i][j]
	}

	memo[i][j] = matrix[i][j] + min(dp(matrix, i-1, j-1), dp(matrix, i-1, j), dp(matrix, i-1, j+1))

	return memo[i][j]
}

func min(a, b, c int) int {
	res := a
	if res > b {
		res = b
	}

	if res > c {
		res = c
	}

	return res
}

func main() {
	matrix := [][]int{[]int{2, 1, 3}, []int{6, 5, 4}, []int{7, 8, 9}}
	mfps := minFallingPathSum(matrix)

	fmt.Printf("mfps: %d\n %+v\n", mfps, memo)
}
