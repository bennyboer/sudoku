package pattern

import (
	"github.com/ob-algdatii-ss19/leistungsnachweis-sudo/model"
	"sort"
)

// Implementation of the naked pair pattern.
type NakedPair struct{}

// Apply pattern on Sudoku.
func (p *NakedPair) Apply(sudoku *model.Sudoku, possibleValuesRef *[][]*map[int]bool) (changed bool) {
	changed = false
	pv := *possibleValuesRef

	// Search rows
	for row := 0; row < model.SudokuSize; row++ {
		possibleRowValues := make([]*map[int]bool, 0, model.SudokuSize)

		for column := 0; column < model.SudokuSize; column++ {
			if pv[row][column] != nil {
				possibleRowValues = append(possibleRowValues, pv[row][column])
			}
		}

		didChange := p.findAndUpdatedNakedPairs(possibleRowValues)
		changed = changed || didChange
	}

	// Search columns
	for column := 0; column < model.SudokuSize; column++ {
		possibleColumnValues := make([]*map[int]bool, 0, model.SudokuSize)

		for row := 0; row < model.SudokuSize; row++ {
			if pv[row][column] != nil {
				possibleColumnValues = append(possibleColumnValues, pv[row][column])
			}
		}

		didChange := p.findAndUpdatedNakedPairs(possibleColumnValues)
		changed = changed || didChange
	}

	// Search blocks
	for block := 0; block < model.SudokuSize; block++ {
		possibleBlockValues := make([]*map[int]bool, 0, model.SudokuSize)

		startRow := block / model.BlockSize * model.BlockSize
		startColumn := (block * model.BlockSize) % model.SudokuSize
		for row := startRow; row < startRow+model.BlockSize; row++ {
			for column := startColumn; column < startColumn+model.BlockSize; column++ {
				if pv[row][column] != nil {
					possibleBlockValues = append(possibleBlockValues, pv[row][column])
				}
			}
		}

		didChange := p.findAndUpdatedNakedPairs(possibleBlockValues)
		changed = changed || didChange
	}

	return
}

// Check for a naked pair in the passed slice and process the changes in the Sudoku and possible value lookup.
func (p *NakedPair) findAndUpdatedNakedPairs(slice []*map[int]bool) bool {
	// First and foremost find all pair possibilities (Only two possible values for a cell).
	pairPossibilities := make([]*[]int, 0, model.SudokuSize)
	pairPossibilitiesLookups := make([]*map[int]bool, 0, model.SudokuSize)

	for _, pvLookup := range slice {
		pair := make([]int, 2)
		index := 0
		add := true
		for value, possible := range *pvLookup {
			if possible {
				if index == 2 {
					add = false
					break
				}

				pair[index] = value
				index++
			}
		}

		if add {
			sort.Ints(pair) // Sort in order to compare it later easily
			pairPossibilities = append(pairPossibilities, &pair)
			pairPossibilitiesLookups = append(pairPossibilitiesLookups, pvLookup)
		}
	}

	if len(pairPossibilities) < 2 {
		return false
	}

	// Compare pair possibilities if equal
	for i1, pairPtr1 := range pairPossibilities {
		for i2, pairPtr2 := range pairPossibilities {
			// Check that not the same slice
			if pairPtr1 != pairPtr2 {
				pair1 := *pairPtr1
				pair2 := *pairPtr2

				// Check if equal
				if pair1[0] == pair2[0] && pair1[1] == pair2[1] {
					// Found one! Remove the two values from the other possible value lookups
					changed := false

					for _, pvLookupPtr := range slice {
						if pvLookupPtr != pairPossibilitiesLookups[i1] && pvLookupPtr != pairPossibilitiesLookups[i2] {
							pvLookup := *pvLookupPtr

							possible, _ := pvLookup[pair1[0]]
							if possible {
								pvLookup[pair1[0]] = false
								changed = true
							}

							possible, _ = pvLookup[pair1[1]]
							if possible {
								pvLookup[pair1[1]] = false
								changed = true
							}
						}
					}

					return changed
				}
			}
		}
	}

	return false
}
