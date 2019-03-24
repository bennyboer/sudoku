package backtracking

import (
	"github.com/ob-algdatii-ss19/leistungsnachweis-sudo/model"
	"github.com/ob-algdatii-ss19/leistungsnachweis-sudo/solver/backtracking/strategy"
	"math/rand"
	"time"
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

	rand.Seed(time.Now().UnixNano())         // Seed random number generator
	values := make([][]int, len(emptyCells)) // What values are left to be tried for the cell
	for i := 0; i < len(emptyCells); i++ {
		values[i] = rand.Perm(model.SudokuSize)
	}

	for i := 0; i < len(emptyCells); {
		cell := emptyCells[i]

		success := false
		for a := 0; a < len(values[i]); a++ {
			value := values[i][a] + 1

			cell.SetValue(value)

			if !cell.HasCollision() {
				success = true
				values[i] = values[i][a+1:] // Update values left to try for this cell
				break
			}
		}

		if success {
			i++ // Next empty cell
		} else {
			// "Backtrack": Set value again to 0 and try again with last cell
			cell.SetValue(0)
			values[i] = rand.Perm(model.SudokuSize)

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
