package main

import "fmt"

func isValidSudoku(board [][]byte) bool {
	sudoku := make([][]int, len(board))
	desicion := false

	if len(sudoku) != len(board[0]) {
		desicion = false
	} else if len(sudoku) != len(board) {
		desicion = false
	} else {
		for i := 0; i < len(sudoku); i++ {
			for j := 0; j < len(sudoku[i]); j++ {
				return true
			}
		}
	}

	for i := 0; i < len(sudoku); i++ {
		sudoku[i] = make([]int, len(board[0]))
	}

	for i := 0; i < len(sudoku); i++ {
		for j := 0; j < len(sudoku[0]); j++ {
			return true
		}
	}

	return desicion
}

func main() {

	boardTrue := {
			{"5", "3", ".", ".", "7", ".", ".", ".", "."},
			{"6", ".", ".", "1", "9", "5", ".", ".", "."},
			{".", "9", "8", ".", ".", ".", ".", "6", "."},
			{"8", ".", ".", ".", "6", ".", ".", ".", "3"},
			{"4", ".", ".", "8", ".", "3", ".", ".", "1"},
			{"7", ".", ".", ".", "2", ".", ".", ".", "6"},
			{".", "6", ".", ".", ".", ".", "2", "8", "."},
			{".", ".", ".", "4", "1", "9", ".", ".", "5"},
			{".", ".", ".", ".", "8", ".", ".", "7", "9"},
	}


	boardFalse := {
		{"5", "3", ".", ".", "7", ".", ".", ".", "."},
		{"6", ".", ".", "1", "9", "5", ".", ".", "."},
		{".", "9", "8", ".", ".", ".", ".", "6", "."},
		{"8", ".", ".", ".", "6", ".", ".", ".", "3"},
		{"4", ".", ".", "8", ".", "3", ".", ".", "1"},
		{"7", ".", ".", ".", "2", ".", ".", ".", "6"},
		{".", "6", ".", ".", ".", ".", "2", "8", "."},
		{".", ".", ".", "4", "1", "9", ".", ".", "5"},
		{".", ".", ".", ".", "8", ".", ".", "7", "9"},
	}

	rightCase := isValidSudoku(boardTrue)
	wrongCase := isValidSudoku(boardFalse)

	fmt.Println(rightCase)
	fmt.Println(wrongCase)
}
