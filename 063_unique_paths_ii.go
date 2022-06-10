package main

import (
	"fmt"
)

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	rows := len(obstacleGrid)
	columns := len(obstacleGrid[0])

	dp := make([][]int, rows)
	for row := 0; row < rows; row++ {
		dp[row] = make([]int, columns)
	}

	for row := 0; row < rows && obstacleGrid[row][0] == 0; row++ {
		dp[row][0] = 1
	}

	for column := 0; column < columns && obstacleGrid[0][column] == 0; column++ {
		dp[0][column] = 1
	}

	for i := 1; i < rows; i++ {
		for j := 1; j < columns; j++ {
			if obstacleGrid[i][j] == 0 {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}

	return dp[rows-1][columns-1]
}

func main() {
	fmt.Printf("number of unique paths: %d\n", uniquePathsWithObstacles([][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}))
}
