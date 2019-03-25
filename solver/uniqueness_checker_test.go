package solver

import (
	"github.com/ob-algdatii-ss19/leistungsnachweis-sudo/model"
	"testing"
)

func TestHasUniqueSolution(t *testing.T) {
	sudoku, _ := model.LoadSudoku(&[9][9]int{
		{0, 0, 5, 3, 0, 0, 0, 0, 0},
		{8, 0, 0, 0, 0, 0, 0, 2, 0},
		{0, 7, 0, 0, 1, 0, 5, 0, 0},
		{4, 0, 0, 0, 0, 5, 3, 0, 0},
		{0, 1, 0, 0, 7, 0, 0, 0, 6},
		{0, 0, 3, 2, 0, 0, 0, 8, 0},
		{0, 6, 0, 5, 0, 0, 0, 0, 9},
		{0, 0, 4, 0, 0, 0, 0, 3, 0},
		{0, 0, 0, 0, 0, 9, 7, 0, 0},
	})
	hasUniqueSolution, _ := HasUniqueSolution(sudoku)

	if !hasUniqueSolution {
		t.Errorf("Expected Sudoku to have an unique solution");
	}
}
