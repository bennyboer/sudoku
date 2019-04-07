package pattern

import (
	"github.com/ob-algdatii-ss19/leistungsnachweis-sudo/model"
	"sort"
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
	// First and foremost find all pair possibilities (Only two possible values for a cell).
	pairPossibilities := make([]*[]int, 0, model.SudokuSize)
	pairPossibilitiesLookups := make(map[*map[int]bool]bool)

	for _, pvLookup := range slice {
		pair := make([]int, 0, 2)
		index := 0
		add := true
		for value, possible := range *pvLookup {
			if possible {
				if index == 2 {
					add = false
					break
				}

				pair = append(pair, value)
				index++
			}
		}

		if add && len(pair) == 2 {
			sort.Ints(pair) // Sort in order to compare it later easily
			pairPossibilities = append(pairPossibilities, &pair)
			pairPossibilitiesLookups[pvLookup] = true
		}
	}

	if len(pairPossibilities) != 3 {
		return false
	}

	// Check that the 3 pairs include only 3 unique values
	valuesSet := make(map[int]bool)
	for _, pairPtr := range pairPossibilities {
		pair := *pairPtr

		for _, value := range pair {
			valuesSet[value] = true
		}
	}

	if len(valuesSet) != 3 {
		return false
	}

	// Now we found a naked triple!
	// -> Eliminate the three unique values from all other lookups in the same unit (Row, column or block)
	changed := false
	for _, lookupPtr := range slice {
		_, ok := pairPossibilitiesLookups[lookupPtr]

		// Check that lookup ptr is none of the cells with a pair
		if !ok {
			lookup := *lookupPtr

			for value, _ := range valuesSet {
				if possible, _ := lookup[value]; possible {
					lookup[value] = false
					changed = true
				}
			}
		}
	}

	return changed
}
