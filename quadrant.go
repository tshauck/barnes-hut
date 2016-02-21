package barneshut

import (
	log "github.com/Sirupsen/logrus"
)

// Long term Quadrant becomes a node.
type Quadrant struct {
	width float64   // the width of the quadrant
	base  []float64 // the "lower-left" point of this quadrant
}

func (q Quadrant) Contains(point []float64) bool {
	for dim := 0; dim < len(point); dim++ {
		base_open := q.base[dim]
		base_close := base_open + q.width

		if !((base_open < point[dim]) && (point[dim] < base_close)) {
			return false
		}
	}

	return true
}

func (q Quadrant) Equals(oq Quadrant) bool {

	for i := range q.base {
		if q.base[i] != oq.base[i] {
			return false
		}
	}

	if q.width != oq.width {
		return false
	}

	return true
}

func (q Quadrant) NewQuadrant(area string) Quadrant {

	log.Debugf("Creating New Quadrant (%v) in area (%s)", q, area)

	switch {
	case area == "NW":
		return Quadrant{width: q.width / 2, base: []float64{q.base[0], q.base[1] + q.width}}
	case area == "NE":
		return Quadrant{width: q.width / 2, base: []float64{q.base[0] + q.width, q.base[1] + q.width}}
	case area == "SW":
		return Quadrant{width: q.width / 2, base: []float64{q.base[0], q.base[1]}}
	case area == "SE":
		return Quadrant{width: q.width / 2, base: []float64{q.base[0] + q.width, q.base[1]}}
	}

	return Quadrant{}

}
