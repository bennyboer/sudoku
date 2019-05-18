package write

import (
	"fmt"
	"github.com/ob-algdatii-ss19/leistungsnachweis-sudo/model"
	"io/ioutil"
	"strings"
)

// Writer writing Sudokus to file.
type SudokuFileWriter struct {
	// Where to write Sudokus to.
	FilePath *string
}

func (w *SudokuFileWriter) Write(sudoku *model.Sudoku) (error) {
	encoded := encodeSudoku(sudoku)

	err := ioutil.WriteFile(*w.FilePath, []byte(encoded), 0644)
	if err != nil {
		return err
	}

	return nil
}

// Encode the passed Sudoku to string.
func encodeSudoku(sudoku *model.Sudoku) string {
	var sb strings.Builder

	for rowIndex, rowCells := range sudoku.Cells {
		for columnIndex, cell := range rowCells {
			sb.WriteString(fmt.Sprintf("%d", cell.Value()))

			if (columnIndex+1)%model.BlockSize == 0 {
				// Is block end
				if columnIndex+1 < model.SudokuSize {
					// Is not last block
					sb.WriteString("  ")
				}
			} else {
				sb.WriteRune(' ')
			}
		}

		if rowIndex+1 < model.SudokuSize {
			// Is not last row
			sb.WriteRune('\n')

			if (rowIndex+1)%model.BlockSize == 0 && rowIndex+1 < model.SudokuSize {
				// Is block end
				sb.WriteRune('\n')
			}
		}
	}

	return sb.String()
}
