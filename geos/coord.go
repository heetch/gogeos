package geos

import (
	"fmt"
)

// Coord represents a coordinate in 3-dimensional space.
type Coord struct {
	X, Y, Z float64
}

// NewCoord is the constructor for a Coord object.
func NewCoord(x, y float64) Coord {
	return Coord{x, y, 0}
}

// String returns a (2d) string representation of a Coord.
func (c Coord) String() string {
	return fmt.Sprintf("%f %f", c.X, c.Y)
}
