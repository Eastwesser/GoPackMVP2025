package main

import "fmt"

func checkRowIsValid(board [][]byte, row int) bool {
	seen := make(map[byte]bool)

	for _, num := range board[row] {
		if num == '.' {
			continue
		}
		if seen[num] {
			return false
		}
		seen[num] = true
	}

	return true
}

func checkColumnIsValid(board [][]byte, col int) bool {
	seen := make(map[byte]bool)

	for i := 0; i < 9; i++ {
		num := board[i][col]
		if num == '.' {
			continue
		}
		if seen[num] {
			return false
		}
		seen[num] = true
	}

	return true
}

func checkSubGridIsValid(board [][]byte, startRow, startCol int) bool {
	seen := make(map[byte]bool)

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			num := board[startRow+i][startCol+j]
			if num == '.' {
				continue
			}
			if seen[num] {
				return false
			}
			seen[num] = true
		}
	}

	return true
}

func isValidSudoku(board [][]byte) bool {
	for i := 0; i < 9; i++ {

		if !checkRowIsValid(board, i) || !checkColumnIsValid(board, i) {
			return false
		}
	}

	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			if !checkSubGridIsValid(board, i, j) {
				return false
			}
		}
	}

	return true
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
