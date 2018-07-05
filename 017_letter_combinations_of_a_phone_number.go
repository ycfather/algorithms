package main

import (
	"flag"
	"fmt"
)

var m = map[rune]string{
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	result := []string{""}
	for _, d := range digits {
		dstr := m[d]
		nResult := make([]string, len(dstr)*len(result))
		nRi := 0
		for _, rv := range result {
			for _, dv := range dstr {
				nResult[nRi] = rv + string(dv)
				nRi++
			}
		}
		result = nResult
	}

	return result
}

var s = flag.String("s", "23", "specify number")

func main() {
	flag.Parse()
	fmt.Printf("letter combinations of %s : %v\n", *s, letterCombinations(*s))
}
