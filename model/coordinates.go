package model

import "fmt"

/// 2-dimensional coordinates.
type Coordinates struct {
	Row    int
	Column int
}

// Retrieve a String representation of the coordinates.
func (c *Coordinates) String() string {
	return fmt.Sprintf("(%d, %d)", c.Row, c.Column)
}
