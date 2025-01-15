package main

import "math"

// generate all permiutations of sudoku set for jstreet puzzle (0-9)
func generatePermiutations(k int, a []int, SudokuBoard *sudoku) {
	if k == 1 {
		for i := 0; i < len(a); i++ {
			generatePermBoards(SudokuBoard, a)
			printBoard(SudokuBoard)
		}
		return
	}

	generatePermiutations(k-1, a, SudokuBoard)

	for i := 0; i < k-1; i++ {
		if k%2 == 0 {
			a[i], a[k-1] = a[k-1], a[i]
		} else {
			a[0], a[k-1] = a[k-1], a[0]
		}
		generatePermiutations(k-1, a, SudokuBoard)
	}
}

// take permiutations and insert them into the jane street Board
func generatePermBoards(SudokuBoard *sudoku, a []int) {
	loadBoard(SudokuBoard, "jane-street.txt")
	for i := 0; i < SudokuBoard.Columns; i++ {
		SudokuBoard.Board[4][i] = a[i]
	}
}

func euclids(a int, b int) int {
	for b != 0 {
		temp := a % b
		a = b
		b = temp
	}
	return a
}

func gcdSet(nums []int) int {
	result := nums[0]
	for i := 1; i < len(nums); i++ {
		result = euclids(result, nums[i])
	}
	return result
}

// should be called to return the gcd of the rows of any sudoku board
func gcd(SudokuBoard *sudoku) int {
	rows := make([]int, SudokuBoard.Rows)
	for i := 0; i < SudokuBoard.Rows; i++ {
		curNum := 0
		for j := 0; j < SudokuBoard.Columns; j++ {
			if SudokuBoard.Board[i][j] == 10 {
				continue
			}

			curNum += SudokuBoard.Board[i][j] * int(math.Pow(10, float64(SudokuBoard.Columns-j-1)))
		}
		rows[i] = curNum
	}

	gcdResult := gcdSet(rows)
	return gcdResult
}
