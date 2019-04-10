package pattern

import (
	"github.com/ob-algdatii-ss19/leistungsnachweis-sudo/model"
	"github.com/ob-algdatii-ss19/leistungsnachweis-sudo/solver/strategy/util"
	"testing"
)

func TestXWing_Apply(t *testing.T) {
	sudoku, _ := model.LoadSudoku(&[][]int{
		{1, 0, 0, 0, 0, 0, 0, 8, 0},
		{8, 0, 0, 1, 0, 0, 0, 2, 4},
		{7, 0, 0, 0, 0, 3, 1, 5, 0},
		{0, 0, 0, 0, 4, 1, 6, 9, 2},
		{0, 9, 0, 6, 7, 0, 4, 1, 3},
		{4, 1, 6, 2, 3, 9, 8, 7, 5},
		{9, 0, 1, 0, 6, 2, 5, 0, 8},
		{0, 0, 0, 3, 0, 0, 9, 0, 1},
		{0, 5, 0, 9, 1, 0, 2, 0, 7},
	})

	possibleValueLookupRef := util.PreparePossibleValueLookup(sudoku)
	pattern := XWing{}

	changed := pattern.Apply(sudoku, possibleValueLookupRef)

	if !changed {
		t.Errorf("Expected pattern to change a thing")
	}

	pv := *possibleValueLookupRef

	possibleValues := *pv[3][2]
	for value, possible := range possibleValues {
		if possible && value == 8 {
			t.Errorf("Expected cell at row 4 and column 3 to not have the following value possible: 8")
		}
	}

	possibleValues = *pv[7][2]
	for value, possible := range possibleValues {
		if possible && value == 8 {
			t.Errorf("Expected cell at row 8 and column 3 to not have the following value possible: 8")
		}
	}

	possibleValues = *pv[7][5]
	for value, possible := range possibleValues {
		if possible && value == 8 {
			t.Errorf("Expected cell at row 8 and column 6 to not have the following value possible: 8")
		}
	}

	changed = pattern.Apply(sudoku, possibleValueLookupRef)

	if !changed {
		t.Errorf("Expected pattern to change a thing")
	}
}
