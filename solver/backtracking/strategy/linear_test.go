package strategy

import (
	"github.com/ob-algdatii-ss19/leistungsnachweis-sudo/model"
	"testing"
)

func TestLinearCellChooser_Get(t *testing.T) {
	cellChooser := LinearCellChooser{}

	sudoku, _ := model.LoadSudoku(&[9][9]int{
		{0, 1, 2, 0, 0, 0, 5, 7, 0},
		{6, 0, 0, 5, 0, 1, 0, 0, 4},
		{4, 0, 0, 0, 2, 0, 0, 0, 8},
		{0, 2, 0, 0, 1, 0, 0, 5, 0},
		{0, 0, 4, 9, 0, 7, 8, 0, 0},
		{0, 7, 0, 0, 8, 0, 0, 1, 0},
		{7, 0, 0, 0, 9, 0, 0, 0, 5},
		{5, 0, 0, 4, 0, 8, 0, 0, 6},
		{0, 3, 8, 0, 0, 0, 9, 4, 0},
	})
	countOfEmptyCells := 49

	emptyCells := *cellChooser.Get(sudoku)
	if len(emptyCells) != 49 {
		t.Errorf("Should have the number of empty cells correctly. Expected: %d, got: %d",
			countOfEmptyCells, len(emptyCells))
	}

	nextEmptyCell := emptyCells[0]
	if nextEmptyCell.Position().Row != 0 && nextEmptyCell.Position().Column != 0 {
		t.Errorf("Expected first empty cell to have position (0, 0); got (%d, %d)",
			nextEmptyCell.Position().Row, nextEmptyCell.Position().Column)
	}

	nextEmptyCell = emptyCells[1]
	if nextEmptyCell.Position().Row != 0 && nextEmptyCell.Position().Column != 3 {
		t.Errorf("Expected first empty cell to have position (0, 3); got (%d, %d)",
			nextEmptyCell.Position().Row, nextEmptyCell.Position().Column)
	}

	nextEmptyCell = emptyCells[6]
	if nextEmptyCell.Position().Row != 1 && nextEmptyCell.Position().Column != 1 {
		t.Errorf("Expected first empty cell to have position (1, 1); got (%d, %d)",
			nextEmptyCell.Position().Row, nextEmptyCell.Position().Column)
	}

	sudoku, _ = model.LoadSudoku(&[9][9]int{
		{9, 1, 2, 8, 4, 6, 5, 7, 3},
		{6, 8, 3, 5, 7, 1, 2, 9, 4},
		{4, 5, 7, 3, 2, 9, 1, 6, 8},
		{8, 2, 9, 6, 1, 3, 4, 5, 7},
		{1, 6, 4, 9, 5, 7, 8, 3, 2},
		{3, 7, 5, 2, 8, 4, 6, 1, 9},
		{7, 4, 6, 1, 9, 2, 3, 8, 5},
		{5, 9, 1, 4, 3, 8, 7, 2, 6},
		{2, 3, 8, 7, 6, 5, 9, 4, 1},
	})

	emptyCells = *cellChooser.Get(sudoku)
	if len(emptyCells) != 0 {
		t.Errorf("Expected 0 empty cells; got %d", len(emptyCells))
	}

	sudoku = model.EmptySudoku()

	emptyCells = *cellChooser.Get(sudoku)
	if len(emptyCells) != 81 {
		t.Errorf("Expected 81 empty cells; got %d", len(emptyCells))
	}
}
