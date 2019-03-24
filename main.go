package main

import (
	"fmt"
	"github.com/ob-algdatii-ss19/leistungsnachweis-sudo/model"
	"github.com/ob-algdatii-ss19/leistungsnachweis-sudo/solver/backtracking"
	"github.com/ob-algdatii-ss19/leistungsnachweis-sudo/solver/backtracking/strategy"
)

func main() {
	backtrackingSolver := backtracking.Solver{CellChooserType: strategy.Linear}
	randomSudoku := model.EmptySudoku()
	_, _ = backtrackingSolver.Solve(randomSudoku)

	fmt.Println(randomSudoku.String())
	if randomSudoku.IsValid() {
		fmt.Printf("Sudoku above is valid")
	} else {
		fmt.Printf("Sudoku above is invalid, which should not happen...")
	}
}
