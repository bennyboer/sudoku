package util

import (
	"github.com/ob-algdatii-ss19/leistungsnachweis-sudo/model"
)

// Prepare a lookup of possible values per Sudoku cell.
// It will return a two-dimensional matrix with lookup tables where for each value either true or false is given.
// True means -> value is possible
// False means -> value is not possible
func PreparePossibleValueLookup(sudoku *model.Sudoku) *[][]*map[int]bool {
	matrix := make([][]*map[int]bool, model.SudokuSize)

	for i := 0; i < model.SudokuSize; i++ {
		matrix[i] = make([]*map[int]bool, model.SudokuSize)
	}

	cellMatrix := sudoku.Cells

	// Fill with possible values
	for row := 0; row < len(cellMatrix); row++ {
		rowCells := cellMatrix[row]

		for column := 0; column < len(rowCells); column++ {
			cell := rowCells[column]

			if cell.IsEmpty() {
				possibleValuesLookup := make(map[int]bool)

				// Fill possible values in the lookup
				for _, value := range cell.PossibleValues() {
					possibleValuesLookup[value] = true
				}

				// Fill impossible values in the lookup
				for value := 1; value <= 9; value++ {
					_, ok := possibleValuesLookup[value]

					if !ok {
						possibleValuesLookup[value] = false
					}
				}

				matrix[row][column] = &possibleValuesLookup
			}
		}
	}

	return &matrix
}
