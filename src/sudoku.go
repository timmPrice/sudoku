package main

import "fmt"

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

func printBoard(SudokuBoard *sudoku) {
	for i := 0; i < SudokuBoard.Rows; i++ {
		for j := 0; j < SudokuBoard.Columns; j++ {
			if SudokuBoard.Board[i][j] == 0 {
				fmt.Print(".")
			} else {
				fmt.Printf("[%v]", SudokuBoard.Board[i][j])
			}
		}
		fmt.Print("\n")
	}
}

func loadBoard(SudokuBoard *sudoku) {
	for i := 0; i < SudokuBoard.Rows; i++ {
		for j := 0; j < SudokuBoard.Columns; j++ {
			if SudokuBoard.Board[i][j] == 0 {
				fmt.Print(".")
			} else {
				fmt.Printf("[%v]", SudokuBoard.Board[i][j])
			}
		}
		fmt.Print("\n")
	}
}
