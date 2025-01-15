package main

import "fmt"

func main() {

	Board := makeBoard(9, 9)
	loadBoard(Board, "jane-street.txt")
	printBoard(Board)

	bruteForce(Board, 0, 0)

	fmt.Println("")

	if checkWin(Board) {
		printBoard(Board)
		fmt.Println("solved!")
	}

	fmt.Println(gcd(Board))
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
