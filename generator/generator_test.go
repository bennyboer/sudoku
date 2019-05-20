package generator

import (
	"github.com/ob-algdatii-ss19/leistungsnachweis-sudo/solver"
	"math"
	"testing"
)

func TestSudokuGeneratorBackTracking_Generate(t *testing.T) {
	generator := NewBacktrackingGenerator()

	sudoku, err := generator.Generate(0.1)
	if err != nil {
		t.Errorf("expected no error")
		return
	}

	if sudoku.IsComplete() {
		t.Errorf("Sudoku is complete")
	}
}

func TestSudokuGeneratorDifficulty_Generate(t *testing.T) {
	generator := NewDifficultyGenerator()

	sudoku, err := generator.Generate(0.1)
	if err != nil {
		t.Errorf("expected no error")
		return
	}

	if sudoku.IsComplete() {
		t.Errorf("sudoku is complete")
	}

	diff, _ := solver.MeasureDifficulty(sudoku)

	if math.Abs(diff-0.1) > 0.05 {
		t.Errorf("difficulty doesn't match. Expected: %f, Got %f", 0.1, diff)
	}
}
