package generator

import (
	"github.com/ob-algdatii-ss19/leistungsnachweis-sudo/model"
	"github.com/ob-algdatii-ss19/leistungsnachweis-sudo/solver/backtracking"
	"github.com/ob-algdatii-ss19/leistungsnachweis-sudo/solver/backtracking/strategy"
	"math/rand"
)

func (sg *SudokuGeneratorBacktracking) Generate(difficulty float64) (*model.Sudoku, error) {
	var lastState *model.Sudoku
	solver := backtracking.Solver{CellChooserType: strategy.Linear}
	sudoku := model.EmptySudoku()
	_, err := solver.Solve(sudoku)
	if err != nil {
		return nil, err
	}

	deletioncount := int((81 - 17) * difficulty)

	for i := 0; i < deletioncount; i++ {
		lastState = sudoku
		sudoku.Cells[rand.Intn(9)][rand.Intn(9)].SetValue(0)
		sudokuCopy, _ := model.LoadSudoku(sudoku.SaveSudoku())

		success, err := solver.Solve(sudokuCopy)
		if !success && err != nil {
			i--
			sudoku = lastState
		}
	}

	return sudoku, nil
}

func (generator *SudokuGeneratorDifficulty) Generate(difficulty float64) (*model.Sudoku, error) {
	solver := backtracking.Solver{CellChooserType: strategy.Linear}
	sudoku := model.EmptySudoku()
	_, err := solver.Solve(sudoku)
	if err != nil {
		return nil, err
	}

	return sudoku, nil
}
