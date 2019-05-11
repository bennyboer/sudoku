package strategy

import (
	"github.com/ob-algdatii-ss19/leistungsnachweis-sudo/model"
	"github.com/ob-algdatii-ss19/leistungsnachweis-sudo/solver/strategy/pattern"
	"github.com/ob-algdatii-ss19/leistungsnachweis-sudo/solver/strategy/util"
)

// Strategy solver using the Sudoku solving patterns in order
// to solve the Sudoku.
type Solver struct{}

// Solve the passed Sudoku using the strategy solver.
func (s *Solver) Solve(sudoku *model.Sudoku) (bool, error) {
	patterns := initPatterns()
	patternCount := len(patterns)
	fallback := pattern.Fallback{}

	possibleValueLookupRef := util.PreparePossibleValueLookup(sudoku)

	for patternIndex := 0; patternIndex < patternCount+1; {
		isFallback := patternIndex >= patternCount
		var changed bool
		if !isFallback {
			// Try the pattern
			changed = patterns[patternIndex].Apply(sudoku, possibleValueLookupRef)
		} else {
			changed = fallback.Apply(sudoku, possibleValueLookupRef)
		}

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
	}

	return sudoku.IsCompleteAndValid(), nil
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
		&pattern.HiddenTriple{},
		&pattern.NakedQuadruple{},
		&pattern.HiddenQuadruple{},
		&pattern.RowBlockCheck{},
		&pattern.BlockRowCheck{},
		&pattern.XWing{},
	}[:]
}
