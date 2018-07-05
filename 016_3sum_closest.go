package main

import (
	"fmt"
	"sort"
)

type Ints []int

func (i Ints) Len() int {
	return len(i)
}

func (i Ints) Swap(a, b int) {
	i[a], i[b] = i[b], i[a]
}

func (i Ints) Less(a, b int) bool {
	return i[a] < i[b]
}

func abs(i int) int {
	if i < 0 {
		return -i
	}

	return i
}

func threeSumClosest(nums []int, target int) int {
	sort.Sort(Ints(nums))
	size := len(nums)
	closest := nums[0] + nums[1] + nums[2]
	for i := range nums {
		l, r := i+1, size-1
		for l < r {
			value := nums[i] + nums[l] + nums[r]
			if value == target {
				return value
			} else if value < target {
				l++
			} else {
				r--
			}

			if abs(target-value) < abs(target-closest) {
				closest = value
			}
		}
	}

	return closest
}

func main() {
	nums := []int{-1, 2, 1, -4}
	fmt.Printf("%v : %v\n", nums, threeSumClosest(nums, 1))
}
