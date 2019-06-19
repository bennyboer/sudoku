package generator

import (
	"errors"
	"fmt"
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
	if difficulty > 1.0 || difficulty < 0 {
		return nil, errors.New("the difficulty must be between 0 and 1")
	}

	var waitgroup sync.WaitGroup
	waitgroup.Add(1)

	go generator.start(difficulty, 0.05, &waitgroup)

	go func() {
		time.Sleep(timeout)
		generator.isCancelled = true
		fmt.Println("Warning: Generator ran into timeout. Cancelling...")
	}()

	// Waiting for results
	waitgroup.Wait()

	// Load best sudoku which the generator could come up with
	bestSudoku, _ := model.LoadSudoku(generator.sudokuSrc)
	return bestSudoku, nil
}

// Try to generate a Sudoku using backtracking and measuring the difficulty of a Sudoku.
// Will return whether the generator could find a Sudoku with the given difficulty.
func (generator *SudokuGeneratorDifficulty) start(difficulty float64, maxDeviation float64, group *sync.WaitGroup) {
	foundSudoku := false
	for !foundSudoku {
		// Generate random Sudoku as starting point
		solver := strategy.Solver{}
		sudoku := model.EmptySudoku()
		_, err := solver.Solve(sudoku)
		if err != nil {
			fmt.Printf("Error: Could not solve Sudoku.. Error:\n%s", err.Error())
			break
		}

		sudokuSrc := sudoku.SaveSudoku()
		generator.sudokuSrc = sudokuSrc

		foundSudoku = generator.find(sudokuSrc, difficulty, maxDeviation, 1.0, group)

		// Check if run into timeout
		if generator.isCancelled {
			break
		}
	}

	group.Done()
}

func (generator *SudokuGeneratorDifficulty) find(sudokuSrcPtr *[][]int, difficulty float64, maxDeviation float64, bestDeviation float64, group *sync.WaitGroup) bool {
	sudokuSrc := *sudokuSrcPtr

	for {
		if generator.isCancelled {
			return false
		}

		// Choose random cell coordinates in the Sudoku
		x := rand.Intn(9)
		y := rand.Intn(9)
		for sudokuSrc[x][y] == 0 {
			x = rand.Intn(9)
			y = rand.Intn(9)
		}

		// Copy old sudoku src and empty cell of randomly chosen cell
		newSudokuSrc := make([][]int, model.SudokuSize)
		for rowIndex, rowValues := range sudokuSrc {
			newSudokuSrc[rowIndex] = make([]int, model.SudokuSize)
			copy(newSudokuSrc[rowIndex], rowValues)
		}
		newSudokuSrc[x][y] = 0 // Set cell value to "empty"

		// Load the altered Sudoku and try to solve it
		solveSudoku, _ := model.LoadSudoku(&newSudokuSrc)
		solver := strategy.Solver{}
		success, err := solver.Solve(solveSudoku)
		if err != nil || !success {
			return false
		}
		localDifficulty := solver.GetLastPassDifficulty()

		// Check if difficulty meets the requirements
		difficultyDiff := math.Abs(difficulty - localDifficulty)
		if difficultyDiff > maxDeviation {
			// Sudoku is not good enough

			// Check if Sudoku is the best one yet
			if difficultyDiff < bestDeviation {
				bestDeviation = difficultyDiff
				generator.sudokuSrc = &newSudokuSrc // Save Sudoku since it is the best one yet
			}

			// Check if difficulty is already too high.
			// Would not make sense to continue if it is already way to high.
			if localDifficulty > difficulty {
				return false
			}

			// Check if timeout has been reached
			if generator.isCancelled {
				return false
			}

			// Try to find a better Sudoku by emptying another cell
			foundBetterSudoku := generator.find(&newSudokuSrc, difficulty, maxDeviation, bestDeviation, group)
			if foundBetterSudoku {
				return true
			}
		} else {
			// Found Sudoku that's good enough! -> Save Sudoku and return
			generator.sudokuSrc = &newSudokuSrc
			return true
		}
	}
}
