package main

import (
	"flag"
	"fmt"
	"strconv"
)

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	if len(num1) < len(num2) { // make num2 to be the shorter one
		num1, num2 = num2, num1
	}

	len1 := len(num1)
	len2 := len(num2)
	table := make([][]int, len2)
	for i := 0; i < len2; i++ {
		table[i] = make([]int, len1+len2)
	}

	shift := 0
	for i := len2 - 1; i >= 0; i-- {
		d2 := num2[i] - '0'
		carry := 0
		for j := len1 - 1; j >= 0; j-- {
			d1 := num1[j] - '0'
			sum := int(d1)*int(d2) + carry
			table[i][len2-shift+j] = sum % 10
			carry = sum / 10
		}
		table[i][len2-shift-1] = carry
		shift++
	}

	result := make([]int, len1+len2)
	carry := 0
	for col := len1 + len2 - 1; col >= 0; col-- {
		sum := 0
		for row := len2 - 1; row >= 0; row-- {
			sum += table[row][col]
		}
		sum += carry
		result[col] = sum % 10
		carry = sum / 10
	}

	str := ""
	for i, v := range result {
		if i == 0 && v == 0 {
			continue
		}
		str += strconv.Itoa(v)
	}
	return str
}

var num1 = flag.String("n1", "1", "specify num1")
var num2 = flag.String("n2", "1", "specify num2")

func main() {
	flag.Parse()
	fmt.Printf("%s * %s = %s\n", *num1, *num2, multiply(*num1, *num2))
}
