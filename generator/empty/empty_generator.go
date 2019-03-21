package empty

import (
	"github.com/ob-algdatii-ss19/leistungsnachweis-sudo/model"
	"github.com/ob-algdatii-ss19/leistungsnachweis-sudo/model/cell"
)

// Size of a Sudoku (amount of cells in rows, columns and blocks).
const sudoku_size int = 9;

// Generator generating empty Sudokus.
type EmptyGenerator struct{}

// Will create an empty Sudoku from scratch.
func (g *EmptyGenerator) Generate(difficulty float32) *model.Sudoku {
	cells := make([][]cell.SudokuCell, sudoku_size, sudoku_size)

	for row := range cells {
		cells[row] = make([]cell.SudokuCell, sudoku_size, sudoku_size)

		for column := range cells[row] {
			cells[row][column] = cell.SudokuCell{
				Position: cell.Coordinates{
					Row:    row,
					Column: column,
				},
			}
		}
	}

	return &model.Sudoku{Cells: cells}
}
