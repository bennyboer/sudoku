package write

import "github.com/ob-algdatii-ss19/leistungsnachweis-sudo/model"

// Writer writing Sudokus somewhere.
type SudokuWriter interface {
	// Write Sudoku
	Write(sudoku *model.Sudoku) error
}
