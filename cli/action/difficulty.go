package action

import (
	"flag"
	"fmt"
	"github.com/ob-algdatii-ss19/leistungsnachweis-sudo/io/read"
	"github.com/ob-algdatii-ss19/leistungsnachweis-sudo/solver"
)

// Action measuring a Sudokus difficulty.
type Difficulty struct {
	// The flag set to use to execute the action.
	flagSet *flag.FlagSet

	// The input to measure (Sudoku in a text file).
	input *string
}

func NewDifficulty() *Difficulty {
	return &Difficulty{}
}

func (a *Difficulty) Name() string {
	return "difficulty"
}

func (a *Difficulty) FlagSet() *flag.FlagSet {
	if a.flagSet == nil {
		a.flagSet = flag.NewFlagSet(a.Name(), flag.ExitOnError)

		a.input = a.flagSet.String(
			"in",
			"sudoku.txt",
			"The Sudoku in a text file to use",
		)
	}

	return a.flagSet
}

func (a *Difficulty) CanExecute() bool {
	return a.flagSet.Parsed()
}

func (a *Difficulty) Execute() {
	fmt.Printf(`Sudoku measuring started...
-----
Sudoku file path: %s
-----
`, *a.input)

	fmt.Println("Loading Sudoku...")
	reader := read.SudokuFileReader{
		FilePath: a.input,
	}
	sudoku, err := reader.Read()
	if err != nil {
		fmt.Printf("Could not load the Sudoku from file\nError:\n%s\n", err.Error())
		return
	}

	fmt.Printf(`Sudoku successfully loaded:
-----
%s
-----
`, sudoku.String())

	fmt.Println("Measuring Sudoku difficulty...")
	difficulty, err := solver.MeasureDifficulty(sudoku)
	if err != nil {
		fmt.Printf(`Could not measure difficulty. Error:
%s`, err.Error())
		return
	}

	fmt.Printf(`-----
Sudoku difficulty is %f
This translates to %s
`, difficulty, a.translateDifficulty(difficulty))
}

// Translate the passed continuous difficulty to a discrete difficulty step.
func (a *Difficulty) translateDifficulty(difficulty float64) string {
	switch {
	case difficulty < 0.1:
		return "VERY EASY"
	case difficulty < 0.2:
		return "EASY"
	case difficulty < 0.4:
		return "MODERATE"
	case difficulty < 0.6:
		return "HARD"
	case difficulty < 0.8:
		return "VERY HARD"
	default:
		return "INSANE"
	}
}
