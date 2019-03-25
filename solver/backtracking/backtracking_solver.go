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
	emptyCells, valuesLeft, e := s.getConfigurationFor(sudoku)
	if e != nil {
		return false, e
	}

	return s.solve(sudoku, emptyCells, valuesLeft, 0), nil
}

// Check if the passed Sudoku is unique.
func (s *Solver) HasUniqueSolution(sudoku model.Sudoku) (bool, error) {
	// Try to find two solutions -> No unique solution.
	emptyCells, valuesLeft, e := s.getConfigurationFor(&sudoku)
	if e != nil {
		return false, e
	}

	solvable := s.solve(&sudoku, emptyCells, valuesLeft, 0)
	if !solvable {
		return false, nil
	}

	solvable = s.solve(&sudoku, emptyCells, valuesLeft, len(*emptyCells)-1)

	return !solvable, nil
}

// Get a configuration for the algorithm for the passed Sudoku.
func (s *Solver) getConfigurationFor(sudoku *model.Sudoku) (*[]*model.SudokuCell, *[][]int, error) {
	cellChooser, e := strategy.Create(s.CellChooserType)
	if e != nil {
		return nil, nil, e
	}

	// Get empty cells in the passed Sudoku
	emptyCells := *cellChooser.Get(sudoku)

	// Generate all values left to try for each empty cell
	rand.Seed(time.Now().UnixNano())         // Seed random number generator
	values := make([][]int, len(emptyCells)) // What values are left to be tried for the cell
	for i := 0; i < len(emptyCells); i++ {
		values[i] = rand.Perm(model.SudokuSize)
	}

	return &emptyCells, &values, nil
}

// Solve Sudoku by filling the passed empty cells in the Sudoku.
// Pass the empty cells of the Sudoku to solve and the values which are left to try for each cell.
// len(emptyCells) == len(valuesLeftPtr)
// Specify the startCellIndex (range [0; len(emptyCells) - 1)) in case you want to alter the algorithm behavior,
// otherwise just pass 0.
func (s *Solver) solve(sudoku *model.Sudoku,
	emptyCellsPtr *[]*model.SudokuCell,
	valuesLeftPtr *[][]int,
	startCellIndex int) bool {
	emptyCells := *emptyCellsPtr
	values := *valuesLeftPtr

	i := startCellIndex
	for ; i < len(emptyCells); {
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
				return false // Not solvable
			}
		}
	}

	return sudoku.IsValid()
}
