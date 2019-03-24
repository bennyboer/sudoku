package strategy

import (
	"fmt"
	"github.com/ob-algdatii-ss19/leistungsnachweis-sudo/model"
)

// Strategy used to find the next empty cell in a Sudoku for the backtracking algorithm.
type CellChoosingStrategy interface {
	// Get empty cells.
	Get(sudoku *model.Sudoku) *[]*model.SudokuCell
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
