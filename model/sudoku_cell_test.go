package model

import "testing"

func TestNewSudokuCell(t *testing.T) {
	cell, e := NewSudokuCell(4, 6, 0)

	if e != nil {
		t.Errorf("NewSudokuCell{4, 6, 0}: An error has been thrown where it should not")
	}

	if cell.Value() != 0 {
		t.Errorf("Value of cell.Value() should have been 0 while it was %d", cell.Value())
	}

	correctPos := Coordinates{Row: 4, Column: 6}
	if cell.Position() != correctPos {
		t.Errorf("cell.Position(): Position should have been %v while it was %v", correctPos, cell.Position())
	}

	_, e = NewSudokuCell(-4, 234, 4)
	if e == nil {
		t.Errorf("NewSudokuCell(-4, 234, 4): Expected to throw error because of negative value -4")
	}

	_, e = NewSudokuCell(345, -6, 4)
	if e == nil {
		t.Errorf("NewSudokuCell(345, -6, 4): Expected to throw error because of too high value 345")
	}

	_, e = NewSudokuCell(2, 3, -1)
	if e == nil {
		t.Errorf("NewSudokuCell(2, 3, -5): Expected to throw error because of invalid value")
	}

	_, e = NewSudokuCell(2, 3, 10)
	if e == nil {
		t.Errorf("NewSudokuCell(2, 3, 10): Expected to throw error because of invalid value")
	}
}

func TestSudokuCell_Value(t *testing.T) {
	cell, _ := NewSudokuCell(4, 6, 3)
	if cell.Value() != 3 {
		t.Errorf("cell.Value(): Expected 3, but received %d", cell.Value())
	}

	cell, _ = NewSudokuCell(4, 6, 9)
	if cell.Value() != 9 {
		t.Errorf("cell.Value(): Expected 9, but received %d", cell.Value())
	}

	cell, _ = NewSudokuCell(4, 6, 0)
	if cell.Value() != 0 {
		t.Errorf("cell.Value(): Expected 0, but received %d", cell.Value())
	}
}

func TestSudokuCell_String(t *testing.T) {
	cell, _ := NewSudokuCell(3, 4, 5)

	if cell.String() != "5 at (3, 4)" {
		t.Errorf("cell.String(): Expected '5 at (3, 4)' but received '%s'", cell.String())
	}
}

func TestSudokuCell_SetValue(t *testing.T) {
	// We need to initialize a complete Sudoku first for SetValue to work properly.
	sudoku := EmptySudoku()

	// Get middle Sudoku cell
	middle := sudoku.Cells[4][4]

	if middle.Value() != 0 {
		t.Errorf("SudokuCell.Value(): should have been an empty value cell but its value was %d", middle.Value())
	}

	middle.SetValue(2)
	if middle.Value() != 2 {
		t.Errorf("SudokuCell.Value(): should give 2, but gave %d", middle.Value())
	}
}

func TestSudokuCell_SetValue_InvalidInput(t *testing.T) {
	// We need to initialize a complete Sudoku first for SetValue to work properly.
	sudoku := EmptySudoku()

	// Get middle Sudoku cell
	middle := sudoku.Cells[4][4]

	middle.SetValue(-1)
	if middle.Value() != 0 {
		t.Errorf("SudokuCell.SetValue(-1): Expected to set value to 0")
	}

	middle.SetValue(10)
	if middle.Value() != 0 {
		t.Errorf("SudokuCell.SetValue(10): Expected to set value to 0")
	}
}

func TestSudokuCell_SetValue_ValueAlreadySet(t *testing.T) {
	// We need to initialize a complete Sudoku first for SetValue to work properly.
	sudoku := EmptySudoku()

	// Get middle Sudoku cell
	middle := sudoku.Cells[4][4]

	if middle.Value() != 0 {
		t.Errorf("Value should have been 0")
	}

	middle.SetValue(0)
	if middle.Value() != 0 {
		t.Errorf("Value should have not updated")
	}
}

