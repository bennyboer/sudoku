package pattern

import (
	"github.com/ob-algdatii-ss19/leistungsnachweis-sudo/model"
	"github.com/ob-algdatii-ss19/leistungsnachweis-sudo/solver/strategy/util"
	"testing"
)

func TestRowBlockCheck_Apply(t *testing.T) {
	sudoku, _ := model.LoadSudoku(&[][]int{
		{1, 2, 0, 0, 5, 6, 0, 8, 0},
		{0, 0, 5, 9, 0, 1, 0, 0, 6},
		{0, 0, 6, 0, 0, 2, 1, 0, 5},
		{0, 1, 2, 0, 0, 0, 4, 0, 7},
		{0, 3, 0, 1, 0, 0, 0, 0, 0},
		{7, 6, 9, 0, 2, 0, 0, 1, 3},
		{0, 0, 7, 0, 1, 8, 0, 9, 0},
		{0, 0, 0, 2, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 4, 3, 0, 7, 0},
	})

	possibleValueLookupRef := util.PreparePossibleValueLookup(sudoku)
	pattern := RowBlockCheck{}

	changed := pattern.Apply(sudoku, possibleValueLookupRef)

	if !changed {
		t.Errorf("Expected pattern to change a thing")
	}

	pv := *possibleValueLookupRef

	possibleValues := *pv[4][4]
	for value, possible := range possibleValues {
		if possible && value == 9 {
			t.Errorf("Expected cell at row 5 and column 5 to not have the following values possible: 9")
		}
	}

	possibleValues = *pv[4][5]
	for value, possible := range possibleValues {
		if possible && value == 9 {
			t.Errorf("Expected cell at row 5 and column 6 to not have the following values possible: 9")
		}
	}

	possibleValues = *pv[3][4]
	for value, possible := range possibleValues {
		if !possible && value == 9 {
			t.Errorf("Expected cell at row 3 and column 5 to have the following values possible: 9")
		}
	}
}
