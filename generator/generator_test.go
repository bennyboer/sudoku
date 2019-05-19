package generator

import "testing"

func TestSudokuGeneratorBackTracking_Generate(t *testing.T) {
	generator := NewBacktrackingGenerator()

	sudoku := generator.Generate(0.1)
	if sudoku.IsComplete() {
		t.Errorf("Sudoku is complete")
	}
}
