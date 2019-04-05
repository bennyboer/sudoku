package pattern

import (
	"github.com/ob-algdatii-ss19/leistungsnachweis-sudo/model"
	"github.com/ob-algdatii-ss19/leistungsnachweis-sudo/solver/strategy/util"
	"testing"
)

func TestNakedPair_Apply_WithChanges(t *testing.T) {
	sudoku, _ := model.LoadSudoku(&[9][9]int{
		{0, 0, 0, 1, 3, 7, 0, 0, 0},
		{7, 0, 0, 5, 9, 6, 1, 3, 0},
		{0, 0, 9, 0, 8, 0, 0, 6, 0},
		{0, 0, 3, 0, 2, 0, 0, 0, 0},
		{5, 0, 0, 8, 0, 0, 9, 2, 0},
		{0, 2, 0, 0, 1, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 8},
		{8, 7, 4, 0, 0, 0, 0, 0, 0},
		{0, 6, 5, 3, 4, 8, 0, 0, 0},
	})

	possibleValueLookupRef := util.PreparePossibleValueLookup(sudoku)
	pattern := NakedPair{}

	changed := pattern.Apply(sudoku, possibleValueLookupRef)

	if !changed {
		t.Errorf("Expected pattern to change a thing")
	}

	pv := *possibleValueLookupRef
	lookup := *pv[2][0]
	for value, possible := range lookup {
		if (value == 2 || value == 4) && possible {
			t.Errorf("Possible values of the cell should not contain 2 or 4 " +
				"as possible values after the pattern has been applied")
		}
	}

	lookup = *pv[2][3]
	possible, _ := lookup[2]
	if !possible {
		t.Errorf("The pair value should not have changed")
	}

	possible, _ = lookup[4]
	if !possible {
		t.Errorf("The pair value should not have changed")
	}
}
