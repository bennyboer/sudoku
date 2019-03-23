package strategy

import "github.com/ob-algdatii-ss19/leistungsnachweis-sudo/model"

// Random cell choosing strategy.
// It will choose the next empty cell randomly.
type RandomCellChooser struct{}

// Initialize the strategy.
func (c *RandomCellChooser) Initialize(sudoku *model.Sudoku) {
	panic("implement me")
}

// Find the next empty cell.
func (c *RandomCellChooser) FindNext() (*model.SudokuCell, error) {
	panic("implement me")
}
