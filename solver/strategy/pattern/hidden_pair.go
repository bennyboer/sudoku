package pattern

import "github.com/ob-algdatii-ss19/leistungsnachweis-sudo/model"

// Implementation of the hidden pair pattern.
type HiddenPair struct{}

// Apply pattern on Sudoku.
func (p *HiddenPair) Apply(sudoku *model.Sudoku, possibleValuesRef *[][]*map[int]bool) (changed bool) {
	// TODO

	return false
}
