package main

import "fmt"

/*
board.length == 9
board[i].length == 9
board[i][j] is a digit 1-9 or '.'.
*/

func isValidSudoku(board [][]byte) bool {
	desicion := true // if there are no dupes - true is as default

	cols := [9][10]bool{}    // idx_col
	rows := [9][10]bool{}    // idx_row
	squares := [9][10]bool{} // idx_row, idx_col

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			num := board[i][j]

			if num == '.' {
				continue
			}

			n := int(num - '0')

			boxIndex := (i/3)*3 + j/3

			if rows[i][n] || cols[j][n] || squares[boxIndex][n] {
				desicion = false
			}

			cols[j][num] = true
			rows[i][num] = true
			squares[boxIndex][num] = true
		}
	}

	return desicion
}

func main() {
	boardTrue := [][]byte{
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

	boardFalse := [][]byte{
		{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}

	fmt.Println("Valid Board:", isValidSudoku(boardTrue))    // true
	fmt.Println("Invalid Board:", isValidSudoku(boardFalse)) // false
}
