package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	ret := make([][]int, 0, len(intervals))
	cur := intervals[0]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] > cur[1] {
			ret = append(ret, cur)
			cur = intervals[i]
		} else if intervals[i][1] > cur[1] {
			cur[1] = intervals[i][1]
		}
	}
	ret = append(ret, cur)

	return ret
}

func main() {
	intervals := [][]int{[]int{1, 3}, []int{2, 6}, []int{8, 10}, []int{15, 18}}
	fmt.Printf("Merged: %+v\n", merge(intervals))
}