func TestSudokuCell_SetValue_NeighbourNotification(t *testing.T) {
	// We need to initialize a complete Sudoku first for SetValue to work properly.
	sudoku, e := LoadSudoku(&[][]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{1, 0, 0, 0, 3, 0, 3, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
	})
	if e != nil {
		t.Errorf("Cannot proceed test, as the Sudoku could not be loaded properly")
	}

	// Get middle Sudoku cell
	middle := sudoku.Cells[4][4]
	if middle.Value() != 3 {
		t.Errorf("Middle cell value should have been 3, but is %d", middle.Value())
	}

	anotherCell := sudoku.Cells[4][0]
	if anotherCell.Value() != 1 {
		t.Errorf("Value of cell should be 1, but is %d", anotherCell.Value())
	}

	if occurrences, taken := anotherCell.taken[3]; !taken || occurrences != 2 {
		t.Errorf("Cell should have found value 3 to be already taken two times")
	}

	middle.SetValue(0)
	if occurrences, taken := anotherCell.taken[3]; !taken || occurrences != 1 {
		t.Errorf("Cell should have found value 3 to be still used by another cell")
	}

	other3Cell := sudoku.Cells[4][6]
	other3Cell.SetValue(0)
	if occurrences, taken := anotherCell.taken[3]; taken || occurrences != 0 {
		t.Errorf("Cell should have found value 3 to be free")
	}

	middle.SetValue(1)
	if occurrences, taken := other3Cell.taken[1]; !taken || occurrences != 2 {
		t.Errorf("The other 3 valued cell should have observed two cells in its neighbours to have value 1")
	}
}

func TestSudokuCell_IsEmpty(t *testing.T) {
	cell, _ := NewSudokuCell(4, 6, 3)

	if cell.IsEmpty() {
		t.Errorf("Cell with value 3 is not empty!")
	}

	cell, _ = NewSudokuCell(4, 6, 0)

	if !cell.IsEmpty() {
		t.Errorf("Cell with value 0 is considered empty!")
	}
}

func TestSudokuCell_HasCollision(t *testing.T) {
	sudoku, _ := LoadSudoku(&[][]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{1, 0, 0, 0, 3, 0, 3, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
	})

	middle := sudoku.Cells[4][4]
	if !middle.HasCollision() {
		t.Errorf("Expected middle cell value to collide with the value of another cell")
	}

	middle.SetValue(9)
	if middle.HasCollision() {
		t.Errorf("Expected cell to not have an collision")
	}

	middle.SetValue(0)
	if middle.HasCollision() {
		t.Errorf("Cell cannot have a collision when its value is empty")
	}
}

func TestSudokuCell_PossibleValues(t *testing.T) {
	sudoku, _ := LoadSudoku(&[][]int{
		{0, 1, 2, 3, 4, 5, 6, 7, 8},
		{0, 0, 0, 0, 1, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{2, 0, 0, 0, 0, 0, 3, 9, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
	})

	middle := sudoku.Cells[4][4]
	possibleValues := middle.PossibleValues()
	expected := map[int]bool{
		5: false,
		6: false,
		7: false,
		8: false,
	}

	for _, value := range possibleValues {
		taken, ok := expected[value]

		if !ok || taken {
			t.Errorf("Expected value %d to be possible", value)
		}

		expected[value] = true
	}

	// Check if all values in the expected map have been there!
	for key, value := range expected {
		if !value {
			t.Errorf("Expected value %d to be found in the possible values", key)
		}
	}
}

func TestSudokuCell_Neighbours(t *testing.T) {
	sudoku, _ := LoadSudoku(&[][]int{
		{0, 1, 2, 3, 4, 5, 6, 7, 8},
		{0, 0, 0, 0, 1, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{2, 0, 0, 0, 0, 0, 3, 9, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
	})

	cell := sudoku.Cells[0][0]
	neighbours := cell.Neighbours()

	if neighbours != &cell.neighbours {
		t.Errorf("Expected cell neighbours to be the same object")
	}
}
