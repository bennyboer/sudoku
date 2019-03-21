package cell

// Cell of a Sudoku.
type SudokuCell struct {
	// Where the cell is located in its corresponding Sudoku.
	Position Coordinates
	// The value of the cell. In range of [0; 9] where 0 means an empty cell.
	Value int
}
