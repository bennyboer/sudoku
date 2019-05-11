package pattern

import (
	"github.com/ob-algdatii-ss19/leistungsnachweis-sudo/model"
)

// Implementation of the naked single pattern.
// It will just look into the possible values and fill all values which are trivially missing,
// meaning:
// - If only one value is missing in a ROW
// - If only one value is missing in a COLUMN
// - If only one value is missing in a BLOCK
type NakedSingle struct{}

// Apply pattern on Sudoku.
func (p *NakedSingle) Apply(sudoku *model.Sudoku, possibleValuesRef *[][]*map[int]bool) (changed bool) {
	changed = false
	pv := *possibleValuesRef

	// Iterate through all possible values and fill missing ones.
	for row := 0; row < model.SudokuSize; row++ {
		for column := 0; column < model.SudokuSize; column++ {
			if pv[row][column] != nil {
				// Not filled yet -> check if only one possible value
				onlyOnePossibleValue := false
				onlyPossibleValue := -1

				for value, possible := range *pv[row][column] {
					if possible {
						if onlyPossibleValue != -1 {
							// Had already a value -> Abort
							onlyOnePossibleValue = false
							break
						}

						onlyPossibleValue = value
						onlyOnePossibleValue = true
					}
				}

				if onlyOnePossibleValue {
					updateValueInSudokuAndLookup(sudoku, possibleValuesRef, row, column, onlyPossibleValue, false)

					changed = true
				}
			}
		}
	}

	return
}
