package main

import (
	"fmt"
)

func intToRoman(num int) string {
	if num <= 0 {
		return ""
	}

	ret := ""
	nums := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	romans := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	for i := 0; i < len(nums) && num > 0; i++ {
		if num < nums[i] {
			continue
		}

		for num >= nums[i] {
			num -= nums[i]
			ret += romans[i]
		}
	}

	return ret
}

func main() {
	fmt.Printf("roman of 1994 : %s\n", intToRoman(1994))
}
