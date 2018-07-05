package main

import (
	"fmt"
)

func nextPermutation(nums []int) {
	size := len(nums)
	if size <= 1 {
		return
	}

	k := size - 1 // from end to start, the last number which is bigger than its previous number
	for k > 0 && nums[k] <= nums[k-1] {
		k--
	}

	if k > 0 {
		i := size - 1
		for ; i >= k; i-- {
			if nums[i] > nums[k-1] {
				break
			}
		}
		nums[k-1], nums[i] = nums[i], nums[k-1]
	}

	r := (k + size - 1) / 2
	for m := k; m <= r; m++ {
		mirror := size - 1 + k - m
		nums[m], nums[mirror] = nums[mirror], nums[m]
	}
}

func main() {
	nums := []int{1, 2, 7, 4, 3, 1}
	fmt.Printf("original : %+v\n", nums)
	nextPermutation(nums)
	fmt.Printf("after : %+v\n", nums)

}
