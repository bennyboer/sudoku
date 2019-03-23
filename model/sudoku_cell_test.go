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

	e := middle.SetValue(2)
	if e != nil {
		t.Errorf("SudokuCell.SetValue(2): should throw no error: %s", e.Error())
	}

	if middle.Value() != 2 {
		t.Errorf("SudokuCell.Value(): should give 2, but gave %d", middle.Value())
	}

	e = middle.SetValue(-1)
	if e == nil {
		t.Errorf("SudokuCell.SetValue(-1): Expected to throw an error but it dit not")
	}

	e = middle.SetValue(10)
	if e == nil {
		t.Errorf("SudokuCell.SetValue(10): Expected to throw an error but it dit not")
	}
}
