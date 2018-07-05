package main

import (
	"fmt"
)

func isValidSudoku(board [][]byte) bool {
	if len(board) != 9 {
		return false
	}
	for i := 0; i < 9; i++ {
		if len(board[i]) != 9 {
			return false
		}
		rowMap := make(map[byte]int)
		colMap := make(map[byte]int)
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' { // check row
				if _, ok := rowMap[board[i][j]]; ok {
					return false
				}
				rowMap[board[i][j]] = 1
			}

			if board[j][i] != '.' { // check column
				if _, ok := colMap[board[j][i]]; ok {
					return false
				}
				colMap[board[j][i]] = 1
			}

			if i%3 == 0 && j%3 == 0 { // check block
				blockMap := make(map[byte]int)
				for bi := i; bi < i+3; bi++ {
					for bj := j; bj < j+3; bj++ {
						if board[bi][bj] != '.' {
							if _, ok := blockMap[board[bi][bj]]; ok {
								return false
							}
							blockMap[board[bi][bj]] = 1
						}
					}
				}
			}
		}
	}

	return true
}

func main() {
	board := [][]byte{[]byte{'5', '3', '.', '.', '7', '.', '.', '.', '.'}, []byte{'6', '.', '.', '1', '9', '5', '.', '.', '.'}, []byte{'.', '9', '8', '.', '.', '.', '.', '6', '.'}, []byte{'8', '.', '.', '.', '6', '.', '.', '.', '3'}, []byte{'4', '.', '.', '8', '.', '3', '.', '.', '1'}, []byte{'7', '.', '.', '.', '2', '.', '.', '.', '6'}, []byte{'.', '6', '.', '.', '.', '.', '2', '8', '.'}, []byte{'.', '.', '.', '4', '1', '9', '.', '.', '5'}, []byte{'.', '.', '.', '.', '8', '.', '.', '7', '9'}}
	ret := isValidSudoku(board)
	fmt.Printf("valid : %v\n", ret)
}
