package pattern

import (
	"github.com/ob-algdatii-ss19/leistungsnachweis-sudo/model"
)

// Implementation of the hidden quadruple pattern.
type HiddenQuadruple struct{}

// Apply pattern on Sudoku.
func (p *HiddenQuadruple) Apply(sudoku *model.Sudoku, possibleValuesRef *[][]*map[int]bool) (changed bool) {
	changed = false

	forEachUnit(func(unit []*map[int]bool) {
		changed = p.findAndUpdateHiddenQuadruple(unit) || changed
	}, possibleValuesRef)

	return
}

// Check for a hidden quadruple in the passed slice and process the changes in the possible value lookup.
func (p *HiddenQuadruple) findAndUpdateHiddenQuadruple(slice []*map[int]bool) bool {
	if len(slice) <= 4 {
		return false // We need at least 5 cells for the algorithm to change a thing!
	}

	htLookups, values := p.findHiddenQuadrupleLookups(slice)

	if values == nil {
		return false
	}

	valueSet := make(map[int]bool)
	for _, value := range values {
		valueSet[value] = true
	}

	changed := false

	for _, htLookupPtr := range htLookups {
		htLookup := *htLookupPtr

		for value, possible := range htLookup {
			if possible {
				if _, keep := valueSet[value]; !keep {
					htLookup[value] = false
					changed = true
				}
			}
		}
	}

	return changed
}

// Find lookups with hidden quadruples in them and the values forming the hidden quadruple.
func (p *HiddenQuadruple) findHiddenQuadrupleLookups(lookups []*map[int]bool) ([]*map[int]bool, []int) {
	length := len(lookups)

	for a := 0; a < length; a++ {
		for b := a + 1; b < length; b++ {
			for c := b + 1; c < length; c++ {
				for d := c + 1; d < length; d++ {
					lookupsToCheck := []*map[int]bool{
						lookups[a],
						lookups[b],
						lookups[c],
						lookups[d],
					}

					otherLookups := make([]*map[int]bool, len(lookups)-len(lookupsToCheck))
					for _, lookupPtr := range lookups {
						if lookupPtr != lookups[a] && lookupPtr != lookups[b] && lookupPtr != lookups[c] && lookupPtr != lookups[d] {
							otherLookups = append(otherLookups, lookupPtr)
						}
					}

					if values := findHiddenNValues(len(lookupsToCheck), lookupsToCheck, otherLookups); values != nil {
						return lookupsToCheck, values
					}
				}
			}
		}
	}

	return nil, nil
}
