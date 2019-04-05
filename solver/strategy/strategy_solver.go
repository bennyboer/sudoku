package strategy

import (
	"fmt"
	"github.com/ob-algdatii-ss19/leistungsnachweis-sudo/model"
	"github.com/ob-algdatii-ss19/leistungsnachweis-sudo/solver/strategy/pattern"
	"sort"
)

// Strategy solver using the Sudoku solving patterns in order
// to solve the Sudoku.
type Solver struct{}

// Solve the passed Sudoku using the strategy solver.
func (s *Solver) Solve(sudoku *model.Sudoku) (bool, error) {
	patterns := initPatterns()
	possibleValueLookupRef := preparePossibleValueLookup(sudoku)

	beforeS := fmt.Sprintf("%v", sudoku)             // TODO Remove
	debugPrintPossibleValues(possibleValueLookupRef) // TODO Remove

	changed := true
	currentPattern := patterns[0] // TODO Select pattern automatically
	for changed {
		changed = currentPattern.Apply(sudoku, possibleValueLookupRef)

		debugPrintPossibleValues(possibleValueLookupRef)
		fmt.Println("-----")
	}

	afterS := fmt.Sprintf("%v", sudoku) // TODO Remove

	fmt.Printf("Before:\n%s\n\n", beforeS) // TODO Remove
	fmt.Printf("After:\n%s\n\n", afterS)   // TODO Remove

	return false, nil
}

func debugPrintPossibleValues(valuesPtr *[][]*map[int]bool) {
	matrix := *valuesPtr

	for row := 0; row < 9; row++ {
		for column := 0; column < 9; column++ {
			fmt.Printf("[row: %d, column: %d] | ", row, column)

			if matrix[row][column] != nil {
				lookup := *matrix[row][column]

				possibleValues := make([]int, 0, 9)
				for value, possible := range lookup {
					if possible {
						possibleValues = append(possibleValues, value)
					}
				}

				sort.Ints(possibleValues)

				if len(possibleValues) > 0 {
					fmt.Printf("%v\n", possibleValues)
				} else {
					fmt.Printf("FILLED\n")
				}
			} else {
				fmt.Printf("FILLED\n")
			}
		}
	}
}

// Prepare a lookup of possible values per Sudoku cell.
// It will return a two-dimensional matrix with lookup tables where for each value either true or false is given.
// True means -> value is possible
// False means -> value is not possible
func preparePossibleValueLookup(sudoku *model.Sudoku) *[][]*map[int]bool {
	matrix := make([][]*map[int]bool, model.SudokuSize)

	for i := 0; i < model.SudokuSize; i++ {
		matrix[i] = make([]*map[int]bool, model.SudokuSize)
	}

	cellMatrix := sudoku.Cells

	// Fill with possible values
	for row := 0; row < len(cellMatrix); row++ {
		rowCells := cellMatrix[row]

		for column := 0; column < len(rowCells); column++ {
			cell := rowCells[column]

			if cell.IsEmpty() {
				possibleValuesLookup := make(map[int]bool)

				// Fill possible values in the lookup
				for _, value := range cell.PossibleValues() {
					possibleValuesLookup[value] = true
				}

				// Fill impossible values in the lookup
				for value := 1; value <= 9; value++ {
					_, ok := possibleValuesLookup[value]

					if !ok {
						possibleValuesLookup[value] = false
					}
				}

				matrix[row][column] = &possibleValuesLookup
			}
		}
	}

	return &matrix
}

// Initialize the patterns slice.
// The patterns are used to be applied on a Sudoku in order to solve it.
func initPatterns() []pattern.Pattern {
	return []pattern.Pattern{
		&pattern.Trivial{},
	}[:]
}
