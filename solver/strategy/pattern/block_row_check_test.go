package pattern

import (
	"github.com/ob-algdatii-ss19/leistungsnachweis-sudo/model"
	"github.com/ob-algdatii-ss19/leistungsnachweis-sudo/solver/strategy/util"
	"testing"
)

func TestBlockRowCheck_Apply(t *testing.T) {
	sudoku, _ := model.LoadSudoku(&[][]int{
		{0, 0, 0, 9, 2, 3, 1, 0, 0},
		{5, 0, 7, 0, 0, 0, 0, 8, 0},
		{2, 0, 1, 0, 8, 0, 0, 0, 0},
		{0, 4, 0, 0, 0, 0, 6, 0, 0},
		{0, 0, 0, 0, 5, 0, 0, 2, 0},
		{0, 1, 0, 0, 0, 0, 0, 0, 0},
		{8, 0, 0, 0, 7, 0, 0, 0, 0},
		{0, 0, 0, 6, 0, 0, 4, 0, 0},
		{0, 0, 0, 3, 0, 0, 0, 0, 9},
	})

	possibleValueLookupRef := util.PreparePossibleValueLookup(sudoku)
	pattern := BlockRowCheck{}

	changed := pattern.Apply(sudoku, possibleValueLookupRef)

	if !changed {
		t.Errorf("Expected pattern to change something")
	}

	pv := *possibleValueLookupRef

	possibleValues := *pv[0][7]
	for value, possible := range possibleValues {
		if possible && value == 4 {
			t.Errorf("Expected cell at row 1 and column 8 to have the following value impossible: 4")
		}
	}
}
