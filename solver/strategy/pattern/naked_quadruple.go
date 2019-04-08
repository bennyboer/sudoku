package pattern

import (
	"github.com/ob-algdatii-ss19/leistungsnachweis-sudo/model"
)

// Implementation of the naked quadruple pattern.
type NakedQuadruple struct{}

// Apply pattern on Sudoku.
func (p *NakedQuadruple) Apply(sudoku *model.Sudoku, possibleValuesRef *[][]*map[int]bool) (changed bool) {
	changed = false

	forEachUnit(func(unit []*map[int]bool) {
		changed = p.findAndUpdateNakedQuadruples(unit) || changed
	}, possibleValuesRef)

	return
}

// Check for a naked quadruples in the passed slice and process the changes in the possible value lookup.
func (p *NakedQuadruple) findAndUpdateNakedQuadruples(slice []*map[int]bool) bool {
	if len(slice) <= 4 {
		return false // We need at least 5 cells, otherwise this wont do anything
	}

	lookups, values := p.findFourMatchingLookups(slice)

	if lookups == nil || values == nil {
		return false
	}

	takenLookups := make(map[*map[int]bool]bool)
	for _, lookupPtr := range lookups {
		takenLookups[lookupPtr] = true
	}

	// Now we found a naked quadruple!
	// -> Eliminate the four unique values from all other lookups in the same unit (Row, column or block)
	changed := false
	for _, lookupPtr := range slice {
		_, ok := takenLookups[lookupPtr]

		// Check that lookup ptr is none of the cells with a quadruple
		if !ok {
			lookup := *lookupPtr

			for _, value := range values {
				if possible, _ := lookup[value]; possible {
					lookup[value] = false
					changed = true
				}
			}
		}
	}

	return changed
}

// Check if there are four lookups among the passed with no more than 4 unique values.
// Will return the four lookups and the four unique values, or nil if there are none matching lookups.
func (p *NakedQuadruple) findFourMatchingLookups(lookups []*map[int]bool) ([]*map[int]bool, []int) {
	length := len(lookups)

	for a := 0; a < length; a++ {
		for b := a + 1; b < length; b++ {
			for c := b + 1; c < length; c++ {
				for d := c + 1; d < length; d++ {
					if values := findNakedValues(4, lookups[a], lookups[b], lookups[c], lookups[d]); values != nil {
						return []*map[int]bool{
							lookups[a],
							lookups[b],
							lookups[c],
							lookups[d],
						}[:], values
					}
				}
			}
		}
	}

	return nil, nil
}
