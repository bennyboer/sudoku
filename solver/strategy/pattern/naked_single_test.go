package pattern

import (
	"github.com/ob-algdatii-ss19/leistungsnachweis-sudo/model"
	"github.com/ob-algdatii-ss19/leistungsnachweis-sudo/solver/strategy/util"
	"testing"
)

func TestNakedSingle_Apply_ShouldChange(t *testing.T) {
	sudoku, _ := model.LoadSudoku(&[9][9]int{
		{0, 0, 0, 0, 1, 0, 0, 0, 0},
		{0, 0, 0, 0, 2, 0, 0, 0, 0},
		{0, 0, 0, 0, 3, 0, 0, 0, 0},
		{0, 0, 0, 0, 4, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 6, 0, 0, 0, 0},
		{0, 0, 0, 0, 7, 0, 0, 0, 0},
		{0, 0, 0, 0, 8, 0, 0, 0, 0},
		{0, 0, 0, 0, 9, 0, 0, 0, 0},
	})

	possibleValueLookupRef := util.PreparePossibleValueLookup(sudoku)
	pattern := NakedSingle{}

	changed := pattern.Apply(sudoku, possibleValueLookupRef)

	if !changed {
		t.Errorf("Expected pattern to change a thing")
	}

	if sudoku.Cells[4][4].Value() != 5 {
		t.Errorf("Expected pattern to apply 5 to the middle Sudoku cell")
	}
}

func TestNakedSingle_Apply_ShouldNotChange(t *testing.T) {
	sudoku, _ := model.LoadSudoku(&[9][9]int{
		{0, 0, 0, 0, 1, 0, 0, 0, 0},
		{0, 0, 0, 0, 2, 0, 0, 0, 0},
		{0, 0, 0, 0, 3, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 6, 0, 0, 0, 0},
		{0, 0, 0, 0, 7, 0, 0, 0, 0},
		{0, 0, 0, 0, 8, 0, 0, 0, 0},
		{0, 0, 0, 0, 9, 0, 0, 0, 0},
	})

	possibleValueLookupRef := util.PreparePossibleValueLookup(sudoku)
	pattern := NakedSingle{}

	changed := pattern.Apply(sudoku, possibleValueLookupRef)

	if changed {
		t.Errorf("Expected pattern to change nothing")
	}

	if sudoku.Cells[4][4].Value() != 0 {
		t.Errorf("Expected pattern to apply nothing to the middle Sudoku cell")
	}
}
