package cli

import (
	"flag"
	"fmt"
	"github.com/ob-algdatii-ss19/leistungsnachweis-sudo/cli/action"
	"os"
)

type CLI struct{}

func NewCLI() *CLI {
	return &CLI{}
}

// Start the CLI.
func (c *CLI) Start() {
	flag.Parse()

	if flag.NArg() < 1 {
		// User MUST provide an argument. Print defaults.
		printDefaults()
		return
	}

	actionName := flag.Arg(0)
	actions := actions()

	// Find the correct action to use
	var actionToUse action.Action
	helpAction := action.NewHelp(actions)
	if actionName == helpAction.Name() {
		actionToUse = helpAction
	}
	if actionToUse == nil {
		for _, a := range actions {
			if a.Name() == actionName {
				actionToUse = a
				break
			}
		}
	}

	if actionToUse == nil {
		fmt.Printf("Action '%s' is unknown\n\n", actionName)
		printDefaults()
		return
	}

	// Try to parse action flag set.
	flagSet := actionToUse.FlagSet()
	err := flagSet.Parse(os.Args[2:])
	if err != nil {
		fmt.Printf("Could not parse flag set for action '%s'. Exiting.\nError: %s\n", actionName, err.Error())
		return
	}

	if !actionToUse.CanExecute() {
		fmt.Println("The arguments you supplied could not be understood.")
		flagSet.PrintDefaults()
		return
	}

	// Execute the action
	actionToUse.Execute()
}

func printDefaults() {
	fmt.Printf(`Sudoku Solver/Generator Tool
-----
Syntax:
   sudoku [action] <flags>
-----
You must provide an action.
The following actions are available:
`)

	actions := actions()
	for _, a := range actions {
		fmt.Printf("   - %s \t\t\t\t| Type `sudoku help %s` for help\n", a.Name(), a.Name())
	}
}

// Get all actions defined for the CLI.
func actions() []action.Action {
	return []action.Action{
		action.NewSolve(),
	}[:]
}
