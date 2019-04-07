package strategy

import (
	"github.com/ob-algdatii-ss19/leistungsnachweis-sudo/model"
	"testing"
)

func TestSolver_SolveSimple(t *testing.T) {
	sudoku, _ := model.LoadSudoku(&[][]int{
		{6, 7, 9, 1, 0, 4, 3, 5, 2},
		{0, 4, 0, 6, 3, 7, 0, 8, 1},
		{3, 1, 0, 9, 2, 0, 4, 6, 0},
		{7, 9, 1, 0, 4, 6, 2, 0, 8},
		{4, 6, 3, 2, 0, 8, 0, 1, 5},
		{2, 0, 5, 0, 1, 3, 6, 4, 0},
		{0, 5, 6, 4, 7, 0, 0, 2, 3},
		{8, 2, 0, 3, 0, 9, 1, 0, 6},
		{1, 0, 7, 0, 6, 2, 5, 9, 4},
	})

	solver := Solver{}

	solvable, e := solver.Solve(sudoku)

	if e != nil {
		t.Errorf("Expected solver to not throw an error")
	}

	if !solvable {
		t.Errorf("Expected solver to solve the Sudoku properly")
	}
}

func TestSolver_SolveAdvanced(t *testing.T) {
	sudoku, _ := model.LoadSudoku(&[][]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{4, 0, 0, 0, 7, 0, 6, 0, 0},
		{0, 7, 0, 2, 0, 4, 3, 0, 0},
		{0, 0, 0, 9, 0, 0, 0, 8, 4},
		{0, 6, 0, 0, 0, 0, 5, 9, 2},
		{5, 9, 0, 0, 0, 0, 0, 0, 0},
		{0, 4, 0, 8, 2, 0, 9, 0, 1},
		{0, 0, 0, 1, 0, 0, 0, 0, 0},
		{6, 5, 1, 7, 0, 0, 0, 0, 0},
	})

	solver := Solver{}

	solvable, e := solver.Solve(sudoku)

	if e != nil {
		t.Errorf("Expected solver to not throw an error")
	}

	if !solvable {
		t.Errorf("Expected solver to solve the Sudoku properly")
	}
}

func TestSolver_SolveComplicated(t *testing.T) {
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

	solver := Solver{}

	solvable, e := solver.Solve(sudoku)

	if e != nil {
		t.Errorf("Expected solver to not throw an error")
	}

	if !solvable {
		t.Errorf("Expected solver to solve the Sudoku properly")
	}
}
