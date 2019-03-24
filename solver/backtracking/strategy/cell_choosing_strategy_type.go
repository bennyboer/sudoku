package strategy

// Strategy type on how to choose the next empty cell for the backtracking algorithm.
type CellChoosingStrategyType int

const (
	// Always try the next cell to the right or the first in the next row.
	Linear CellChoosingStrategyType = 0
	// Choose empty cells randomly.
	Random CellChoosingStrategyType = 1
)
