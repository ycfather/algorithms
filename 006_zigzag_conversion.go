package main

import (
	"bytes"
	"flag"
	"fmt"
)

func convert(s string, numRows int) string {
	if numRows <= 1 {
		return s
	}

	size := len(s)
	var zigzag bytes.Buffer
	delta := 2*numRows - 2
	for i := 0; i < numRows; i++ {
		for j := i; j < size; j += delta {
			zigzag.WriteByte(s[j])
			bias := j + delta - 2*i
			if i != 0 && i != numRows-1 && bias < size {
				zigzag.WriteByte(s[bias])
			}
		}
	}

	return zigzag.String()
}

var s = flag.String("s", "PAYPALISHIRING", "specify string")
var r = flag.Int("r", 3, "specify zigzag rows")

func main() {
	flag.Parse()
	fmt.Printf("zigzag convertion of '%s'(rows=%d) : %s\n", *s, *r, convert(*s, *r))
}
