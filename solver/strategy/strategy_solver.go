package strategy

import (
	"errors"
	"github.com/ob-algdatii-ss19/leistungsnachweis-sudo/model"
	"github.com/ob-algdatii-ss19/leistungsnachweis-sudo/solver/strategy/pattern"
	"github.com/ob-algdatii-ss19/leistungsnachweis-sudo/solver/strategy/util"
	"math"
)

// A sudoku needs at least 17 filled sudoku cells to be solvable!
const minFilledSudokuCells = 17

// A sudoku mustn't have more than 81 - 17 empty cells to be solvable
const maxEmptySudokuCells = model.SudokuSize*model.SudokuSize - minFilledSudokuCells

// Strategy solver using the Sudoku solving patterns in order
// to solve the Sudoku.
type Solver struct {
	// Difficulty of the last solved Sudoku.
	lastPassDifficulty *float64
}

// Solve the passed Sudoku using the strategy solver.
func (s *Solver) Solve(sudoku *model.Sudoku) (bool, error) {
	patterns := initPatterns()
	patternCount := len(patterns)
	fallback := pattern.Fallback{}

	possibleValueLookupRef := util.PreparePossibleValueLookup(sudoku)

	// Count empty Sudoku cells
	emptyCells := 0
	for row := 0; row < model.SudokuSize; row++ {
		for column := 0; column < model.SudokuSize; column++ {
			if sudoku.Cells[row][column].IsEmpty() {
				emptyCells++
			}
		}
	}

	maxPatternIndex := 0
	for patternIndex := 0; patternIndex < patternCount+1; {
		isFallback := patternIndex >= patternCount
		var changed bool
		if !isFallback {
			// Try the pattern
			changed = patterns[patternIndex].Apply(sudoku, possibleValueLookupRef)
		} else {
			changed = fallback.Apply(sudoku, possibleValueLookupRef)
		}

		if changed {
			if patternIndex > maxPatternIndex {
				maxPatternIndex = patternIndex
			}

			// Go back to the first pattern, since the Sudoku has been changed by the current pattern.
			// New opportunities may have risen.
			patternIndex = 0
		} else {
			patternIndex++
		}

		if sudoku.IsCompleteAndValid() {
			s.calculateDifficulty(emptyCells, maxPatternIndex, patternCount)
			return true, nil // Early exit because the solver finished its job
		}
	}

	return sudoku.IsCompleteAndValid(), nil
}

// Calculate the difficulty of the last solved Sudoku.
func (s *Solver) calculateDifficulty(emptyFields, maxPatternIndex, patternCount int) {
	difficulty := float64(maxPatternIndex) / float64(patternCount) * 0.8
	difficulty += float64(emptyFields) / float64(maxEmptySudokuCells) * 0.2

	difficulty = math.Min(1.0, math.Max(0.0, difficulty)) // Make sure it has the correct range

	s.lastPassDifficulty = &difficulty
}

// Measure difficulty of the passed Sudoku.
func (s *Solver) Measure(sudoku *model.Sudoku) (float64, error) {
	// Copy sudoku in order to not modify it
	c, err := model.LoadSudoku(sudoku.SaveSudoku())
	if err != nil {
		return 0, err
	}

	// Solve Sudoku with strategy solver, which is able to measure the difficulty.
	solvable, err := s.Solve(c)
	if err != nil {
		return 0, err
	}

	if !solvable {
		return 0, errors.New("could not measure the difficulty of the passed Sudoku, since it is NOT solvable")
	}

	if s.lastPassDifficulty == nil {
		return 0, errors.New("measuring difficulty did not deliver a result")
	}

	return *s.lastPassDifficulty, nil
}

// Initialize the patterns slice.
// The patterns are used to be applied on a Sudoku in order to solve it.
func initPatterns() []pattern.Pattern {
	return []pattern.Pattern{
		&pattern.NakedSingle{},
		&pattern.HiddenSingle{},
		&pattern.NakedPair{},
		&pattern.HiddenPair{},
		&pattern.NakedTriple{},
		&pattern.HiddenTriple{},
		&pattern.NakedQuadruple{},
		&pattern.HiddenQuadruple{},
		&pattern.RowBlockCheck{},
		&pattern.BlockRowCheck{},
		&pattern.XWing{},
	}
}
