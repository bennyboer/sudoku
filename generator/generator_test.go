package generator

import (
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

	sudoku, err := generator.Generate(0.8)
	if err != nil {
		t.Errorf("expected no error")
		return
	}

	if sudoku.IsComplete() {
		t.Errorf("sudoku is complete")
	}
}
