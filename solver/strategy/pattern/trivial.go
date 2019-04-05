package pattern

import (
	"github.com/ob-algdatii-ss19/leistungsnachweis-sudo/model"
)

// Implementation of the trivial pattern.
// It will just look into the possible values and fill all values which are trivially missing,
// meaning:
// - If only one value is missing in a ROW
// - If only one value is missing in a COLUMN
// - If only one value is missing in a BLOCK
type Trivial struct{}

// Apply pattern on Sudoku.
func (p *Trivial) Apply(sudoku *model.Sudoku, possibleValuesRef *[][]*map[int]bool) (changed bool) {
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
					// Fill cell in Sudoku and update possible value lookup
					cell := sudoku.Cells[row][column]
					cell.SetValue(onlyPossibleValue)

					// Update possible value lookup
					(*pv[row][column])[onlyPossibleValue] = false
					// Update neighbour lookups
					for _, neighbour := range cell.Neighbours().All {
						position := neighbour.Position()

						// Set as "no more possible" in lookup
						lookupPtr := pv[position.Row][position.Column]
						if lookupPtr != nil {
							(*lookupPtr)[onlyPossibleValue] = false // Mark the value as "no more possible"
						}
					}

					changed = true
				}
			}
		}
	}

	return
}
