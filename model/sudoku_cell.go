package model

import (
	"errors"
	"fmt"
)

// Cell of a Sudoku.
type SudokuCell struct {
	// Where the cell is located in its corresponding Sudoku.
	position Coordinates
	// The value of the cell. In range of [0; 9] where 0 means an empty cell.
	value int
	// All neighbouring cells of this cell.
	neighbours CellNeighbours
	// All currently taken (not usable) values for this cell (keys) and how often they are found in the neighbours (value).
	taken map[int]int
}

// Create a new Sudoku cell.
func NewSudokuCell(row int, column int, value int) (*SudokuCell, error) {
	if (row < 0 || row >= SudokuSize) ||
		(column < 0 || column >= SudokuSize) {
		return nil, errors.New("Row and column coordinates of a Sudoku cell must be in range [0; 8]")
	}

	if value < 0 || value > 9 {
		return nil, errors.New("Value is not in range [0; 9]")
	}

	return &SudokuCell{
		position: Coordinates{
			Row:    row,
			Column: column,
		},
		value: value,
	}, nil
}

func (c *SudokuCell) Value() int {
	return c.value
}

// Set a new value to this Sudoku cell.
// Will return nil error if it worked successfully.
func (c *SudokuCell) SetValue(value int) error {
	oldValue := c.value

	// Check if value already set -> do nothing.
	if value == oldValue {
		return nil
	}

	// Check if value in range.
	if value < 0 || value > SudokuSize {
		return errors.New(fmt.Sprintf("Sudoku cell can only hold values between %d and %d", 0, SudokuSize))
	}

	// Notify neighbours that old value is no more used by this cell
	if oldValue != 0 {
		for _, neighbour := range c.neighbours.All {
			occurrences, taken := neighbour.taken[oldValue]

			if taken {
				if occurrences == 1 {
					// Value is no more taken -> Remove from taken map
					delete(neighbour.taken, oldValue)
				} else {
					// Value is taken more than once in the neighbours neighbours -> reduce occurrence count by one
					neighbour.taken[oldValue] = occurrences - 1
				}
			}
		}
	}

	// Notify neighbours that new value is used by this cell and thus no more available (taken)
	if value != 0 {
		for _, neighbour := range c.neighbours.All {
			occurrences, taken := neighbour.taken[value]

			if taken {
				// Value already taken for the neighbour -> increase occurrence count
				neighbour.taken[value] = occurrences + 1
			} else {
				// Value was not taken before for neighbor -> Create new entry in map
				neighbour.taken[value] = 1
			}
		}
	}

	// Finally set the value
	c.value = value

	return nil
}

func (c *SudokuCell) Position() Coordinates {
	return c.position
}

// Initialize this cells lookups using the passed Sudoku cells.
func (c *SudokuCell) Init(cellsPtr *[][]SudokuCell) {
	c.initNeighbours(cellsPtr)
	c.initTaken()
}

// Initialize the neighbour cell lookups for this cell.
func (c *SudokuCell) initNeighbours(cellsPtr *[][]SudokuCell) {
	cells := *cellsPtr

	neighbourCount := SudokuSize - 1

	rowNeighbours := make([]*SudokuCell, neighbourCount)
	columnNeighbours := make([]*SudokuCell, neighbourCount)

	blockNeighbours := make([]*SudokuCell, neighbourCount)
	blockStartRow := (c.position.Row / BlockSize) * BlockSize
	blockStartColumn := (c.position.Column / BlockSize) * BlockSize

	allNeighbours := make([]*SudokuCell, NeighbourCount)

	rowNeighbourCounter := 0
	columnNeighbourCounter := 0
	blockNeighbourCounter := 0
	allNeighbourCounter := 0

	for row := 0; row < SudokuSize; row++ {
		for column := 0; column < SudokuSize; column++ {
			cellPtr := &cells[row][column]
			isNeighbourCell := false

			// Check for row neighbour
			if row == c.position.Row && column != c.position.Column {
				rowNeighbours[rowNeighbourCounter] = cellPtr
				rowNeighbourCounter++
				isNeighbourCell = true
			}

			// Check for column neighbour
			if column == c.position.Column && row != c.position.Row {
				columnNeighbours[columnNeighbourCounter] = cellPtr
				columnNeighbourCounter++
				isNeighbourCell = true
			}

			// Check for block neighbour
			if row >= blockStartRow && row < blockStartRow+BlockSize &&
				column >= blockStartColumn && column < blockStartColumn+BlockSize {
				if row != c.position.Row || column != c.position.Column {
					blockNeighbours[blockNeighbourCounter] = cellPtr
					blockNeighbourCounter++
					isNeighbourCell = true
				}
			}

			if isNeighbourCell {
				allNeighbours[allNeighbourCounter] = cellPtr
				allNeighbourCounter++
			}
		}
	}

	c.neighbours = CellNeighbours{
		Row:    rowNeighbours,
		Column: columnNeighbours,
		Block:  blockNeighbours,
		All:    allNeighbours,
	}
}

// Initialize lookup for all taken (not usable) values for this cell.
func (c *SudokuCell) initTaken() {
	taken := make(map[int]int)

	for _, neighbour := range c.neighbours.All {
		if neighbour.value > 0 {
			occurrences, ok := taken[neighbour.value]

			if ok {
				taken[neighbour.value] = occurrences + 1 // Increment occurrences
			} else {
				taken[neighbour.value] = 1 // Initialize occurrences
			}
		}
	}

	c.taken = taken
}

// Get a String representation of a Sudoku cell.
func (c *SudokuCell) String() string {
	return fmt.Sprintf("%d at %s", c.value, c.position.String())
}
