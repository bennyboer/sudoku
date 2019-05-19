package generator

import (
	"github.com/ob-algdatii-ss19/leistungsnachweis-sudo/model"
	"github.com/ob-algdatii-ss19/leistungsnachweis-sudo/solver/backtracking"
	"github.com/ob-algdatii-ss19/leistungsnachweis-sudo/solver/backtracking/strategy"
	"math/rand"
	"time"
)

func (sg *SudokuGeneratorBackTracking) Generate(difficulty float64) *model.Sudoku {
	var lastState *model.Sudoku
	solver := backtracking.Solver{CellChooserType: strategy.Linear}
	sudoku := model.EmptySudoku()
	rand.Seed(time.Now().UnixNano())
	deletioncount := int(((81 - 15) * difficulty) / 100)

	for i := 0; i < deletioncount; i++ {
		lastState = sudoku
		sudoku.Cells[rand.Intn(9)][rand.Intn(9)].Value() = 0

		success, err := solver.Solve(sudoku)
		if !success && err != nil {
			i--
			sudoku = lastState
		}
	}

	return sudoku
}
