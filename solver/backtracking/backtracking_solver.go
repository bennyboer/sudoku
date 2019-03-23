package backtracking

import (
	"github.com/ob-algdatii-ss19/leistungsnachweis-sudo/model"
	"github.com/ob-algdatii-ss19/leistungsnachweis-sudo/solver/backtracking/strategy"
)

// Solver using the backtracking algorithm.
type BacktrackingSolver struct {
	AlgorithmType            AlgorithmType
	CellChoosingStrategyType strategy.CellChoosingStrategyType
}

// Solve the passed Sudoku.
func (s *BacktrackingSolver) Solve(sudoku *model.Sudoku) {

}
