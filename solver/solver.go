package solver

import (
	"errors"
	"github.com/ob-algdatii-ss19/leistungsnachweis-sudo/model"
	"github.com/ob-algdatii-ss19/leistungsnachweis-sudo/solver/backtracking"
	s "github.com/ob-algdatii-ss19/leistungsnachweis-sudo/solver/backtracking/strategy"
	"github.com/ob-algdatii-ss19/leistungsnachweis-sudo/solver/strategy"
)

/// Base interface for all Sudoku solvers.
type SudokuSolver interface {
	// Solve the passed Sudoku.
	// Return if the Sudoku was solvable.
	Solve(sudoku *model.Sudoku) (bool, error)
}

// Get an algorithm by its name.
func GetAlgorithmForName(name string) (SudokuSolver, error) {
	for n, solver := range *AllSolverAlgorithms() {
		if n == name {
			return solver, nil
		}
	}

	return nil, errors.New("algorithm name unknown")
}

// Return all available solver algorithms.
func AllSolverAlgorithms() *map[string]SudokuSolver {
	aMap := make(map[string]SudokuSolver)

	aMap["backtracking"] = &backtracking.Solver{CellChooserType: s.Linear}
	aMap["strategy"] = &strategy.Solver{}

	return &aMap
}
