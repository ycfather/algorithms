package main

import (
	"fmt"
	"math"
)

func myAtoi(str string) int {
	if len(str) == 0 {
		return 0
	}

	res := 0
	sign := 1
	index := 0 // first non-empty char

	bytes := []byte(str)
	for i, b := range bytes {
		if b == ' ' {
			continue
		} else {
			index = i
			break
		}
	}

	if bytes[index] == '-' {
		sign = -1
		index++
	} else if bytes[index] == '+' {
		sign = 1
		index++
	}

	for i := index; i < len(bytes); i++ {
		digit := bytes[i] - '0'
		if digit < 0 || digit > 9 { // non-digit
			break
		}

		if math.MaxInt32/10 < res || math.MaxInt32/10 == res && math.MaxInt32%10 < digit {
			if sign == 1 {
				return math.MaxInt32
			} else {
				return math.MinInt32
			}
		}

		res = res*10 + int(digit)
	}

	return res * sign
}

func main() {
	str := "42"
	fmt.Printf("atoi(%s) : %d\n", str, myAtoi(str))

	str = "  -42"
	fmt.Printf("atoi(%s) : %d\n", str, myAtoi(str))

	str = "4193 with words"
	fmt.Printf("atoi(%s) : %d\n", str, myAtoi(str))

	str = "words and 987"
	fmt.Printf("atoi(%s) : %d\n", str, myAtoi(str))

	str = "-91283472332"
	fmt.Printf("atoi(%s) : %d\n", str, myAtoi(str))
}
