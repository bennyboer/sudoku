package strategy

import (
	"github.com/ob-algdatii-ss19/leistungsnachweis-sudo/model"
)

// Linear cell choosing strategy.
// It will choose the next empty cell row by row, cell by cell.
type LinearCellChooser struct{}

// Get empty cells for the passed Sudoku.
func (c *LinearCellChooser) Get(sudoku *model.Sudoku) *[]*model.SudokuCell {
	emptyCells := make([]*model.SudokuCell, 0, model.SudokuSize^2)

	for row := 0; row < model.SudokuSize; row++ {
		for column := 0; column < model.SudokuSize; column++ {
			cell := sudoku.Cells[row][column]

			if cell.IsEmpty() {
				emptyCells = append(emptyCells, cell)
			}
		}
	}

	return &emptyCells
}
