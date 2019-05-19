package read

import "github.com/ob-algdatii-ss19/leistungsnachweis-sudo/model"

// Reader for sudokus.
type SudokuReader interface {
	// Read sudoku
	Read() (*model.Sudoku, error)
}
