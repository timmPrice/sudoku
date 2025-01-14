package main

import "fmt"

func main() {
	set := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	generatePermiutations(len(set), set)
	fmt.Println("done")
}

// func main() {
// 	Board := makeBoard(9, 9)
// 	loadBoard(Board, "single-solution.txt")
// 	printBoard(Board)
//
// 	bruteForce(Board, 0, 0)
//
// 	fmt.Println("")
// 	if checkWin(Board) {
// 		printBoard(Board)
// 		fmt.Println("solved!")
// 	}
//
// }
