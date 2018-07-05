package main

import (
	"flag"
	"fmt"
)

func isPalindrome(x int) bool {
	if x < 0 || x > 0 && x%10 == 0 {
		return false
	}

	var reverted int32
	i := x
	for i != 0 {
		t := reverted*10 + int32(i)%10
		if t/10 != reverted {
			reverted = 0
			break
		}

		reverted = t
		i /= 10
	}

	fmt.Printf("original : %d, reverted : %d\n", x, reverted)

	return int(reverted) == x
}

var x = flag.Int("n", 0, "specify the number")

func main() {
	flag.Parse()
	fmt.Printf("%d is palindrome number? %v\n", *x, isPalindrome(*x))
}
