package pattern

import (
	"github.com/ob-algdatii-ss19/leistungsnachweis-sudo/model"
)

// Implementation of the hidden single pattern.
// It searches each row, column and block for values which can only be applied to one single field!
type HiddenSingle struct{}

// Apply pattern on Sudoku.
// Search all rows, columns and blocks for values which can only be applied to one single cell.
func (p *HiddenSingle) Apply(sudoku *model.Sudoku, possibleValuesRef *[][]*map[int]bool) (changed bool) {
	changed = false

	forEachUnit(func(unit []*map[int]bool) {
		changed = p.findAndUpdateHiddenSingles(unit) || changed
	}, possibleValuesRef)

	return
}

// Check for a hidden singles in the passed slice of possible values
// and process the changes in the possible value lookup.
func (p *HiddenSingle) findAndUpdateHiddenSingles(slice []*map[int]bool) bool {
	changed := false

	for value := 1; value <= model.SudokuSize; value++ {
		var uniqueLookupPtr *map[int]bool

		// Check if value is unique in unit slice
		for _, lookupPtr := range slice {
			lookup := *lookupPtr

			if possible, ok := lookup[value]; ok && possible {
				if uniqueLookupPtr != nil {
					uniqueLookupPtr = nil
					break // Value occurs more than once -> No need to continue
				}

				uniqueLookupPtr = lookupPtr
			}
		}

		if uniqueLookupPtr != nil {
			// Value occurs exclusively in the uniqueLookup.
			// Mark all other values of the lookup as impossible.
			uniqueLookup := *uniqueLookupPtr

			for v := range uniqueLookup {
				if v != value {
					uniqueLookup[v] = false
				}
			}

			changed = true
		}
	}

	return changed
}
