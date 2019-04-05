package strategy

import (
	"fmt"
	"github.com/ob-algdatii-ss19/leistungsnachweis-sudo/model"
	"github.com/ob-algdatii-ss19/leistungsnachweis-sudo/solver/strategy/pattern"
	"github.com/ob-algdatii-ss19/leistungsnachweis-sudo/solver/strategy/util"
	"sort"
)

// Strategy solver using the Sudoku solving patterns in order
// to solve the Sudoku.
type Solver struct{}

// Solve the passed Sudoku using the strategy solver.
func (s *Solver) Solve(sudoku *model.Sudoku) (bool, error) {
	patterns := initPatterns()
	possibleValueLookupRef := util.PreparePossibleValueLookup(sudoku)

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

// Initialize the patterns slice.
// The patterns are used to be applied on a Sudoku in order to solve it.
func initPatterns() []pattern.Pattern {
	return []pattern.Pattern{
		&pattern.Trivial{},
	}[:]
}
