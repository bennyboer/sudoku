package strategy

import (
	"github.com/ob-algdatii-ss19/leistungsnachweis-sudo/model"
	"testing"
)

func TestSolver_Solve(t *testing.T) {
	sudoku, _ := model.LoadSudoku(&[9][9]int{
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

	solver := Solver{}

	_, _ = solver.Solve(sudoku)
}
