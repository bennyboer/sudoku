package action

import (
	"flag"
	"fmt"
	"github.com/ob-algdatii-ss19/leistungsnachweis-sudo/generator"
	"github.com/ob-algdatii-ss19/leistungsnachweis-sudo/io/write"
)

// Action generating a Sudoku.
type Generate struct {
	// The flag set to use to execute the action.
	flagSet *flag.FlagSet

	// How difficult the generated Sudoku should be
	difficulty *float64

	// Where to write the generated Sudoku (optional).
	output *string
}

func NewGenerate() *Generate {
	return &Generate{}
}

func (a *Generate) Name() string {
	return "generate"
}

func (a *Generate) FlagSet() *flag.FlagSet {
	if a.flagSet == nil {
		a.flagSet = flag.NewFlagSet(a.Name(), flag.ExitOnError)

		a.difficulty = a.flagSet.Float64(
			"difficulty",
			0.2,
			"How difficult the generated Sudoku should be (Range from 0.0 (VERY EASY) to 1.0 (INSANE))",
		)

		a.output = a.flagSet.String(
			"out",
			"",
			"Where to write the generated Sudoku (optional)",
		)
	}

	return a.flagSet
}

func (a *Generate) CanExecute() bool {
	return a.flagSet.Parsed()
}

func (a *Generate) Execute() {
	fmt.Printf(`Sudoku Generator started...
-----
Difficulty: %f
`, *a.difficulty)

	if len(*a.output) != 0 {
		// Output defined -> store the result to the file
		fmt.Printf("Output file path (where to save the generated Sudoku to): '%s'\n", *a.output)
	}

	fmt.Println("-----")

	sudoku := generator.NewBacktrackingGenerator().Generate(*a.difficulty)
	writer := write.SudokuFileWriter{
		FilePath: a.output,
	}

	err := writer.Write(sudoku)
	if err != nil {
		fmt.Printf("There was an error trying to save your sudoku :C %s", err)
	}
}
