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

func (sg *SudokuGeneratorSimple) Generate(difficulty float64) (*model.Sudoku, error) {
	var lastState *model.Sudoku
	solver := backtracking.Solver{CellChooserType: s.Linear}
	sudoku := model.EmptySudoku()
	_, err := solver.Solve(sudoku)
	if err != nil {
		return nil, err
	}

	if difficulty > 1.0 || difficulty < 0 {
		return nil, errors.New("the difficulty must be between 0 and 1")
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
	solver := strategy.Solver{}
	sudoku := model.EmptySudoku()
	_, err := solver.Solve(sudoku)
	if err != nil {
		return nil, err
	}

	generator.backtrack(sudoku, difficulty)

	return generator.sudoku, nil
}

func (generator *SudokuGeneratorDifficulty) backtrack(sudoku *model.Sudoku, difficulty float64) {
	solver := strategy.Solver{}

	x := rand.Intn(9)
	y := rand.Intn(9)
	for sudoku.Cells[x][y].Value() == 0 {
		x = rand.Intn(9)
		y = rand.Intn(9)
	}

	sudoku.Cells[x][y].SetValue(0)
	sudokuCopy, _ := model.LoadSudoku(sudoku.SaveSudoku())

	success, err := solver.Solve(sudokuCopy)
	localDifficulty := solver.GetLastPassDifficulty()
	if success && err == nil {
		if localDifficulty > generator.difficulty {
			generator.sudoku = sudoku
			generator.difficulty = localDifficulty
		}
	}

	if (!success || err != nil) ||
		math.Abs(difficulty-localDifficulty) < 0.05 ||
		difficulty < localDifficulty ||
		generator.isCancelled {

		return
	}

	generator.backtrack(sudoku, difficulty)
}
