package solver

import "github.com/ob-algdatii-ss19/leistungsnachweis-sudo/model"

/// Base interface for all Sudoku solvers.
type SudokuSolver interface {
	// Solve the passed Sudoku.
	Solve(sudoku *model.Sudoku)
}
