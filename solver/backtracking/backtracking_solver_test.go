package backtracking

import (
	"github.com/ob-algdatii-ss19/leistungsnachweis-sudo/model"
	"github.com/ob-algdatii-ss19/leistungsnachweis-sudo/solver/backtracking/strategy"
	"testing"
)

func TestSolver_Solve_Linear(t *testing.T) {
	sudoku, _ := model.LoadSudoku(&[9][9]int{
		{0, 1, 2, 0, 0, 0, 5, 7, 0},
		{6, 0, 0, 5, 0, 1, 0, 0, 4},
		{4, 0, 0, 0, 2, 0, 0, 0, 8},
		{0, 2, 0, 0, 1, 0, 0, 5, 0},
		{0, 0, 4, 9, 0, 7, 8, 0, 0},
		{0, 7, 0, 0, 8, 0, 0, 1, 0},
		{7, 0, 0, 0, 9, 0, 0, 0, 5},
		{5, 0, 0, 4, 0, 8, 0, 0, 6},
		{0, 3, 8, 0, 0, 0, 9, 4, 0},
	})
	sudokuSolution := [9][9]int{
		{9, 1, 2, 8, 4, 6, 5, 7, 3},
		{6, 8, 3, 5, 7, 1, 2, 9, 4},
		{4, 5, 7, 3, 2, 9, 1, 6, 8},
		{8, 2, 9, 6, 1, 3, 4, 5, 7},
		{1, 6, 4, 9, 5, 7, 8, 3, 2},
		{3, 7, 5, 2, 8, 4, 6, 1, 9},
		{7, 4, 6, 1, 9, 2, 3, 8, 5},
		{5, 9, 1, 4, 3, 8, 7, 2, 6},
		{2, 3, 8, 7, 6, 5, 9, 4, 1},
	}

	solver := Solver{CellChooserType: strategy.Linear}

	solvable, e := solver.Solve(sudoku)
	if !solvable || e != nil {
		t.Errorf("Expected Sudoku to be solvable")
	}

	if !sudoku.IsValid() {
		t.Errorf("Expected solved Sudoku to be valid!")
	}

	for row := 0; row < 9; row++ {
		for column := 0; column < 9; column++ {
			expected := sudokuSolution[row][column]
			got := sudoku.Cells[row][column].Value()

			if got != expected {
				t.Errorf("Expected value at (%d, %d) to be %d; got %d", row, column, expected, got)
			}
		}
	}
}

func TestSolver_Solve_Random(t *testing.T) {
	sudoku, _ := model.LoadSudoku(&[9][9]int{
		{9, 1, 2, 0, 0, 0, 5, 7, 3},
		{6, 8, 0, 5, 7, 1, 2, 9, 4},
		{4, 5, 0, 0, 2, 9, 1, 6, 8},
		{8, 2, 0, 0, 1, 3, 4, 5, 7},
		{1, 6, 4, 9, 5, 7, 8, 3, 2},
		{3, 7, 5, 2, 8, 4, 6, 1, 9},
		{7, 4, 6, 1, 9, 2, 3, 8, 5},
		{5, 9, 1, 4, 3, 8, 7, 2, 6},
		{2, 3, 8, 7, 6, 5, 9, 4, 1},
	})
	sudokuSolution := [9][9]int{
		{9, 1, 2, 8, 4, 6, 5, 7, 3},
		{6, 8, 3, 5, 7, 1, 2, 9, 4},
		{4, 5, 7, 3, 2, 9, 1, 6, 8},
		{8, 2, 9, 6, 1, 3, 4, 5, 7},
		{1, 6, 4, 9, 5, 7, 8, 3, 2},
		{3, 7, 5, 2, 8, 4, 6, 1, 9},
		{7, 4, 6, 1, 9, 2, 3, 8, 5},
		{5, 9, 1, 4, 3, 8, 7, 2, 6},
		{2, 3, 8, 7, 6, 5, 9, 4, 1},
	}

	solver := Solver{CellChooserType: strategy.Random}

	solvable, e := solver.Solve(sudoku)
	if !solvable || e != nil {
		t.Errorf("Expected Sudoku to be solvable")
	}

	if !sudoku.IsValid() {
		t.Errorf("Expected solved Sudoku to be valid!")
	}

	for row := 0; row < 9; row++ {
		for column := 0; column < 9; column++ {
			expected := sudokuSolution[row][column]
			got := sudoku.Cells[row][column].Value()

			if got != expected {
				t.Errorf("Expected value at (%d, %d) to be %d; got %d", row, column, expected, got)
			}
		}
	}
}

func TestSolver_Solve_UnknownCellChoosingStrategy(t *testing.T) {
	ty := strategy.CellChoosingStrategyType(4)
	solver := Solver{CellChooserType: ty}

	sudoku := model.EmptySudoku()

	_, e := solver.Solve(sudoku)
	if e == nil {
		t.Errorf("Expected to return error")
	}
}

func TestSolver_Solve_NotSolvable(t *testing.T) {
	sudoku, _ := model.LoadSudoku(&[9][9]int{
		{5, 1, 6, 8, 4, 9, 7, 3, 2},
		{3, 0, 7, 6, 0, 5, 0, 0, 0},
		{8, 0, 9, 7, 0, 0, 0, 6, 5},
		{1, 3, 5, 0, 6, 0, 9, 0, 7},
		{4, 7, 2, 5, 9, 1, 0, 0, 6},
		{9, 6, 8, 3, 7, 0, 0, 5, 0},
		{2, 5, 3, 1, 8, 6, 0, 7, 4},
		{6, 8, 4, 2, 0, 7, 5, 0, 0},
		{7, 9, 1, 0, 5, 0, 6, 0, 8},
	})

	solver := Solver{CellChooserType: strategy.Linear}

	solvable, _ := solver.Solve(sudoku)
	if solvable {
		t.Errorf("Expected Sudoku to be unsolvable")
	}

	if !sudoku.IsValid() {
		t.Errorf("Expected solved Sudoku to be valid!")
	}
}

func TestSolver_HasUniqueSolution(t *testing.T) {
	solver := Solver{CellChooserType: strategy.Linear}

	empty := model.EmptySudoku()
	hasUniqueSolution, _ := solver.HasUniqueSolution(*empty)

	if hasUniqueSolution {
		t.Errorf("Expected empty Sudoku to have a lot of solutions and not an unique one")
	}

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
	hasUniqueSolution, _ = solver.HasUniqueSolution(*sudoku)

	if !hasUniqueSolution {
		t.Errorf("Expected Sudoku to only have one unique solution.")
	}
}

func TestSolver_HasUniqueSolution_Errors(t *testing.T) {
	solver := Solver{CellChooserType: strategy.CellChoosingStrategyType(5)}

	empty := model.EmptySudoku()
	_, e := solver.HasUniqueSolution(*empty)
	if e == nil {
		t.Errorf("Expected to throw error because of incorrect cell choosing strategy type")
	}

	notSolvableSudoku, _ := model.LoadSudoku(&[9][9]int{
		{1, 1, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1, 0, 1},
		{1, 1, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1, 1, 1},
	})
	solver = Solver{CellChooserType: strategy.Linear}
	solvable, _ := solver.HasUniqueSolution(*notSolvableSudoku)
	if solvable {
		t.Errorf("Expected Sudoku to be unsolvable")
	}
}
