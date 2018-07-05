package main

import (
	"fmt"
)

func solveSudoku(board [][]byte) {
	if len(board) != 9 {
		return
	}

	for _, v := range board {
		if len(v) != 9 {
			return
		}
	}
	solve(board)
}

func solve(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				for b := '1'; b <= '9'; b++ {
					if isValid(board, i, j, byte(b)) {
						board[i][j] = byte(b)
						if solve(board) {
							return true
						} else {
							board[i][j] = '.'
						}
					}
				}
				return false
			}
		}
	}
	return true
}

func isValid(board [][]byte, row, col int, b byte) bool {
	for i := 0; i < 9; i++ {
		if board[row][i] == b || board[i][col] == b {
			return false
		}

		if board[row/3*3+i/3][col/3*3+i%3] == b {
			return false
		}
	}

	return true
}

func prettyPrint(board [][]byte) {
	for _, v := range board {
		for _, b := range v {
			fmt.Printf("%s ", string(b))
		}
		fmt.Println()
	}
}

func main() {
	board := [][]byte{[]byte{'5', '3', '.', '.', '7', '.', '.', '.', '.'}, []byte{'6', '.', '.', '1', '9', '5', '.', '.', '.'}, []byte{'.', '9', '8', '.', '.', '.', '.', '6', '.'}, []byte{'8', '.', '.', '.', '6', '.', '.', '.', '3'}, []byte{'4', '.', '.', '8', '.', '3', '.', '.', '1'}, []byte{'7', '.', '.', '.', '2', '.', '.', '.', '6'}, []byte{'.', '6', '.', '.', '.', '.', '2', '8', '.'}, []byte{'.', '.', '.', '4', '1', '9', '.', '.', '5'}, []byte{'.', '.', '.', '.', '8', '.', '.', '7', '9'}}
	fmt.Printf("original : \n")
	prettyPrint(board)
	solveSudoku(board)
	fmt.Printf("solved : \n")
	prettyPrint(board)
}
