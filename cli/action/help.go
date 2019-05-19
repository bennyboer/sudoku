package action

import (
	"flag"
	"fmt"
)

// Action printing help for a action.
type Help struct {
	// The flag set to use to execute the action.
	flagSet *flag.FlagSet
	// Slice of all actions available
	actions []Action
}

func NewHelp(actions []Action) *Help {
	return &Help{
		actions: actions,
	}
}

func (a *Help) Name() string {
	return "help"
}

func (a *Help) FlagSet() *flag.FlagSet {
	if a.flagSet == nil {
		a.flagSet = flag.NewFlagSet(a.Name(), flag.ExitOnError)
	}

	return a.flagSet
}

func (a *Help) CanExecute() bool {
	return a.flagSet.Parsed() && a.flagSet.NArg() == 1
}

func (a *Help) Execute() {
	actionName := a.flagSet.Arg(0)

	for _, a := range a.actions {
		if a.Name() == actionName {
			a.FlagSet().PrintDefaults()
			return
		}
	}

	fmt.Printf("Could not find help for action with the name '%s'\n", actionName)
}
