package main

import (
	"fmt"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func insert(intervals [][]int, newInterval []int) [][]int {
	ret := make([][]int, 0, len(intervals)+1)
	last := -1
	breakIndex := -1
	for i, interval := range intervals {
		if newInterval[1] < interval[0] {
			breakIndex = i
			break
		} else if newInterval[0] > interval[1] {
			last = i
		} else {
			newInterval[0] = min(interval[0], newInterval[0])
			newInterval[1] = max(interval[1], newInterval[1])
		}
	}
	ret = append(ret, intervals[:(last+1)]...)
	ret = append(ret, newInterval)
	if breakIndex > -1 {
		ret = append(ret, intervals[breakIndex:]...)
	}

	return ret
}

func main() {
	intervals := [][]int{[]int{1, 2}, []int{3, 5}, []int{6, 7}, []int{8, 10}, []int{12, 16}}
	fmt.Printf("Merged: %+v\n", insert(intervals, []int{4, 8}))

	intervals = [][]int{[]int{3, 5}, []int{12, 15}}
	fmt.Printf("Merged: %+v\n", insert(intervals, []int{6, 6}))

	intervals = [][]int{[]int{1, 5}}
	fmt.Printf("Merged: %+v\n", insert(intervals, []int{2, 3}))
}
