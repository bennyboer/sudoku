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
	Cells [][]*SudokuCell
}

// Get a new empty Sudoku.
func EmptySudoku() *Sudoku {
	return &Sudoku{Cells: *createCells(nil)}
}

// Load a Sudoku with the passed values.
func LoadSudoku(values *[9][9]int) (*Sudoku, error) {
	// Validate input first
	if values == nil {
		return nil, fmt.Errorf("cannot load Sudoku from no nil pointer")
	}

	for _, row := range *values {
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
func createCells(values *[9][9]int) *[][]*SudokuCell {
	cells := make([][]*SudokuCell, SudokuSize)

	for row := 0; row < SudokuSize; row++ {
		cells[row] = make([]*SudokuCell, SudokuSize)

		for column := 0; column < SudokuSize; column++ {
			value := 0

			if values != nil {
				value = values[row][column]
			}

			newCell, _ := NewSudokuCell(row, column, value)

			cells[row][column] = newCell
		}
	}

	// Initialize cell lookups
	for row := 0; row < SudokuSize; row++ {
		for column := 0; column < SudokuSize; column++ {
			cells[row][column].Init(&cells)
		}
	}

	return &cells
}

// Check if the Sudoku is valid.
func (s *Sudoku) IsValid() bool {
	for row := 0; row < SudokuSize; row++ {
		for column := 0; column < SudokuSize; column++ {
			cell := s.Cells[row][column]

			if cell.HasCollision() {
				return false
			}
		}
	}

	return true
}

// Check whether the Sudoku is completely filled.
// This does NOT mean that it is valid!
func (s *Sudoku) IsComplete() bool {
	for row := 0; row < SudokuSize; row++ {
		for column := 0; column < SudokuSize; column++ {
			cell := s.Cells[row][column]

			if cell.IsEmpty() {
				return false
			}
		}
	}

	return true;
}

// Check whether the Sudoku is completely filled AND valid.
func (s *Sudoku) IsCompleteAndValid() bool {
	for row := 0; row < SudokuSize; row++ {
		for column := 0; column < SudokuSize; column++ {
			cell := s.Cells[row][column]

			if cell.IsEmpty() || cell.HasCollision() {
				return false
			}
		}
	}

	return true;
}

// Get a String representation of the Sudoku.
func (s *Sudoku) String() string {
	var sb strings.Builder

	for rowIndex, rowCells := range s.Cells {
		for columnIndex, cell := range rowCells {
			if cell.value == 0 {
				sb.WriteRune('_')
			} else {
				sb.WriteString(fmt.Sprintf("%d", cell.value))
			}

			if (columnIndex+1)%BlockSize == 0 {
				// Is block end
				if columnIndex+1 < SudokuSize {
					// Is not last block
					sb.WriteString("   ")
				}
			} else {
				sb.WriteRune(' ')
			}
		}

		if rowIndex+1 < SudokuSize {
			// Is not last row
			sb.WriteRune('\n')

			if (rowIndex+1)%BlockSize == 0 && rowIndex+1 < SudokuSize {
				// Is block end
				sb.WriteRune('\n')
			}
		}
	}

	return sb.String()
}
