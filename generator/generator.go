package generator

import "github.com/ob-algdatii-ss19/leistungsnachweis-sudo/model"

// Interface for Sudoku generators.
type SudokuGenerator interface {
	// Generates a Sudoku with the passed difficulty of range [0.0; 1.0],
	// where 0.0 is as easy as possible and 1.0 as difficult as possible.
	Generate(difficulty float64) *model.Sudoku
}

type SudokuGeneratorBackTracking struct {
}
