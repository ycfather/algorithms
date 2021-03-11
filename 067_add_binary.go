package main

import (
	"fmt"
)

func addBinary(a string, b string) string {
	if len(a) > len(b) {
		delta := len(a) - len(b)
		for delta > 0 {
			b = "0" + b
			delta--
		}
	} else {
		delta := len(b) - len(a)
		for delta > 0 {
			a = "0" + a
			delta--
		}
	}

	// sz := len(a)
	fmt.Printf("%s + %s = \n", a, b)

	sz := len(a)
	res := make([]byte, sz)
	carry := byte(0)
	for i := sz - 1; i >= 0; i-- {
		s := (a[i] - '0') + (b[i] - '0') + carry
		res[i] = (s % 2) + '0'
		fmt.Println(res[i])
		carry = s / 2
	}

	if carry > 0 {
		res = append([]byte{carry + '0'}, res...)
	}

	return string(res)
}

func main() {
	fmt.Printf("binary sum: %s\n", addBinary("1010", "1011010"))
}
