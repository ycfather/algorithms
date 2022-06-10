package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	ret := [][]int{}
	size := len(nums)
	if size < 3 {
		return ret
	}

	sort.Sort(Ints(nums))
	dif, l, r := 0, 0, 0
	for i := 0; i < size-2; i++ {
		if nums[i] > 0 {
			break
		}

		if i > 0 && nums[i-1] == nums[i] {
			continue
		}

		l = i + 1
		r = size - 1
		dif = 0 - nums[i]
		for l < r {
			if nums[l]+nums[r] == dif {
				ret = append(ret, []int{nums[i], nums[l], nums[r]})
				for l < r && nums[l+1] == nums[l] {
					l += 1
				}
				for l < r && nums[r-1] == nums[r] {
					r -= 1
				}
				l += 1
				r -= 1
			} else if nums[l]+nums[r] < dif {
				l += 1
			} else {
				r -= 1
			}
		}
	}

	return ret
}

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Printf("%v : %v\n", nums, threeSum(nums))
}
