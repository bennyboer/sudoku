package pattern

import (
	"github.com/ob-algdatii-ss19/leistungsnachweis-sudo/model"
	"sort"
)

// Implementation of the hidden pair pattern.
type HiddenPair struct{}

// Apply pattern on Sudoku.
func (p *HiddenPair) Apply(sudoku *model.Sudoku, possibleValuesRef *[][]*map[int]bool) (changed bool) {
	changed = false

	forEachUnit(func(unit []*map[int]bool) {
		changed = p.findAndUpdateHiddenPairs(unit) || changed
	}, possibleValuesRef)

	return
}

// Check for a hidden pairs in the passed slice of possible values
// and process the changes in the possible value lookup.
func (p *HiddenPair) findAndUpdateHiddenPairs(slice []*map[int]bool) bool {
	// Count occurrences of each value in the slice
	occurrences := make(map[int]int)
	for _, lookupPtr := range slice {
		lookup := *lookupPtr

		for value, possible := range lookup {
			if possible {
				count, ok := occurrences[value]

				if ok {
					occurrences[value] = count + 1
				} else {
					occurrences[value] = 1
				}
			}
		}
	}

	// Collect all values occurring twice
	occursTwice := make([]int, 0, model.SudokuSize)
	for value, count := range occurrences {
		if count == 2 {
			occursTwice = append(occursTwice, value)
		}
	}

	// We need at least two values here, otherwise exit early.
	if len(occursTwice) < 2 {
		return false
	}

	sort.Ints(occursTwice)

	// Calculate all combinations = n elements
	// combinations = (n - 1) + (n - 2) + ... + 1
	combinations := 1
	for i := 2; i <= len(occursTwice); i++ {
		combinations += i
	}

	// Collect all combinations of the twice-occurring values
	pairs := make([][]int, 0, combinations)
	for i := 0; i < len(occursTwice); i++ {
		for ii := i + 1; ii < len(occursTwice); ii++ {
			pair := make([]int, 2)
			pair[0] = occursTwice[i]
			pair[1] = occursTwice[ii]
			pairs = append(pairs, pair)
		}
	}

	// Find two possible value lookups in slice with two equal of the twice-occurring values
	for i := 0; i < len(pairs); i++ {
		pair := pairs[i]

		var lookup1Ptr *map[int]bool = nil
		var lookup2Ptr *map[int]bool = nil

		for _, lookupPtr := range slice {
			lookup := *lookupPtr

			possible1, foundValue1 := lookup[pair[0]]
			possible2, foundValue2 := lookup[pair[1]]

			if foundValue1 && foundValue2 && possible1 && possible2 {
				if lookup1Ptr == nil {
					lookup1Ptr = lookupPtr
				} else {
					lookup2Ptr = lookupPtr
				}
			}
		}

		if lookup1Ptr != nil && lookup2Ptr != nil {
			// Found two lookups with the two values -> Remove all other possible values from them
			changed := false

			lookup := *lookup1Ptr
			for v, possible := range lookup {
				if possible && v != pair[0] && v != pair[1] {
					lookup[v] = false
					changed = true
				}
			}

			lookup = *lookup2Ptr
			for v, possible := range *lookup2Ptr {
				if possible && v != pair[0] && v != pair[1] {
					lookup[v] = false
					changed = true
				}
			}

			return changed
		}
	}

	return false
}
