package solver

import (
	"github.com/ob-algdatii-ss19/leistungsnachweis-sudo/model"
	"github.com/ob-algdatii-ss19/leistungsnachweis-sudo/solver/backtracking"
	"github.com/ob-algdatii-ss19/leistungsnachweis-sudo/solver/backtracking/strategy"
)

// Checks Sudoku for a unique solution.
type UniqueSolutionChecker interface {
	// Check if the passed Sudoku has a unique solution.
	HasUniqueSolution(sudoku model.Sudoku) (bool, error)
}

// Check if the passed Sudoku has a unique solution.
func HasUniqueSolution(sudoku *model.Sudoku) (bool, error) {
	return createUniqueSolutionChecker().HasUniqueSolution(*sudoku)
}

// Create an unique solution checker.
func createUniqueSolutionChecker() UniqueSolutionChecker {
	return &backtracking.Solver{CellChooserType: strategy.Linear}
}
