package pattern

import (
	"github.com/ob-algdatii-ss19/leistungsnachweis-sudo/model"
	"github.com/ob-algdatii-ss19/leistungsnachweis-sudo/solver/backtracking"
	"github.com/ob-algdatii-ss19/leistungsnachweis-sudo/solver/backtracking/strategy"
)

// Implementation of the fallback pattern, which is used as fallback, in
// case none of the provided pattern is able to move forward.
// It will simply guess on value with the backtracking solver.
type Fallback struct{}

// Apply pattern on the passed sudoku.
func (p *Fallback) Apply(sudoku *model.Sudoku, possibleValuesRef *[][]*map[int]bool) (changed bool) {
	solver := backtracking.Solver{CellChooserType: strategy.Linear}

	sudokuToModify, err := model.LoadSudoku(sudoku.SaveSudoku())
	if err != nil {
		return false
	}

	solvable, err := solver.Solve(sudokuToModify)
	if err != nil || !solvable {
		return false
	}

	pv := *possibleValuesRef

	for row := 0; row < len(pv); row++ {
		possibleValuesRow := pv[row]
		for column := 0; column < len(possibleValuesRow); column++ {
			pvLookup := possibleValuesRow[column]
			if pvLookup != nil {
				// Found not filled value -> Fill
				updateValueInSudokuAndLookup(sudoku, possibleValuesRef, row, column, sudokuToModify.Cells[row][column].Value(), false)
				return true
			}
		}
	}

	return false
}
