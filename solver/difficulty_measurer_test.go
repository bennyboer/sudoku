package solver

import (
	"github.com/ob-algdatii-ss19/leistungsnachweis-sudo/model"
	"testing"
)

func TestMeasureDifficulty(t *testing.T) {
	sudoku, _ := model.LoadSudoku(&[][]int{
		{2, 6, 0, 0, 0, 0, 0, 0, 0},
		{1, 0, 0, 8, 0, 0, 0, 0, 0},
		{0, 0, 0, 9, 6, 2, 0, 0, 0},
		{0, 0, 0, 3, 0, 0, 7, 0, 0},
		{0, 0, 0, 0, 5, 0, 0, 0, 8},
		{0, 0, 2, 7, 4, 0, 0, 9, 0},
		{4, 0, 0, 0, 0, 0, 0, 2, 0},
		{7, 0, 3, 0, 0, 6, 0, 0, 0},
		{0, 0, 0, 0, 3, 0, 0, 8, 1},
	})

	difficulty, err := MeasureDifficulty(sudoku)
	if err != nil {
		t.Errorf("expected no error")
	}

	if difficulty < 0.0 || difficulty > 1.0 {
		t.Errorf("difficulty %f out of range [0.0; 1.0]", difficulty)
	}

	if !sudoku.Cells[0][2].IsEmpty() {
		t.Errorf("expected sudoku to not have changed!")
	}
}
