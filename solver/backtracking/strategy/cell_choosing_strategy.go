package strategy

import "github.com/ob-algdatii-ss19/leistungsnachweis-sudo/model"

// Strategy used to find the next empty cell in a Sudoku for the backtracking algorithm.
type CellChoosingStrategy interface {
	// Initialize the strategy with the passed Sudoku.
	Initialize(sudoku *model.Sudoku)
	// Find the next cell.
	// May return an error if the strategy is not initialized.
	FindNext() (*model.SudokuCell, error)
}
