package generator

import "testing"

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
