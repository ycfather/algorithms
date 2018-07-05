package main

import (
	"flag"
	"fmt"
)

func divide(dividend int, divisor int) int {
	sign := 1
	if dividend > 0 && divisor < 0 || dividend < 0 && divisor > 0 {
		sign = -1
	}

	lDividend := int64(dividend)
	if lDividend < 0 {
		lDividend = 0 - lDividend
	}

	lDivisor := int64(divisor)
	if lDivisor < 0 {
		lDivisor = 0 - lDivisor
	}

	quotient := int64(0)
	for lDividend >= lDivisor {
		tDivisor := lDivisor
		i := int64(1)
		for lDividend >= tDivisor {
			lDividend -= tDivisor
			quotient += i
			i <<= 1
			tDivisor <<= 1
		}
	}

	if sign == -1 {
		quotient = -quotient
	}

	MAX_INT32 := int64(1<<31 - 1)
	MIN_INT32 := int64(-1 << 31)
	if quotient < MIN_INT32 {
		quotient = MIN_INT32
	}

	if quotient > MAX_INT32 {
		quotient = MAX_INT32
	}
	return int(quotient)
}

var dividend = flag.Int("e", 1, "specify dividend")
var divisor = flag.Int("s", 1, "specify divisor")

func main() {
	flag.Parse()
	fmt.Printf("%d/%d = %d\n", *dividend, *divisor, divide(*dividend, *divisor))
}
