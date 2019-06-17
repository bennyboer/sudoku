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

func (generator *SudokuGeneratorSimple) Generate(difficulty float64, timeout time.Duration) (*model.Sudoku, error) {
	var waitgroup sync.WaitGroup

	if difficulty > 1.0 || difficulty < 0 {
		return nil, errors.New("the difficulty must be between 0 and 1")
	}

	waitgroup.Add(1)
	go generator.simpleGenerate(difficulty, &waitgroup)
	go func() {
		time.Sleep(timeout)
		generator.isCancelled = true
	}()

	waitgroup.Wait()
	return generator.sudoku, nil
}

func (generator *SudokuGeneratorDifficulty) Generate(difficulty float64, timeout time.Duration) (*model.Sudoku, error) {
	solver := strategy.Solver{}
	sudoku := model.EmptySudoku()
	_, err := solver.Solve(sudoku)
	var waitgroup sync.WaitGroup
	if err != nil {
		return nil, err
	}

	if difficulty > 1.0 || difficulty < 0 {
		return nil, errors.New("the difficulty must be between 0 and 1")
	}

	waitgroup.Add(1)
	go generator.backtrack(sudoku, difficulty, &waitgroup)
	go func() {
		time.Sleep(timeout)
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

	backupSudoku, _ := model.LoadSudoku(sudoku.SaveSudoku())

	sudoku.Cells[x][y].SetValue(0)
	sudokuCopy, _ := model.LoadSudoku(sudoku.SaveSudoku())

	success, err := solver.Solve(sudokuCopy)
	localDifficulty := solver.GetLastPassDifficulty()

	if (!success || err != nil) ||
		math.Abs(difficulty-localDifficulty) < 0.05 ||
		localDifficulty > difficulty+0.05 ||
		generator.isCancelled {

		if generator.isCancelled {
			group.Done()
			return
		}

		group.Add(1)
		go generator.backtrack(backupSudoku, difficulty, group)
		group.Done()
		return
	}

	if success {
		if localDifficulty > generator.difficulty && !generator.isCancelled {
			generator.lock.Lock()
			generator.sudoku = sudoku
			generator.difficulty = localDifficulty
			generator.lock.Unlock()
		}

		if math.Abs(difficulty-localDifficulty) < 0.05 {
			generator.isCancelled = true
		}
	}

	group.Add(1)
	go generator.backtrack(sudoku, difficulty, group)
	group.Done()
}
