package generator

import (
	"errors"
	"github.com/ob-algdatii-ss19/leistungsnachweis-sudo/model"
	"github.com/ob-algdatii-ss19/leistungsnachweis-sudo/solver/backtracking"
	s "github.com/ob-algdatii-ss19/leistungsnachweis-sudo/solver/backtracking/strategy"
	"github.com/ob-algdatii-ss19/leistungsnachweis-sudo/solver/strategy"
	"math"
	"math/rand"
	"sync"
	"time"
)

func (generator *SudokuGeneratorSimple) simpleGenerate(difficulty float64, group *sync.WaitGroup) {
	var lastState *model.Sudoku
	solver := backtracking.Solver{CellChooserType: s.Linear}
	sudoku := model.EmptySudoku()
	_, err := solver.Solve(sudoku)
	if err != nil {
		print("Error while solving")
	}

	deletioncount := int((81 - 17) * difficulty)

	for i := 0; i < deletioncount && !generator.isCancelled; i++ {
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

	group.Done()
	generator.sudoku = sudoku
}

func (sg *SudokuGeneratorSimple) Generate(difficulty float64) (*model.Sudoku, error) {
	var waitgroup = &sync.WaitGroup{}

	if difficulty > 1.0 || difficulty < 0 {
		return nil, errors.New("the difficulty must be between 0 and 1")
	}

	waitgroup.Add(1)
	go sg.simpleGenerate(difficulty, waitgroup)
	go func() {
		time.Sleep(2 * time.Second)
		sg.isCancelled = true
	}()

	waitgroup.Wait()
	return sg.sudoku, nil
}

func (generator *SudokuGeneratorDifficulty) Generate(difficulty float64) (*model.Sudoku, error) {
	solver := strategy.Solver{}
	sudoku := model.EmptySudoku()
	_, err := solver.Solve(sudoku)
	waitgroup := &sync.WaitGroup{}
	if err != nil {
		return nil, err
	}

	waitgroup.Add(1)
	go generator.backtrack(sudoku, difficulty, waitgroup)
	go func() {
		time.Sleep(2 * time.Second)
		generator.isCancelled = true
	}()
	// Waiting for results
	waitgroup.Wait()

	return generator.sudoku, nil
}

func (generator *SudokuGeneratorDifficulty) backtrack(sudoku *model.Sudoku, difficulty float64, group *sync.WaitGroup) {
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
		group.Done()
		return
	}

	group.Add(1)
	go generator.backtrack(sudoku, difficulty, group)
	group.Done()
}
