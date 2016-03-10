// Copyright (C) 2016 Trent Hauck - All Rights Reserved

package barneshut

import (
	"fmt"
)

// Quadrant is a struct that contains with width of the quadrant and the
// point where the lower-left (LL) is. The quadrant is assumed to be a square.
type Quadrant struct {
	Width float64   `json:"Width"` // the Width of the quadrant
	LL    []float64 `LL:"Base"`    // the "lower-left" point of this quadrant
}

// String returns a pretty printed version of the Quadrant type.
func (q Quadrant) String() string {
	return fmt.Sprintf("Quadrant{Width: %v, LL: %v}", q.Width, q.LL)
}

// ContainsBody returns true if the body (of type *Body) is in the Quadrant.
func (q Quadrant) ContainsBody(body *Body) bool {
	return q.Contains(body.R)
}

// Contains returns true if the point is within its boarders.
func (q Quadrant) Contains(point []float64) bool {

	var LLOpen float64
	var LLClose float64

	for dim := 0; dim < len(point); dim++ {
		LLOpen = q.LL[dim]
		LLClose = LLOpen + q.Width

		if !((LLOpen < point[dim]) && (point[dim] < LLClose)) {
			return false
		}
	}

	return true
}

// Equals determines if the calling Quadrant is equal to the passed Quadrant.
func (q Quadrant) Equals(oq Quadrant) bool {

	if len(oq.LL) != len(q.LL) {
		return false
	}

	if q.Width != oq.Width {
		return false
	}

	for i := range q.LL {
		if q.LL[i] != oq.LL[i] {
			return false
		}
	}

	return true
}

// Subdivide creates a N-Tree within the calling quadrant and returns
// those Quadrants as a list.
func (q Quadrant) Subdivide() []Quadrant {

	cnt_new_quadrants := Pow(2, len(q.LL))

	var quadrants []Quadrant
	var offset []int

	for i := 0; i < cnt_new_quadrants; i++ {
		offset = Index2Offset(i, len(q.LL))

		var new_points []float64

		j := 0
		for {
			new_points = append(new_points, q.LL[j]+float64(offset[j])*(q.Width/2))
			if j == len(offset)-1 {
				break
			}
			j++
		}

		quadrants = append(quadrants, Quadrant{LL: new_points, Width: q.Width / 2})
	}

	return quadrants
}

func (q Quadrant) Min() float64 {
	return q.LL[0]
}

func (q Quadrant) Max() float64 {
	return q.LL[0] + q.Width
}
