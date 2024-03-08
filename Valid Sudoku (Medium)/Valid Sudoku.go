package main

import "fmt"

func isValidSudoku(board [][]byte) bool {

	for i := 0; i < 9; i++ {
		rowMap := make(map[byte]struct{})
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				if _, exists := rowMap[board[i][j]]; exists {
					return false
				}
				rowMap[board[i][j]] = struct{}{}
			}
		}

	}

	for j := 0; j < 9; j++ {
		colMap := make(map[byte]struct{})
		for i := 0; i < 9; i++ {
			if board[i][j] != '.' {
				if _, exists := colMap[board[i][j]]; exists {
					return false
				}
				colMap[board[i][j]] = struct{}{}
			}
		}
	}

	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			squareMap := make(map[byte]struct{})
			for x := i; x < i+3; x++ {
				for y := j; y < j+3; y++ {
					if board[x][y] != '.' {
						if _, exists := squareMap[board[x][y]]; exists {
							return false
						}
						squareMap[board[x][y]] = struct{}{}
					}
				}
			}
		}
	}

	return true
}

func main() {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}

	fmt.Println(isValidSudoku(board))
}
