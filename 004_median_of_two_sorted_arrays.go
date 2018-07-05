package main

import (
	"fmt"
	"math"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	N1 := len(nums1)
	N2 := len(nums2)

	if N1 < N2 {
		return findMedianSortedArrays(nums2, nums1)
	}

	lo := 0
	hi := N2 * 2
	for lo <= hi {
		C2 := (lo + hi) / 2 // Try Cut 2
		C1 := N1 + N2 - C2  // Calculate Cut 1 accordingly

		var L1, R1, L2, R2 int
		// Get L1, R1, L2, R2 respectively
		if C1 == 0 {
			L1 = math.MinInt32
		} else {
			L1 = nums1[(C1-1)/2]
		}

		if C2 == 0 {
			L2 = math.MinInt32
		} else {
			L2 = nums2[(C2-1)/2]
		}

		if C1 == N1*2 {
			R1 = math.MaxInt32
		} else {
			R1 = nums1[(C1)/2]
		}

		if C2 == N2*2 {
			R2 = math.MaxInt32
		} else {
			R2 = nums2[(C2)/2]
		}

		if L1 > R2 { // A1's lower half is too big; need to move C1 left (C2 right)
			lo = C2 + 1
		} else if L2 > R1 { // A2's lower half too big; need to move C2 left.
			hi = C2 - 1
		} else { // Otherwise, that's the right cut.
			return (math.Max(float64(L1), float64(L2)) + math.Min(float64(R1), float64(R2))) / 2
		}
	}
	return -1
}

func main() {
	nums1 := []int{6, 9, 13, 18}
	nums2 := []int{6, 9, 11, 13, 18}

	fmt.Printf("median of %+v and %+v is : %v\n", nums1, nums2, findMedianSortedArrays(nums1, nums2))
}
