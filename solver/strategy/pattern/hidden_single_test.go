package pattern

import (
	"github.com/ob-algdatii-ss19/leistungsnachweis-sudo/model"
	"github.com/ob-algdatii-ss19/leistungsnachweis-sudo/solver/strategy/util"
	"testing"
)

func TestHiddenSingle_Apply_WithChanges(t *testing.T) {
	sudoku, _ := model.LoadSudoku(&[9][9]int{
		{2, 6, 0, 0, 0, 0, 0, 0, 0},
		{1, 0, 0, 8, 0, 0, 0, 0, 0},
		{0, 0, 0, 9, 6, 2, 0, 0, 0},
		{0, 0, 0, 3, 0, 0, 7, 0, 0},
		{0, 0, 0, 0, 5, 0, 0, 0, 8},
		{0, 0, 2, 7, 4, 0, 0, 9, 0},
		{4, 0, 0, 0, 0, 0, 0, 2, 0},
		{7, 0, 3, 0, 0, 6, 0, 0, 0},
		{0, 0, 0, 0, 3, 0, 0, 8, 1},
	})

	possibleValueLookupRef := util.PreparePossibleValueLookup(sudoku)
	pattern := HiddenSingle{}

	changed := pattern.Apply(sudoku, possibleValueLookupRef)

	if !changed {
		t.Errorf("Expected pattern to change something")
	}

	if sudoku.Cells[6][8].Value() != 7 {
		t.Errorf("Expected pattern to apply value 7 to the cell at row 7 and column 9")
	}
}

func TestHiddenSingle_Apply_WithoutChanges(t *testing.T) {
	sudoku, _ := model.LoadSudoku(&[9][9]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 7, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 2, 7},
		{7, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 8, 1},
	})

	possibleValueLookupRef := util.PreparePossibleValueLookup(sudoku)
	pattern := HiddenSingle{}

	changed := pattern.Apply(sudoku, possibleValueLookupRef)

	if changed {
		t.Errorf("Expected pattern to change nothing")
	}
}
