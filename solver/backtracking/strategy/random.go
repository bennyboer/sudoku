package strategy

import (
	"github.com/ob-algdatii-ss19/leistungsnachweis-sudo/model"
	"math/rand"
	"time"
)

// Random cell choosing strategy.
// It will choose the next empty cell randomly.
type RandomCellChooser struct{}

// Get empty cells for the passed sudoku in random order.
func (c *RandomCellChooser) Get(sudoku *model.Sudoku) *[]*model.SudokuCell {
	emptyCells := make([]*model.SudokuCell, 0, model.SudokuSize^2)

	for row := 0; row < model.SudokuSize; row++ {
		for column := 0; column < model.SudokuSize; column++ {
			cell := sudoku.Cells[row][column]

			if cell.IsEmpty() {
				emptyCells = append(emptyCells, cell)
			}
		}
	}

	// Shuffle empty cell slice
	shuffled := make([]*model.SudokuCell, len(emptyCells))

	rand.Seed(time.Now().UnixNano())
	permutation := rand.Perm(len(emptyCells))
	for oldIndex, newIndex := range permutation {
		shuffled[newIndex] = emptyCells[oldIndex]
	}

	return &shuffled
}
