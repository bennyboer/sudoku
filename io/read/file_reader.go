package read

import (
	"fmt"
	"github.com/ob-algdatii-ss19/leistungsnachweis-sudo/model"
	"io/ioutil"
	"strconv"
	"strings"
)

// Reader reading Sudokus from file.
type SudokuFileReader struct {
	// Path to file to read from
	FilePath *string
}

func (r *SudokuFileReader) Read() (*model.Sudoku, error) {
	bytes, err := ioutil.ReadFile(*r.FilePath)
	if err != nil {
		return nil, err
	}

	text := string(bytes)

	return loadSudoku(&text)
}

// Load the Sudoku from the passed serialized Sudoku string.
func loadSudoku(encoded *string) (*model.Sudoku, error) {
	text := *encoded

	lines := strings.Split(text, "\n")
	if len(lines) < model.SudokuSize+2 {
		return nil, fmt.Errorf("expected at least %d lines; got %d", model.SudokuSize+2, len(lines))
	}

	// Prepare source slice for Sudoku
	src := make([][]int, model.SudokuSize)
	for i := 0; i < model.SudokuSize; i++ {
		src[i] = make([]int, model.SudokuSize)
	}

	// Parse lines
	row := 0
	for _, line := range lines {
		line = strings.Trim(line, "\n\r")
		if len(line) == 0 {
			continue
		}

		cells := strings.Split(line, " ")
		column := 0

		if len(cells) < model.SudokuSize+2 {
			return nil, fmt.Errorf("expected rows in the format: 1 2 3  4 5 6  7 8 9; got '%s'", line)
		}

		for _, cell := range cells {
			if len(cell) > 0 {
				// Try to parse the cell value
				cell = strings.Trim(cell, "\n\r")
				value, err := strconv.ParseInt(cell, 10, 64)
				if err != nil {
					return nil, fmt.Errorf("expected an integer number as cell value; got %s", cell)
				}

				src[row][column] = int(value)

				column++
			}
		}

		row++
		if row >= model.SudokuSize {
			break
		}
	}

	return model.LoadSudoku(&src)
}
