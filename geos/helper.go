package geos

// MustCoords is a helper that wraps a call to a function returning ([]Coord, error)
// and panics if the error is non-nil.
func MustCoords(c []Coord, err error) []Coord {
	if err != nil {
		panic(err)
	}
	return c
}
