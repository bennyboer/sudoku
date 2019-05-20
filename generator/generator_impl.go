package generator

import (
	"errors"
	"github.com/ob-algdatii-ss19/leistungsnachweis-sudo/model"
	"github.com/ob-algdatii-ss19/leistungsnachweis-sudo/solver/backtracking"
	s "github.com/ob-algdatii-ss19/leistungsnachweis-sudo/solver/backtracking/strategy"
	"github.com/ob-algdatii-ss19/leistungsnachweis-sudo/solver/strategy"
	"math"
	"math/rand"
)

func (sg *SudokuGeneratorBacktracking) Generate(difficulty float64) (*model.Sudoku, error) {
	var lastState *model.Sudoku
	solver := backtracking.Solver{CellChooserType: s.Linear}
	sudoku := model.EmptySudoku()
	_, err := solver.Solve(sudoku)
	if err != nil {
		return nil, err
	}

	if difficulty > 1.0 || difficulty < 0 {
		return nil, errors.New("The difficulty must be between 0 and 1")
	}

	deletioncount := int((81 - 17) * difficulty)

	for i := 0; i < deletioncount; i++ {
		lastState = sudoku

		x := rand.Intn(9)
		y := rand.Intn(9)
		for sudoku.Cells[x][y].Value() == 0 {
			x = rand.Intn(9)
			y = rand.Intn(9)
		}

		sudoku.Cells[x][y].SetValue(0)
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
	var lastState *model.Sudoku
	solver := strategy.Solver{}
	sudoku := model.EmptySudoku()
	_, err := solver.Solve(sudoku)
	if err != nil {
		return nil, err
	}

	localDifficulty := solver.GetLastPassDifficulty()

	for math.Abs(difficulty-localDifficulty) > 0.05 {
		x := rand.Intn(9)
		y := rand.Intn(9)
		for sudoku.Cells[x][y].Value() == 0 {
			x = rand.Intn(9)
			y = rand.Intn(9)
		}

		sudoku.Cells[x][y].SetValue(0)
		sudokuCopy, _ := model.LoadSudoku(sudoku.SaveSudoku())

		success, err := solver.Solve(sudokuCopy)
		if !success && err != nil {
			sudoku = lastState
		}

		localDifficulty = solver.GetLastPassDifficulty()
	}

	return sudoku, nil
}
