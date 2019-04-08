package pattern

import (
	"github.com/ob-algdatii-ss19/leistungsnachweis-sudo/model"
)

// Helper function to update a value in the passed Sudoku as well in the possible values lookup matrix.
func updateValueInSudokuAndLookup(sudoku *model.Sudoku, possibleValuesRef *[][]*map[int]bool, row int, column int, value int) {
	pv := *possibleValuesRef

	// Fill cell in Sudoku and update possible value lookup
	cell := sudoku.Cells[row][column]
	cell.SetValue(value)

	// Update possible value lookup
	if pv[row][column] != nil {
		pv[row][column] = nil
	}

	// Update neighbour lookups
	for _, neighbour := range cell.Neighbours().All {
		position := neighbour.Position()

		// Set as "no more possible" in lookup
		lookupPtr := pv[position.Row][position.Column]
		if lookupPtr != nil {
			(*lookupPtr)[value] = false // Mark the value as "no more possible"
		}
	}
}

// Function processing a unit (Row, column or block) of possible value maps.
type unitFunction func([]*map[int]bool)

// Process the passed function for each unit (Row, column or block) of possible value maps.
func forEachUnit(fn unitFunction, possibleValuesRef *[][]*map[int]bool) {
	// Rows
	for row := 0; row < model.SudokuSize; row++ {
		fn(getRowPossibleValues(row, possibleValuesRef))
	}

	// Columns
	for column := 0; column < model.SudokuSize; column++ {
		fn(getColumnPossibleValues(column, possibleValuesRef))
	}

	// Blocks
	for block := 0; block < model.SudokuSize; block++ {
		fn(getBlockPossibleValues(block, possibleValuesRef))
	}
}

// Get all possible values for a block in range [0; 8]
func getBlockPossibleValues(block int, possibleValuesRef *[][]*map[int]bool) []*map[int]bool {
	pv := *possibleValuesRef

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

	return possibleBlockValues
}

// Get all possible values for a row in range [0; 8]
func getRowPossibleValues(row int, possibleValuesRef *[][]*map[int]bool) []*map[int]bool {
	pv := *possibleValuesRef

	possibleRowValues := make([]*map[int]bool, 0, model.SudokuSize)

	for column := 0; column < model.SudokuSize; column++ {
		if pv[row][column] != nil {
			possibleRowValues = append(possibleRowValues, pv[row][column])
		}
	}

	return possibleRowValues
}

// Get all possible values for a column in range [0; 8]
func getColumnPossibleValues(column int, possibleValuesRef *[][]*map[int]bool) []*map[int]bool {
	pv := *possibleValuesRef

	possibleColumnValues := make([]*map[int]bool, 0, model.SudokuSize)

	for row := 0; row < model.SudokuSize; row++ {
		if pv[row][column] != nil {
			possibleColumnValues = append(possibleColumnValues, pv[row][column])
		}
	}

	return possibleColumnValues
}

// Count the possible values in the passed possible value lookup.
func countPossibleValues(lookup *map[int]bool) int {
	result := 0

	if lookup != nil {
		for _, possible := range *lookup {
			if possible {
				result++
			}
		}
	}

	return result
}

// Count the occurrences of all values in the passed possible value lookups.
func countValueOccurrences(lookups []*map[int]bool) *map[int]int {
	occurrences := make(map[int]int)

	for _, lookupPtr := range lookups {
		if lookupPtr != nil {
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
	}

	return &occurrences
}

// Find naked N (count) values in the passed possible value lookups or nil if none are found.
func findNakedValues(count int, lookups ...*map[int]bool) []int {
	valuesSet := make(map[int]bool)

	for _, lookupPtr := range lookups {
		lookup := *lookupPtr

		for value, possible := range lookup {
			if possible {
				valuesSet[value] = true
			}
		}
	}

	if len(valuesSet) <= count {
		values := make([]int, 0, len(valuesSet))

		for value, _ := range valuesSet {
			values = append(values, value)
		}

		return values
	}

	return nil
}

// Return all hidden values (count) of the passed lookups or if none, then nil.
func findHiddenNValues(count int, lookups []*map[int]bool, otherLookups []*map[int]bool) []int {
	// We need at least one lookup with more than [count] values,
	// otherwise the algorithm won't have anything to reduce afterwards which would
	// not make any sense.
	maxCount := 0
	for _, lookupPtr := range lookups {
		if c := countPossibleValues(lookupPtr); c > maxCount {
			maxCount = c
		}
	}

	// Check if at least [count]+1 values in a lookup.
	if count >= maxCount {
		return nil
	}

	// Check values in common of lookups
	occurrences1 := countValueOccurrences(lookups)
	occurrences2 := countValueOccurrences(otherLookups)

	// Values of occurrences2 mustn't occur in occurrences1 -> Subtract
	occurrences := make(map[int]int)
	for value := 1; value <= model.SudokuSize; value++ {
		c1, hasValue1 := (*occurrences1)[value]
		c2, hasValue2 := (*occurrences2)[value]

		if hasValue1 && !hasValue2 {
			c := c1 - c2
			if c > 0 {
				occurrences[value] = c
			}
		}
	}

	if len(occurrences) == count {
		// Success! Collect values and return.
		values := make([]int, 0, count)
		for value, _ := range occurrences {
			values = append(values, value)
		}

		return values
	}

	return nil
}
