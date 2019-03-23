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
func (s *Solver) Solve(sudoku *model.Sudoku) error {
	var cellChooser strategy.CellChoosingStrategy
	if str, e := strategy.Create(s.CellChooserType); e != nil {
		return e
	} else {
		cellChooser = str
	}

	cellChooser.Initialize(sudoku)

	nextEmptyCell, e := cellChooser.FindNext()
	if e != nil {
		return e
	}

	for nextEmptyCell != nil {
		// TODO Implement algorithm
	}

	return nil
}
