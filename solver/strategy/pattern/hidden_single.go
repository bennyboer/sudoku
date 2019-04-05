package pattern

import "github.com/ob-algdatii-ss19/leistungsnachweis-sudo/model"

// Implementation of the hidden single pattern.
type HiddenSingle struct{}

// Apply pattern on Sudoku.
func (p *HiddenSingle) Apply(sudoku *model.Sudoku, possibleValuesRef *[][]*map[int]bool) (changed bool) {
	// TODO
	return false
}
