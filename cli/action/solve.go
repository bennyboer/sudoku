package action

import (
	"flag"
	"fmt"
	"github.com/ob-algdatii-ss19/leistungsnachweis-sudo/io/read"
	"github.com/ob-algdatii-ss19/leistungsnachweis-sudo/solver"
	"strings"
)

type Solve struct {
	// The flag set to use to execute the action.
	flagSet *flag.FlagSet

	// The input to solve (Sudoku in a text file).
	input *string

	// Where to store the solved Sudoku to.
	output *string

	// Algorithm to use to solve the Sudoku.
	algorithm *string
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

		a.input = a.flagSet.String(
			"in",
			"sudoku.txt",
			"The Sudoku in a text file to use",
		)

		a.output = a.flagSet.String(
			"out",
			"",
			"Save the solved Sudoku to somewhere (Leave empty if you do not want to save it)")

		solverAlgorithms := *solver.AllSolverAlgorithms()
		algorithmNames := make([]string, 0, len(solverAlgorithms))
		for name := range solverAlgorithms {
			algorithmNames = append(algorithmNames, name)
		}

		a.algorithm = a.flagSet.String(
			"algorithm",
			"strategy",
			fmt.Sprintf("The algorithm to use to solve the Sudoku (%s).", strings.Join(algorithmNames, ", ")),
		)
	}

	return a.flagSet
}

func (a *Solve) CanExecute() bool {
	return a.flagSet.Parsed()
}

func (a *Solve) Execute() {
	fmt.Printf(`Sudoku solver started...
-----
Using algorithm: '%s'
Sudoku file path: '%s'
-----
`, *a.algorithm, *a.input)

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

	s, err := solver.GetAlgorithmForName(*a.algorithm)
	if err != nil {
		fmt.Printf("Could not find solver algorithm with name '%s'\n", *a.algorithm)
		return
	}

	fmt.Println("Trying to solve Sudoku...")
	solvable, err := s.Solve(sudoku)
	if err != nil {
		fmt.Printf(`An error occurred while trying to solve the Sudoku:
%s
`, err.Error())
		return
	}

	if solvable {
		fmt.Printf(`The Sudoku is solvable!
-----
%s
-----
`, sudoku.String())

		if len(*a.output) > 0 {
			fmt.Printf("Storing the result to '%s'\n", *a.output)
		}
	} else {
		fmt.Println("Sudoku is NOT solvable. :(")
	}
}
