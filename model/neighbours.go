package model

// Collection of all neighbouring cells of a Sudoku cell.
type CellNeighbours struct {
	// Neighbours in the same row as this cell.
	Row []*SudokuCell
	// Neighbours in the same column as this cell.
	Column []*SudokuCell
	// Neighbours in the same block as this cell.
	Block []*SudokuCell
	// All neighbour cells
	All []*SudokuCell
}
