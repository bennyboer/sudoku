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

	beforeS := fmt.Sprintf("%v", sudoku) // TODO Remove
	//debugPrintPossibleValues(possibleValueLookupRef) // TODO Remove

	iteration := 1
	for patternIndex := 0; patternIndex < len(patterns); {
		currentPattern := patterns[patternIndex]

		changed := currentPattern.Apply(sudoku, possibleValueLookupRef)

		fmt.Printf("\nITERATION %d WITH PATTERN %d -----\n", iteration, patternIndex) // TODO Remove

		if sudoku.IsCompleteAndValid() {
			return true, nil // Early exit because the solver finished its job
		}

		if changed {
			// Go back to the first pattern, since the Sudoku has been changed by the current pattern.
			// New opportunities may have risen.
			patternIndex = 0
		} else {
			patternIndex++
		}

		iteration++
	}

	afterS := fmt.Sprintf("%v", sudoku) // TODO Remove

	fmt.Printf("Before:\n%s\n\n", beforeS) // TODO Remove
	fmt.Printf("After:\n%s\n\n", afterS)   // TODO Remove

	return sudoku.IsCompleteAndValid(), nil
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
		&pattern.NakedSingle{},
		&pattern.HiddenSingle{},
		&pattern.NakedPair{},
		&pattern.HiddenPair{},
		&pattern.NakedTriple{},
	}[:]
}
