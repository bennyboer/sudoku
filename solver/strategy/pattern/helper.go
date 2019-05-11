package pattern

import (
	"errors"
	"github.com/ob-algdatii-ss19/leistungsnachweis-sudo/model"
)

// A change of the updateValueInSudokuAndLookup method in the lookup.
type lookupUpdateChange struct {
	row           int
	column        int
	possibleValue int
}

// Helper function to update a value in the passed Sudoku as well in the possible values lookup matrix.
func updateValueInSudokuAndLookup(sudoku *model.Sudoku, possibleValuesRef *[][]*map[int]bool, row int, column int, value int, saveChanges bool) *[]lookupUpdateChange {
	pv := *possibleValuesRef

	var changes []lookupUpdateChange = nil
	if saveChanges {
		changes = make([]lookupUpdateChange, 0, model.NeighbourCount)
	}

	// Fill cell in Sudoku and update possible value lookup
	cell := sudoku.Cells[row][column]
	cell.SetValue(value)

	// Update possible value lookup
	if saveChanges {
		// First save all possible values in changes slice
		for v, p := range *pv[row][column] {
			if p {
				changes = append(changes, lookupUpdateChange{
					row:           row,
					column:        column,
					possibleValue: v,
				})
			}
		}
	}
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

			if saveChanges {
				// Save change
				changes = append(changes, lookupUpdateChange{
					row:           position.Row,
					column:        position.Column,
					possibleValue: value,
				})
			}
		}
	}

	return &changes
}

// Undo the changes done with updateValueInSudokuAndLookup.
func undoValueUpdateInSudokuAndLookup(sudoku *model.Sudoku, possibleValuesRef *[][]*map[int]bool, row int, column int, value int, lookupChanges *[]lookupUpdateChange) error {
	if lookupChanges == nil {
		return errors.New("expected lookup changes slice and not a nil pointer")
	}

	pv := *possibleValuesRef

	// Undo cell filling
	cell := sudoku.Cells[row][column]
	cell.SetValue(0)

	// Undo cell possible lookup deletion
	newPVLookup := make(map[int]bool)
	for v := 1; v <= model.SudokuSize; v++ {
		newPVLookup[v] = false
	}
	pv[row][column] = &newPVLookup

	// Undo updated possible value lookups
	for i := 0; i < len(*lookupChanges); i++ {
		change := (*lookupChanges)[i]

		pvLookup := pv[change.row][change.column]
		if pvLookup != nil {
			(*pvLookup)[change.possibleValue] = true
		}
	}

	return nil
}

// Function processing a unit (Row, column or block) of possible value maps.
type unitFunction func([]*map[int]bool)

// Process the passed function for each unit (Row, column or block) of possible value maps.
func forEachUnit(fn unitFunction, possibleValuesRef *[][]*map[int]bool) {
	// Rows
	for row := 0; row < model.SudokuSize; row++ {
		fn(getRowPossibleValues(row, possibleValuesRef, false))
	}

	// Columns
	for column := 0; column < model.SudokuSize; column++ {
		fn(getColumnPossibleValues(column, possibleValuesRef, false))
	}

	// Blocks
	for block := 0; block < model.SudokuSize; block++ {
		fn(getBlockPossibleValues(block, possibleValuesRef, false))
	}
}

// Get all possible values for a block in range [0; 8]
func getBlockPossibleValues(block int, possibleValuesRef *[][]*map[int]bool, includeNilPtrs bool) []*map[int]bool {
	pv := *possibleValuesRef

	possibleBlockValues := make([]*map[int]bool, 0, model.SudokuSize)

	startRow := block / model.BlockSize * model.BlockSize
	startColumn := (block * model.BlockSize) % model.SudokuSize
	for row := startRow; row < startRow+model.BlockSize; row++ {
		for column := startColumn; column < startColumn+model.BlockSize; column++ {
			if pv[row][column] != nil || includeNilPtrs {
				possibleBlockValues = append(possibleBlockValues, pv[row][column])
			}
		}
	}

	return possibleBlockValues
}

// Get all possible values for a row in range [0; 8]
func getRowPossibleValues(row int, possibleValuesRef *[][]*map[int]bool, includeNilPtrs bool) []*map[int]bool {
	pv := *possibleValuesRef

	possibleRowValues := make([]*map[int]bool, 0, model.SudokuSize)

	for column := 0; column < model.SudokuSize; column++ {
		if pv[row][column] != nil || includeNilPtrs {
			possibleRowValues = append(possibleRowValues, pv[row][column])
		}
	}

	return possibleRowValues
}

// Get all possible values for a column in range [0; 8]
func getColumnPossibleValues(column int, possibleValuesRef *[][]*map[int]bool, includeNilPtrs bool) []*map[int]bool {
	pv := *possibleValuesRef

	possibleColumnValues := make([]*map[int]bool, 0, model.SudokuSize)

	for row := 0; row < model.SudokuSize; row++ {
		if pv[row][column] != nil || includeNilPtrs {
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

// Calculate the binomial coefficient of n and k.
func binomialCoefficient(n, k int) int {
	result := 1

	// binomialCoefficient(n, k) = binomialCoefficient(n, n - k)
	if k > n-k {
		k = n - k
	}

	for i := 0; i < k; i++ {
		result *= n - i
		result /= i + 1
	}

	return result
}

// Get all possible combinations of [count] rows out of [all].
func getCombinations(count, all int) [][]int {
	combinationCount := binomialCoefficient(all, count)
	combinations := make([][]int, combinationCount)

	counter := make([]int, count)
	// Initialize counter
	for i := 0; i < count; i++ {
		counter[i] = i
	}

	// Collect all combinations.
	for i := 0; i < combinationCount; i++ {
		combinations[i] = make([]int, count)
		for a := 0; a < count; a++ {
			combinations[i][a] = counter[a]
		}

		// Increase counter
		if i != combinationCount-1 {
			counterIndex := count - 1
			counter[counterIndex]++
			for counter[counterIndex] > model.SudokuSize-(count-counterIndex) {
				counterIndex--
				counter[counterIndex]++
			}
			for a := counterIndex + 1; a < count; a++ {
				counter[a] = counter[counterIndex] + (a - counterIndex)
			}
		}
	}

	return combinations
}
