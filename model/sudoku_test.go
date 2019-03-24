package model

import "testing"

func TestLoadSudoku(t *testing.T) {
	values := [9][9]int{
		{1, 2, 3, 4, 5, 4, 3, 2, 1},
		{2, 3, 4, 5, 6, 5, 4, 3, 2},
		{3, 4, 5, 6, 7, 6, 5, 4, 3},
		{4, 5, 6, 7, 8, 7, 6, 5, 4},
		{5, 6, 7, 8, 9, 8, 7, 6, 5},
		{4, 5, 6, 7, 8, 7, 6, 5, 4},
		{3, 4, 5, 6, 7, 6, 5, 4, 3},
		{2, 3, 4, 5, 6, 5, 4, 3, 2},
		{1, 2, 3, 4, 5, 4, 3, 2, 1},
	}

	sudoku, _ := LoadSudoku(&values)

	for row := 0; row < 9; row++ {
		for column := 0; column < 9; column++ {
			cell := sudoku.Cells[row][column]
			if cell.value != values[row][column] {
				t.Errorf("Value in position (%d, %d) should have been %d; got %d",
					row, column, values[row][column], cell.value)
			}

			if cell.position.Row != row || cell.position.Column != column {
				t.Errorf("Cell position should have been (%d, %d); got (%d, %d)",
					row, column, cell.position.Row, cell.position.Column)
			}
		}
	}
}

func TestLoadSudoku_InvalidUsage(t *testing.T) {
	// Try passing nil pointer as values
	_, e := LoadSudoku(nil)
	if e == nil {
		t.Errorf("Anticipated error that sudoku cannot be loaded from nil pointer values array")
	}

	// Try passing invalid ranged values out of range [0; 9] as values
	_, e = LoadSudoku(&[9][9]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, -1, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
	})
	if e == nil {
		t.Errorf("Anticipated error that sudoku cannot be loaded from invalid ranged values")
	}

	_, e = LoadSudoku(&[9][9]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 10, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
	})
	if e == nil {
		t.Errorf("Anticipated error that sudoku cannot be loaded from invalid ranged values")
	}
}

func TestEmptySudoku(t *testing.T) {
	sudoku := EmptySudoku()

	for row := 0; row < 9; row++ {
		for column := 0; column < 9; column++ {
			cell := sudoku.Cells[row][column]
			if cell.value != 0 {
				t.Errorf("Value in position (%d, %d) should have been 0; got %d",
					row, column, cell.value)
			}
		}
	}
}

func TestSudoku_SaveSudoku(t *testing.T) {
	sudoku := EmptySudoku()

	savedValues := *sudoku.SaveSudoku()

	for row := 0; row < 9; row++ {
		for column := 0; column < 9; column++ {
			savedValue := savedValues[row][column]
			if savedValue != 0 {
				t.Errorf("Expected saved value to be 0; got %d", savedValue)
			}
		}
	}

	valuesToSet := [9][9]int{
		{1, 2, 3, 4, 5, 4, 3, 2, 1},
		{2, 3, 4, 5, 6, 5, 4, 3, 2},
		{3, 4, 5, 6, 7, 6, 5, 4, 3},
		{4, 5, 6, 7, 8, 7, 6, 5, 4},
		{5, 6, 7, 8, 0, 8, 7, 6, 5},
		{4, 5, 6, 7, 8, 7, 6, 5, 4},
		{3, 4, 5, 6, 7, 6, 5, 4, 3},
		{2, 3, 4, 5, 6, 5, 4, 3, 2},
		{1, 2, 3, 4, 5, 4, 3, 2, 1},
	}
	for row := 0; row < 9; row++ {
		for column := 0; column < 9; column++ {
			sudoku.Cells[row][column].SetValue(valuesToSet[row][column])
		}
	}

	savedValues = *sudoku.SaveSudoku()

	for row := 0; row < 9; row++ {
		for column := 0; column < 9; column++ {
			savedValue := savedValues[row][column]
			if savedValue != valuesToSet[row][column] {
				t.Errorf("Expected saved value to be %d; got %d", valuesToSet[row][column], savedValue)
			}
		}
	}
}

