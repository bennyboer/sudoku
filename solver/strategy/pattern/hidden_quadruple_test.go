package pattern

import (
	"github.com/ob-algdatii-ss19/leistungsnachweis-sudo/model"
	"github.com/ob-algdatii-ss19/leistungsnachweis-sudo/solver/strategy/util"
	"testing"
)

func TestHiddenQuadruple_Apply(t *testing.T) {
	sudoku, _ := model.LoadSudoku(&[][]int{
		{0, 0, 0, 7, 1, 0, 2, 5, 0},
		{0, 3, 1, 6, 0, 0, 0, 0, 8},
		{0, 5, 7, 9, 0, 0, 0, 1, 0},
		{0, 0, 0, 0, 4, 0, 0, 0, 0},
		{0, 7, 0, 0, 6, 2, 1, 0, 5},
		{0, 0, 6, 0, 9, 7, 8, 0, 2},
		{0, 0, 9, 2, 0, 1, 0, 6, 0},
		{0, 0, 0, 0, 7, 9, 3, 2, 1},
		{0, 0, 0, 0, 0, 6, 0, 8, 9},
	})

	possibleValueLookupRef := util.PreparePossibleValueLookup(sudoku)
	pattern := HiddenQuadruple{}

	changed := pattern.Apply(sudoku, possibleValueLookupRef)

	if !changed {
		t.Errorf("Expected pattern to change a thing")
	}

	pv := *possibleValueLookupRef

	possibleValues := *pv[6][0]
	for value, possible := range possibleValues {
		if possible && (value == 4 || value == 5 || value == 8) {
			t.Errorf("Expected cell at row 7 and column 1 to not have any of the following values: 4, 5, 8")
		}
	}

	possibleValues = *pv[8][0]
	for value, possible := range possibleValues {
		if possible && (value == 4 || value == 5) {
			t.Errorf("Expected cell at row 9 and column 1 to not have any of the following values: 4, 5")
		}
	}

	possibleValues = *pv[8][2]
	for value, possible := range possibleValues {
		if possible && (value == 4 || value == 5) {
			t.Errorf("Expected cell at row 9 and column 3 to not have any of the following values: 4, 5")
		}
	}

	possibleValues = *pv[7][0]
	for value, possible := range possibleValues {
		if possible && value != 4 && value != 5 && value != 6 && value != 8 {
			t.Errorf("Expected cell at row 8 and column 1 to have the following values: 4, 5, 6, 8")
		}
	}
}
