package model

import (
	"fmt"
	sodukuCell "github.com/ob-algdatii-ss19/leistungsnachweis-sudo/model/cell"
	"strings"
)

// A model for a Sudoku.
type Sudoku struct {
	// All cells of a sudoku.
	Cells [][]sodukuCell.SudokuCell
}

// Get a String representation of the Sudoku.
func (s *Sudoku) String() string {
	var sb strings.Builder

	for _, rowCells := range s.Cells {
		for _, cell := range rowCells {
			if cell.Value == 0 {
				sb.WriteString("_ ")
			} else {
				sb.WriteString(fmt.Sprintf("%s ", cell.Value))
			}
		}

		sb.WriteRune('\n')
	}

	return sb.String()
}