func TestSudoku_String(t *testing.T) {
	sudoku, _ := LoadSudoku(&[9][9]int{
		{1, 2, 3, 4, 5, 4, 3, 2, 1},
		{2, 3, 4, 5, 6, 5, 4, 3, 2},
		{3, 4, 5, 6, 7, 6, 5, 4, 3},
		{4, 5, 6, 7, 8, 7, 6, 5, 4},
		{5, 6, 7, 8, 0, 8, 7, 6, 5},
		{4, 5, 6, 7, 8, 7, 6, 5, 4},
		{3, 4, 5, 6, 7, 6, 5, 4, 3},
		{2, 3, 4, 5, 6, 5, 4, 3, 2},
		{1, 2, 3, 4, 5, 4, 3, 2, 1},
	})

	anticipated := `1 2 3   4 5 4   3 2 1
2 3 4   5 6 5   4 3 2
3 4 5   6 7 6   5 4 3

4 5 6   7 8 7   6 5 4
5 6 7   8 _ 8   7 6 5
4 5 6   7 8 7   6 5 4

3 4 5   6 7 6   5 4 3
2 3 4   5 6 5   4 3 2
1 2 3   4 5 4   3 2 1`

	if sudoku.String() != anticipated {
		t.Errorf("Expected sudoku to look like this:\n'%s'\n\nand not like this:\n'%s'", anticipated, sudoku.String())
	}
}

func TestSudoku_IsValid(t *testing.T) {
	sudoku, _ := LoadSudoku(&[9][9]int{
		{1, 2, 3, 4, 5, 4, 3, 2, 1},
		{2, 3, 4, 5, 6, 5, 4, 3, 2},
		{3, 4, 5, 6, 7, 6, 5, 4, 3},
		{4, 5, 6, 7, 8, 7, 6, 5, 4},
		{5, 6, 7, 8, 0, 8, 7, 6, 5},
		{4, 5, 6, 7, 8, 7, 6, 5, 4},
		{3, 4, 5, 6, 7, 6, 5, 4, 3},
		{2, 3, 4, 5, 6, 5, 4, 3, 2},
		{1, 2, 3, 4, 5, 4, 3, 2, 1},
	})
	if sudoku.IsValid() {
		t.Errorf("The Sudoku\n%s\nclaims to be valid, but is not!", sudoku.String())
	}

	sudoku = EmptySudoku()
	if !sudoku.IsValid() {
		t.Errorf("An empty Sudoku is always valid")
	}

	sudoku, _ = LoadSudoku(&[9][9]int{
		{0, 1, 2, 0, 0, 0, 5, 7, 0},
		{6, 0, 0, 5, 0, 1, 0, 0, 4},
		{4, 0, 0, 0, 2, 0, 0, 0, 8},
		{0, 2, 0, 0, 1, 0, 0, 5, 0},
		{0, 0, 4, 9, 0, 7, 8, 0, 0},
		{0, 7, 0, 0, 8, 0, 0, 1, 0},
		{7, 0, 0, 0, 9, 0, 0, 0, 5},
		{5, 0, 0, 4, 0, 8, 0, 0, 6},
		{0, 3, 8, 0, 0, 0, 9, 4, 0},
	})
	if !sudoku.IsValid() {
		t.Errorf("The Sudoku\n%s\nis valid but claims to be invalid", sudoku.String())
	}

	sudoku, _ = LoadSudoku(&[9][9]int{
		{9, 1, 2, 8, 4, 6, 5, 7, 3},
		{6, 8, 3, 5, 7, 1, 2, 9, 4},
		{4, 5, 7, 3, 2, 9, 1, 6, 8},
		{8, 2, 9, 6, 1, 3, 4, 5, 7},
		{1, 6, 4, 9, 5, 7, 8, 3, 2},
		{3, 7, 5, 2, 8, 4, 6, 1, 9},
		{7, 4, 6, 1, 9, 2, 3, 8, 5},
		{5, 9, 1, 4, 3, 8, 7, 2, 6},
		{2, 3, 8, 7, 6, 5, 9, 4, 1},
	})
	if !sudoku.IsValid() {
		t.Errorf("The Sudoku\n%s\nis valid but claims to be invalid", sudoku.String())
	}

	sudoku.Cells[0][0].SetValue(3)
	if sudoku.IsValid() {
		t.Errorf("The Sudoku\n%s\nis invalid but claims to be valid", sudoku.String())
	}
}
