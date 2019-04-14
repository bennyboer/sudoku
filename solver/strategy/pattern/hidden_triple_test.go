package pattern

import (
	"github.com/ob-algdatii-ss19/leistungsnachweis-sudo/model"
	"github.com/ob-algdatii-ss19/leistungsnachweis-sudo/solver/strategy/util"
	"testing"
)

func TestHiddenTriple_Apply(t *testing.T) {
	sudoku, _ := model.LoadSudoku(&[][]int{
		{0, 0, 0, 7, 4, 0, 0, 0, 8},
		{4, 9, 6, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 2, 0, 0, 4, 0},
		{0, 0, 0, 0, 5, 7, 6, 0, 0},
		{8, 0, 0, 0, 0, 0, 0, 2, 1},
		{0, 0, 3, 4, 0, 0, 0, 0, 0},
		{0, 0, 0, 3, 0, 0, 0, 0, 0},
		{1, 2, 4, 0, 7, 0, 0, 5, 0},
		{0, 0, 0, 0, 0, 0, 7, 0, 0},
	})

	possibleValueLookupRef := util.PreparePossibleValueLookup(sudoku)
	pattern := HiddenTriple{}

	changed := pattern.Apply(sudoku, possibleValueLookupRef)

	if !changed {
		t.Errorf("Expected pattern to change something")
	}

	pv := *possibleValueLookupRef

	possibleValues := *pv[4][1]
	for value, possible := range possibleValues {
		if possible && value == 6 {
			t.Errorf("Expected value 6 to be impossible for cell at row 5 and column 2")
		}
	}

	possibleValues = *pv[4][2]
	for value, possible := range possibleValues {
		if possible && value == 9 {
			t.Errorf("Expected value 9 to be impossible for cell at row 5 and column 3")
		}
	}

	possibleValues = *pv[4][6]
	for value, possible := range possibleValues {
		if possible && (value == 3 || value == 9) {
			t.Errorf("Expected value 9 or 3 to be impossible for cell at row 5 and column 7")
		}
	}

	possibleValues = *pv[4][3]
	for value, possible := range possibleValues {
		if possible && value != 6 && value != 9 {
			t.Errorf("Expected value 6 or p to be the only possible for cell at row 5 and column 4")
		}
	}
}
