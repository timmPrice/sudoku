package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type sudoku struct {
	Rows    int
	Columns int
	Board   [][]int
}

func makeBoard(rows int, columns int) *sudoku {
	board := make([][]int, rows)
	for i := range board {
		board[i] = make([]int, columns)
	}
	return &sudoku{
		Rows:    rows,
		Columns: columns,
		Board:   board,
	}
}

// returns false if the Column is invalid and true if it is valid
func checkAllColumns(SudokuBoard *sudoku) bool {
	// potential speedup by only checking finished lines or specific lines
	for i := 0; i < SudokuBoard.Columns; i++ {
		seen := make(map[int]bool)
		for j := 0; j < SudokuBoard.Rows; j++ {
			num := SudokuBoard.Board[j][i]
			if num == 0 {
				continue
			}

			if seen[num] {
				fmt.Printf("Column %v is invalid\n", i)
				return false
			}
		}
	}
	return true
}

// returns false if the Row is invalid and true if it is valid
func checkAllRows(SudokuBoard *sudoku) bool {
	for i := 0; i < SudokuBoard.Rows; i++ {
		seen := make(map[int]bool)
		for j := 0; j < SudokuBoard.Columns; j++ {
			num := SudokuBoard.Board[i][j]
			if num == 0 {
				continue
			}

			if seen[num] {
				fmt.Printf("Row %v is invalid\n", i)
				return false
			}

			seen[num] = true
		}
	}
	return true
}

// checks all Sub Grids and returns true if valid and false if invalid
func checkAllGrids(SudokuBoard *sudoku) bool {
	return checkGridsPriv(SudokuBoard, 0, 0)
}

func checkGridsPriv(SudokuBoard *sudoku, startRow int, startCol int) bool {
	if startRow >= SudokuBoard.Rows {
		return true
	}

	if checkSubGrid(SudokuBoard, startRow, startCol) != true {
		return false
	}

	if startCol+3 < SudokuBoard.Columns {
		return checkGridsPriv(SudokuBoard, startRow, startCol+3)
	} else {
		return checkGridsPriv(SudokuBoard, startRow+3, 0)
	}
}

// Assumes 3x3 subsections
// returns true if the block is valid and false if the block is invalid
func checkSubGrid(SudokuBoard *sudoku, startRow int, startCol int) bool {
	seen := make(map[int]bool)
	for i := startRow; i < startRow+3; i++ {
		for j := startCol; j < startCol+3; j++ {
			num := SudokuBoard.Board[i][j]

			if num == 0 {
				continue
			}

			if seen[num] {
				fmt.Printf("3x3 subgrid starting at (%v, %v) is invalid\n", startRow, startCol)
				return false
			}

			seen[num] = true
		}
	}
	return true
}

// runs all test to verify that the board is solved
func checkWin() {
	// todo
}

func printBoard(SudokuBoard *sudoku) {
	for i := 0; i < SudokuBoard.Rows; i++ {
		for j := 0; j < SudokuBoard.Columns; j++ {
			if SudokuBoard.Board[i][j] == 0 {
				fmt.Print("[.]")
			} else {
				fmt.Printf("[%v]", SudokuBoard.Board[i][j])
			}
		}

		fmt.Print("\n")
	}
}

func loadBoard(SudokuBoard *sudoku, filename string) {
	file, err := os.ReadFile("./puzzles/" + filename)
	if err != nil {
		fmt.Println("Error reading file: ", err)
		return
	}

	arrayPosition := 0
	cleanedFile := strings.ReplaceAll(string(file), "\n", "")
	data := strings.Split(cleanedFile, ",")

	for i := 0; i < SudokuBoard.Rows; i++ {
		for j := 0; j < SudokuBoard.Columns; j++ {
			value, err := strconv.Atoi(data[arrayPosition])
			if err != nil {
				arrayPosition++
				continue
			}

			SudokuBoard.Board[i][j] = value
			arrayPosition++
		}
	}
}

// makes new board file and adds it to "runable list"
func generateBoard() {
	// todo
}
