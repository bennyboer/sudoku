package pattern

import (
	"github.com/ob-algdatii-ss19/leistungsnachweis-sudo/model"
	"github.com/ob-algdatii-ss19/leistungsnachweis-sudo/solver/strategy/util"
	"testing"
)

func TestHiddenPair_Apply(t *testing.T) {
	sudoku, _ := model.LoadSudoku(&[][]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{4, 0, 0, 0, 7, 0, 6, 0, 0},
		{0, 7, 0, 2, 0, 4, 3, 0, 0},
		{0, 0, 0, 9, 0, 0, 0, 8, 4},
		{0, 6, 0, 0, 0, 0, 5, 9, 2},
		{5, 9, 0, 0, 0, 0, 0, 0, 0},
		{0, 4, 0, 8, 2, 0, 9, 0, 1},
		{0, 0, 0, 1, 0, 0, 0, 0, 0},
		{6, 5, 1, 7, 0, 0, 0, 0, 0},
	})

	possibleValueLookupRef := util.PreparePossibleValueLookup(sudoku)
	pattern := HiddenPair{}

	changed := pattern.Apply(sudoku, possibleValueLookupRef)

	if !changed {
		t.Errorf("Expected pattern to change something")
	}

	possibleValues := *(*possibleValueLookupRef)[3][4]
	for value, possible := range possibleValues {
		if (value == 1 || value == 3) && possible {
			t.Errorf("Expected cell at row 4 and column 5 to not have 1 or 3 as possible values")
		}
	}

	possibleValues = *(*possibleValueLookupRef)[3][5]
	for value, possible := range possibleValues {
		if (value == 1 || value == 2 || value == 3 || value == 7) && possible {
			t.Errorf("Expected cell at row 4 and column 6 to not have 1, 2, 3 or 7 as possible values")
		}
	}
}
