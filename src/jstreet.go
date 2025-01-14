package main

import (
	"fmt"
)

// generate all permiutations of sudoku set for jstreet puzzle (0-9)
func generatePermiutations(k int, a []int) {
	if k == 1 {
		fmt.Println(a)
		return
	}

	generatePermiutations(k-1, a)

	for i := 0; i < k-1; i++ {
		if k%2 == 0 {
			a[i], a[k-1] = a[k-1], a[i]
		} else {
			a[0], a[k-1] = a[k-1], a[0]
		}
		generatePermiutations(k-1, a)
	}
}

// take permiutations and insert them into the jane street Board
func generateBoards() {

}
