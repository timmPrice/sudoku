package main

import "fmt"

func main() {
	Board := makeBoard(9, 9)
	loadBoard(Board, "valid-solution.txt")
	printBoard(Board)

	fmt.Println("")

	checkAllColumns(Board)
	checkAllRows(Board)
	checkAllGrids(Board)
}
