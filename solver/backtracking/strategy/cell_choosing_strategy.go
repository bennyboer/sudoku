package strategy

import (
	"fmt"
	"github.com/ob-algdatii-ss19/leistungsnachweis-sudo/model"
)

// Strategy used to find the next empty cell in a Sudoku for the backtracking algorithm.
type CellChoosingStrategy interface {
	// Initialize the strategy with the passed Sudoku.
	Initialize(sudoku *model.Sudoku)
	// Find the next cell.
	// May return an error if the strategy is not initialized.
	// Returns nil if there are no more cells.
	FindNext() (*model.SudokuCell, error)
}

// Create a strategy instance.
func Create(t CellChoosingStrategyType) (CellChoosingStrategy, error) {
	switch t {
	case Linear:
		return &LinearCellChooser{}, nil
	case Random:
		return &RandomCellChooser{}, nil
	default:
		return nil, fmt.Errorf("CellChoosingStrategyType unknown")
	}
}
