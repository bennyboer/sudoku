package generator

import (
	"github.com/ob-algdatii-ss19/leistungsnachweis-sudo/model"
	"github.com/ob-algdatii-ss19/leistungsnachweis-sudo/solver/backtracking"
	"github.com/ob-algdatii-ss19/leistungsnachweis-sudo/solver/backtracking/strategy"
	"math/rand"
)

func (sg *SudokuGeneratorBacktracking) Generate(difficulty float64) *model.Sudoku {
	var lastState *model.Sudoku
	solver := backtracking.Solver{CellChooserType: strategy.Linear}
	sudoku := model.EmptySudoku()
	deletioncount := int(((81 - 17) * difficulty) / 100)

	for i := 0; i < deletioncount; i++ {
		lastState = sudoku
		sudoku.Cells[rand.Intn(9)][rand.Intn(9)].SetValue(0)

		success, err := solver.Solve(sudoku)
		if !success && err != nil {
			i--
			sudoku = lastState
		}
	}

	return sudoku
}
