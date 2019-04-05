package pattern

import (
	"github.com/ob-algdatii-ss19/leistungsnachweis-sudo/model"
)

// Implementation of the hidden single pattern.
// It searches for each block for values which can only be applied to one single field!
type HiddenSingle struct{}

// Apply pattern on Sudoku.
// Search block by block values which can only be applied to one single cell.
func (p *HiddenSingle) Apply(sudoku *model.Sudoku, possibleValuesRef *[][]*map[int]bool) (changed bool) {
	changed = false
	pv := *possibleValuesRef

	for block := 0; block < model.SudokuSize; block++ {
		startRow := block / model.BlockSize * model.BlockSize;
		startColumn := (block * model.BlockSize) % model.SudokuSize;
		for value := 1; value < model.SudokuSize; value++ {
			isValueUnique := false
			var uniqueRow int
			var uniqueColumn int

			// Check if value is unique in block
		Outer:
			for row := startRow; row < startRow+model.BlockSize; row++ {
				for column := startColumn; column < startColumn+model.BlockSize; column++ {
					if pv[row][column] != nil {
						possible, ok := (*pv[row][column])[value]

						if ok && possible {
							if isValueUnique {
								isValueUnique = false
								break Outer // Value occurs more than once -> No need to continue
							} else {
								isValueUnique = true
								uniqueRow = row
								uniqueColumn = column
							}
						}
					}
				}
			}

			if isValueUnique {
				// Value is unique in block! Fill the cell!
				updateValueInSudokuAndLookup(sudoku, possibleValuesRef, uniqueRow, uniqueColumn, value)

				changed = true
			}
		}
	}

	return
}
