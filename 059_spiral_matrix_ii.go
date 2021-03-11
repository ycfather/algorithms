package main

import (
	"fmt"
)

func generateMatrix(n int) [][]int {
	ret := make([][]int, n)
	for i := 0; i < n; i++ {
		ret[i] = make([]int, n)
	}

	seq := 1
	max := n * n
	left := 0
	right := n - 1
	top := 0
	bottom := n - 1

	for seq <= max {
		for i := left; i <= right; i++ {
			ret[top][i] = seq
			seq++
		}
		top++
		for i := top; i <= bottom; i++ {
			ret[i][right] = seq
			seq++
		}
		right--
		for i := right; i >= left; i-- {
			ret[bottom][i] = seq
			seq++
		}
		bottom--
		for i := bottom; i >= top; i-- {
			ret[i][left] = seq
			seq++
		}
		left++
	}

	return ret
}

func main() {
	fmt.Printf("Spiral Matrix: %+v\n", generateMatrix(4))
}
