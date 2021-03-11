package main

import (
	"fmt"
)

func plusOne(digits []int) []int {
	sz := len(digits)
	carry := (digits[sz-1] + 1) / 10
	digits[sz-1] = (digits[sz-1] + 1) % 10
	for i := sz - 2; i >= 0; i-- {
		s := digits[i] + carry
		digits[i] = s % 10
		carry = s / 10
	}

	if carry > 0 {
		digits = append([]int{carry}, digits...)
	}

	return digits
}

func main() {
	digits := []int{8, 9, 9, 9}
	fmt.Printf("result: %v\n", plusOne(digits))
}
