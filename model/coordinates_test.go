package model

import "testing"

func TestCoordinates_String(t *testing.T) {
	c := Coordinates{Row: 4, Column: 6}
	str := c.String()

	if str != "(4, 6)" {
		t.Errorf("Coordinates.String() = %s; wanted (4, 6)", str)
	}

	c = Coordinates{Row: -5, Column: 234}
	str = c.String()

	if str != "(-5, 234)" {
		t.Errorf("Coordinates.String() = %s; wanted (-5, 234)", str)
	}
}
