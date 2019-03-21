package cell

// Cell of a Sudoku.
type SudokuCell struct {
	// Where the cell is located in its corresponding Sudoku.
	position Coordinates
}

// Get a reference of the cells coordinates.
// Handle with care and do not update the cells coordinates.
func (c *SudokuCell) Position() *Coordinates {
	return &c.position
}
