package solver

import (
	"github.com/ob-algdatii-ss19/leistungsnachweis-sudo/model"
	"github.com/ob-algdatii-ss19/leistungsnachweis-sudo/solver/strategy"
)

// Measures Sudoku difficulty.
type DifficultyMeasurer interface {
	// Measure the difficulty of the passed Sudoku.
	Measure(sudoku *model.Sudoku) (float64, error)
}

// Measure the difficulty of the passed Sudoku.
func MeasureDifficulty(sudoku *model.Sudoku) (float64, error) {
	return createDifficultyMeasurer().Measure(sudoku)
}

// Create the default difficulty measurer.
func createDifficultyMeasurer() DifficultyMeasurer {
	return &strategy.Solver{}
}
