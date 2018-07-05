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

func fourSum(nums []int, target int) [][]int {
	ret := [][]int{}
	size := len(nums)
	if size < 4 {
		return ret
	}

	sort.Sort(Ints(nums))

	for i := 0; i < size-3; i++ {
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}

		for j := i + 1; j < size-2; j++ {
			if j > i+1 && nums[j-1] == nums[j] {
				continue
			}

			l := j + 1
			r := size - 1
			for l < r {
				sum := nums[i] + nums[j] + nums[l] + nums[r]
				if sum == target {
					ret = append(ret, []int{nums[i], nums[j], nums[l], nums[r]})
					for l < r && nums[l+1] == nums[l] {
						l++
					}
					for l < r && nums[r-1] == nums[r] {
						r--
					}
					l++
					r--
				} else if sum < target {
					l++
				} else {
					r--
				}
			}
		}
	}

	return ret
}

func main() {
	nums := []int{1, 0, -1, 0, -2, 2}
	fmt.Printf("%v : %v\n", nums, fourSum(nums, 0))
}
