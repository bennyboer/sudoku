package action

import (
	"flag"
	"fmt"
)

type Solve struct {
	// The flag set to use to execute the action.
	flagSet *flag.FlagSet

	// The input to solve (Sudoku in a text file).
	input *string
}

func NewSolve() *Solve {
	return &Solve{}
}

func (a *Solve) Name() string {
	return "solve"
}

func (a *Solve) FlagSet() *flag.FlagSet {
	if a.flagSet == nil {
		a.flagSet = flag.NewFlagSet(a.Name(), flag.ExitOnError)

		a.input = a.flagSet.String("in", "sudoku.txt", "The Sudoku in a text file to use")
	}

	return a.flagSet
}

func (a *Solve) CanExecute() bool {
	return a.flagSet.Parsed()
}

func (a *Solve) Execute() {
	fmt.Println("I should solve something...")
	fmt.Printf("Sudoko in %s\n", *a.input)
}
