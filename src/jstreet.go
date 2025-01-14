package main

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
	loadBoard(SudokuBoard, "jane-street-current.txt")
	for i := 0; i < SudokuBoard.Columns; i++ {
		SudokuBoard.Board[4][i] = a[i]
	}
}
