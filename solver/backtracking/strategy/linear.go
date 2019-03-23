package strategy

import "github.com/ob-algdatii-ss19/leistungsnachweis-sudo/model"

// Linear cell choosing strategy.
// It will choose the next empty cell row by row, cell by cell.
type LinearCellChooser struct{}

// Initialize the strategy.
func (LinearCellChooser) Initialize(sudoku *model.Sudoku) {
	panic("implement me")
}

// Find the next empty cell.
func (LinearCellChooser) FindNext() (*model.SudokuCell, error) {
	panic("implement me")
}
