package backtracking

import (
	"github.com/ob-algdatii-ss19/leistungsnachweis-sudo/model"
	"github.com/ob-algdatii-ss19/leistungsnachweis-sudo/solver/backtracking/strategy"
)

// Solver using the backtracking algorithm.
type Solver struct {
	CellChooserType strategy.CellChoosingStrategyType
}

// Solve the passed Sudoku.
// Return if the Sudoku was solvable.
func (s *Solver) Solve(sudoku *model.Sudoku) (bool, error) {
	cellChooser, e := strategy.Create(s.CellChooserType)
	if e != nil {
		return false, e
	}

	emptyCells := *cellChooser.Get(sudoku)

	for i := 0; i < len(emptyCells); {
		cell := emptyCells[i]

		success := false
		for value := cell.Value() + 1; value <= 9; value++ {
			cell.SetValue(value)

			if !cell.HasCollision() {
				success = true
				break
			}
		}

		if success {
			i++ // Next empty cell
		} else {
			// "Backtrack": Set value again to 0 and try again with last cell
			cell.SetValue(0)

			// First and foremost, check if backtracking is possible
			if i > 0 {
				i-- // Backtracking is possible -> Get last cell
			} else {
				return false, nil // Not solvable
			}
		}
	}

	return true, nil
}
