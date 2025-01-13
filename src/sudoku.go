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
