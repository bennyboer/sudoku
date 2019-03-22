package model

import (
	"fmt"
	"strings"
)

const (
	// Number of cells in rows, columns and blocks of a Sudoku.
	SudokuSize int = 9
	// Number of cells in rows, columns of a block.
	BlockSize int = 3
)

// A model for a Sudoku.
type Sudoku struct {
	// All cells of a sudoku.
	Cells [][]SudokuCell
}

// Get a new empty Sudoku.
func EmptySudoku() *Sudoku {
	return &Sudoku{Cells: *createCells()}
}

// Generate empty Sudoku cells.
func createCells() *[][]SudokuCell {
	cells := make([][]SudokuCell, SudokuSize, SudokuSize)

	for row := range cells {
		cells[row] = make([]SudokuCell, SudokuSize, SudokuSize)

		for column := range cells[row] {
			cells[row][column] = *NewSudokuCell(row, column, 0)
		}
	}

	// Initialize cell lookups
	for _, rowCells := range cells {
		for _, cell := range rowCells {
			cell.Init(&cells)
		}
	}

	return &cells
}

// Get a String representation of the Sudoku.
func (s *Sudoku) String() string {
	var sb strings.Builder

	for rowIndex, rowCells := range s.Cells {
		for columnIndex, cell := range rowCells {
			if cell.Value == 0 {
				sb.WriteString("_ ")
			} else {
				sb.WriteString(fmt.Sprintf("%d ", cell.Value))
			}

			if (columnIndex+1)%BlockSize == 0 {
				// Is block end
				sb.WriteString("  ")
			}
		}

		sb.WriteRune('\n')

		if (rowIndex+1)%BlockSize == 0 {
			// Is block end
			sb.WriteRune('\n')
		}
	}

	return sb.String()
}
