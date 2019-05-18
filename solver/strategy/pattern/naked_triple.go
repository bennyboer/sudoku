package pattern

import (
	"github.com/ob-algdatii-ss19/leistungsnachweis-sudo/model"
)

// Implementation of the naked triple pattern.
type NakedTriple struct{}

// Apply pattern on Sudoku.
func (p *NakedTriple) Apply(sudoku *model.Sudoku, possibleValuesRef *[][]*map[int]bool) (changed bool) {
	changed = false

	forEachUnit(func(unit []*map[int]bool) {
		changed = p.findAndUpdateNakedTriples(unit) || changed
	}, possibleValuesRef)

	return
}

// Check for a naked triple in the passed slice and process the changes in the possible value lookup.
func (p *NakedTriple) findAndUpdateNakedTriples(slice []*map[int]bool) bool {
	if len(slice) <= 3 {
		return false // We need at least 4 cells, otherwise this wont do anything
	}

	lookups, values := p.findThreeMatchingLookups(slice)

	if lookups == nil || values == nil {
		return false
	}

	takenLookups := make(map[*map[int]bool]bool)
	for _, lookupPtr := range lookups {
		takenLookups[lookupPtr] = true
	}

	// Now we found a naked triple!
	// -> Eliminate the three unique values from all other lookups in the same unit (Row, column or block)
	changed := false
	for _, lookupPtr := range slice {
		_, ok := takenLookups[lookupPtr]

		// Check that lookup ptr is none of the cells with a triple
		if !ok {
			lookup := *lookupPtr

			for _, value := range values {
				if possible := lookup[value]; possible {
					lookup[value] = false
					changed = true
				}
			}
		}
	}

	return changed
}

// Check if there are three lookups among the passed with no more than 3 unique values.
// Will return the three lookups and the three unique values, or nil if there are none matching lookups.
func (p *NakedTriple) findThreeMatchingLookups(lookups []*map[int]bool) ([]*map[int]bool, []int) {
	length := len(lookups)

	for a := 0; a < length; a++ {
		for b := a + 1; b < length; b++ {
			for c := b + 1; c < length; c++ {
				if values := findNakedValues(3, lookups[a], lookups[b], lookups[c]); values != nil {
					return []*map[int]bool{
						lookups[a],
						lookups[b],
						lookups[c],
					}[:], values
				}
			}
		}
	}

	return nil, nil
}
