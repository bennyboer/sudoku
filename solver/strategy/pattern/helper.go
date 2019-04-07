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
