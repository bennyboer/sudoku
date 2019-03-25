package main

import (
	"fmt"
	"github.com/ob-algdatii-ss19/leistungsnachweis-sudo/model"
	"github.com/ob-algdatii-ss19/leistungsnachweis-sudo/solver/backtracking"
	"github.com/ob-algdatii-ss19/leistungsnachweis-sudo/solver/backtracking/strategy"
)

func main() {
	//backtrackingSolver := backtracking.Solver{CellChooserType: strategy.Linear}
	//randomSudoku := model.EmptySudoku()
	//_, _ = backtrackingSolver.Solve(randomSudoku)
	//
	//fmt.Println(randomSudoku.String())
	//if randomSudoku.IsValid() {
	//	fmt.Printf("Sudoku above is valid\n")
	//} else {
	//	fmt.Printf("Sudoku above is invalid, which should not happen...\n")
	//}
	//
	//s := model.EmptySudoku()
	//
	//unique, _ := solver.HasUniqueSolution(s)
	//if unique {
	//	fmt.Println("Sudoku has an unique solution")
	//} else {
	//	fmt.Println("Sudoku has no unique solution")
	//}

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

	solver := backtracking.Solver{CellChooserType: strategy.Linear}

	hasUniqueSolution, _ := solver.HasUniqueSolution(*sudoku)
	fmt.Println(sudoku.String())
	if hasUniqueSolution {
		fmt.Println("Sudoku has an unique solution")
	} else {
		fmt.Println("Sudoku does not have an unique solution")
	}

	_, _ = solver.Solve(sudoku)

	fmt.Println(sudoku.String())
}
