package pattern

import "github.com/ob-algdatii-ss19/leistungsnachweis-sudo/model"

// Sudoku pattern.
type Pattern interface {
	// Apply the pattern on the given sudoku until it cannot change more.
	// The possible values per cell are passed as reference and must be updated when
	// filling values in the Sudoku.
	// Returns whether the pattern could change at least one value in the Sudoku or the possible cell values.
	Apply(sudoku *model.Sudoku, possibleValuesRef *[][]*map[int]bool) (changed bool)
}
