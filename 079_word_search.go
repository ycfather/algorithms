package main

import (
	"fmt"
)

var directs [][]int = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func exist(board [][]byte, word string) bool {
	rows := len(board)
	if rows == 0 {
		return false
	}
	columns := len(board[0])

	mark := make([][]byte, rows)
	for i := 0; i < rows; i++ {
		mark[i] = make([]byte, columns)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			if board[i][j] == word[0] {
				mark[i][j] = 1
				if backtrack(i, j, &mark, board, word[1:]) {
					return true
				} else {
					mark[i][j] = 0
				}
			}
		}
	}

	return false
}

func backtrack(i, j int, mark *[][]byte, board [][]byte, word string) bool {
	if len(word) == 0 {
		return true
	}

	for _, direct := range directs {
		ni := i + direct[0]
		nj := j + direct[1]

		if ni >= 0 && ni < len(board) && nj >= 0 && nj < len(board[0]) && board[ni][nj] == word[0] {
			if (*mark)[ni][nj] == 1 {
				continue
			}

			(*mark)[ni][nj] = 1
			if backtrack(ni, nj, mark, board, word[1:]) {
				return true
			} else {
				(*mark)[ni][nj] = 0
			}
		}
	}

	return false
}

func main() {
	board := [][]byte{[]byte{'A', 'B', 'C', 'E'}, []byte{'S', 'F', 'C', 'S'}, []byte{'A', 'D', 'E', 'E'}}

	words := []string{"ABCCED", "SEE", "ABCB"}
	for _, word := range words {
		fmt.Printf("%s: %v\n", word, exist(board, word))
	}
}
