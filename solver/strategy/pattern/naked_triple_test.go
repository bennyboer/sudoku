package pattern

import (
	"github.com/ob-algdatii-ss19/leistungsnachweis-sudo/model"
	"github.com/ob-algdatii-ss19/leistungsnachweis-sudo/solver/strategy/util"
	"testing"
)

func TestNakedTriple_Apply(t *testing.T) {
	sudoku, _ := model.LoadSudoku(&[9][9]int{
		{0, 7, 0, 0, 0, 0, 8, 0, 0},
		{0, 2, 0, 8, 0, 0, 9, 5, 0},
		{0, 0, 0, 0, 0, 9, 6, 0, 2},
		{0, 0, 0, 3, 0, 4, 2, 8, 9},
		{0, 0, 3, 9, 0, 0, 1, 6, 4},
		{4, 0, 0, 6, 1, 0, 5, 0, 0},
		{0, 4, 0, 2, 9, 6, 0, 1, 5},
		{0, 0, 0, 0, 0, 0, 0, 9, 6},
		{0, 0, 0, 5, 3, 0, 4, 2, 8},
	})

	possibleValueLookupRef := util.PreparePossibleValueLookup(sudoku)
	pattern := NakedTriple{}

	changed := pattern.Apply(sudoku, possibleValueLookupRef)

	if !changed {
		t.Errorf("Expected pattern to change a thing")
	}

	pv := *possibleValueLookupRef

	possibleValues := *pv[0][0]
	for value, possible := range possibleValues {
		if possible && (value == 1 || value == 3) {
			t.Errorf("Expected possible values of cell at row 1 and column 1 to not include 1 or 3")
		}
	}

	possibleValues = *pv[0][2]
	for value, possible := range possibleValues {
		if possible && (value == 1 || value == 4) {
			t.Errorf("Expected possible values of cell at row 1 and column 3 to not include 1 or 4")
		}
	}

	possibleValues = *pv[0][3]
	for value, possible := range possibleValues {
		if possible && value != 1 && value != 4 {
			t.Errorf("Expected possible values of cell at row 1 and column 3 to only include 1 and 4")
		}
	}
}
