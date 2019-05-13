package action

import "flag"

// An action the CLI supports.
type Action interface {
	// Name of the action.
	Name() string
	// Get the flag set defined for the action.
	FlagSet() *flag.FlagSet
	// Check whether the action can be executed.
	CanExecute() bool
	// Execute the action.
	Execute()
}
