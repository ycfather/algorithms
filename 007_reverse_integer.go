package main

import (
	"fmt"
)

func reverse(x int) int {
	var res int32
	for x != 0 {
		t := res*10 + int32(x)%10
		if t/10 != res {
			return 0
		}

		res = t
		x /= 10
	}

	return int(res)
}

func main() {
	fmt.Printf("reverse integer : %d\n", reverse(1534236469))
	fmt.Printf("reverse integer : %d\n", reverse(-123))
	fmt.Printf("reverse integer : %d\n", reverse(120))
}
