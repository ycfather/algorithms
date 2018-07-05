package main

import (
	"flag"
	"fmt"
)

var r = flag.String("r", "XCIX", "specify the roman number")

var m = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	ret := 0
	size := len(s)
	for i := range s {
		if i+1 >= size || m[s[i+1]] <= m[s[i]] {
			ret += m[s[i]]
		} else {
			ret -= m[s[i]]
		}
	}

	return ret
}

func main() {
	flag.Parse()
	fmt.Printf("'%s' to integer : %d\n", *r, romanToInt(*r))
}
