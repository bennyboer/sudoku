package model

import "github.com/ob-algdatii-ss19/leistungsnachweis-sudo/model/cell"

// A model for a Sudoku.
type Sudoku struct {
	// All cells of a sudoku.
	cells [][]cell.SudokuCell
}

// Retrieve a reference to all cells in the Sudoku.
func (s *Sudoku) Cells() *[][]cell.SudokuCell {
	return &s.cells
}
