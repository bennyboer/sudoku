package pattern

import "github.com/ob-algdatii-ss19/leistungsnachweis-sudo/model"

// Helper function to update a value in the passed Sudoku as well in the possible values lookup matrix.
func updateValueInSudokuAndLookup(sudoku *model.Sudoku, possibleValuesRef *[][]*map[int]bool, row int, column int, value int) {
	pv := *possibleValuesRef

	// Fill cell in Sudoku and update possible value lookup
	cell := sudoku.Cells[row][column]
	cell.SetValue(value)

	// Update possible value lookup
	if pv[row][column] != nil {
		(*pv[row][column])[value] = false
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
