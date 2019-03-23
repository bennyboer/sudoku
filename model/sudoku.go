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
	// Count of neighbours each cell has.
	NeighbourCount int = 20
)

// A model for a Sudoku.
type Sudoku struct {
	// All cells of a sudoku.
	Cells [][]SudokuCell
}

// Get a new empty Sudoku.
func EmptySudoku() *Sudoku {
	return &Sudoku{Cells: *createCells(nil)}
}

// Load a Sudoku with the passed values.
func LoadSudoku(values *[][]int) (*Sudoku, error) {
	// Validate input first
	if values == nil {
		return nil, fmt.Errorf("cannot load Sudoku from no nil pointer")
	}

	if len(*values) != SudokuSize {
		return nil, fmt.Errorf("cannot load Sudoku from slice with less or more than 9 rows")
	}

	for _, row := range *values {
		if len(row) != SudokuSize {
			return nil, fmt.Errorf("cannot load Sudoku from slice with less or more than 9 columns")
		}

		for _, value := range row {
			if value < 0 || value > SudokuSize {
				return nil, fmt.Errorf("the values to load need to be in range [0; 9]")
			}
		}
	}

	return &Sudoku{Cells: *createCells(values)}, nil
}

// Save a Sudoku.
func (s *Sudoku) SaveSudoku() *[][]int {
	values := make([][]int, SudokuSize)

	for rowIndex, row := range s.Cells {
		values[rowIndex] = make([]int, SudokuSize)

		for columnIndex, cell := range row {
			values[rowIndex][columnIndex] = cell.Value()
		}
	}

	return &values
}

// Generate Sudoku cells filled with the passed values or if nil is given empty cells.
func createCells(values *[][]int) *[][]SudokuCell {
	cells := make([][]SudokuCell, SudokuSize)

	for row := range cells {
		cells[row] = make([]SudokuCell, SudokuSize)

		for column := range cells[row] {
			value := 0

			if values != nil {
				value = (*values)[row][column]
			}

			newCell, _ := NewSudokuCell(row, column, value)

			cells[row][column] = *newCell
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
			if cell.value == 0 {
				sb.WriteString("_ ")
			} else {
				sb.WriteString(fmt.Sprintf("%d ", cell.value))
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
