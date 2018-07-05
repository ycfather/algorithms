package main

import (
	"fmt"
	"unsafe"
)

func singleNumber(nums []int) int {
	ret := 0
	i := uint(0)
	bits := uint(8 * unsafe.Sizeof(ret))
	for ; i < bits; i++ {
		count := 0 // number of 1 in the i'th bit
		for _, num := range nums {
			count += (num >> i) & 1
		}
		ret += (count % 3) << i
	}

	return ret
}

func main() {
	nums := []int{-2, -2, 1, 1, -3, 1, -3, -3, -4, -2}
	fmt.Printf("the single number : %d\n", singleNumber(nums))
}
