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

// returns false if the Row is invalid and true if it is valid
func checkAllColumns(SudokuBoard *sudoku) bool {
	// potential speedup by only checking finished lines or specific lines
	for i := 0; i < SudokuBoard.Columns; i++ {
		for j := 0; j < SudokuBoard.Rows; j++ {
			for k := j + 1; k < SudokuBoard.Rows; k++ {
				if SudokuBoard.Board[j][i] == 0 {
					continue
				}

				if SudokuBoard.Board[j][i] == SudokuBoard.Board[k][i] {
					fmt.Printf("Column %v is invalid\n", i)
					return false
				}
			}
		}
	}
	return true
}

func checkAllRows(SudokuBoard *sudoku) bool {
	for i := 0; i < SudokuBoard.Rows; i++ {
		for j := 0; j < SudokuBoard.Columns; j++ {
			for k := j + 1; k < SudokuBoard.Columns; k++ {
				if SudokuBoard.Board[i][j] == 0 {
					continue
				}

				if SudokuBoard.Board[i][j] == SudokuBoard.Board[i][k] {
					fmt.Printf("Row %v is invalid\n", i)
					return false
				}
			}
		}
	}
	return true
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
