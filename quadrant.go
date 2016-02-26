package barneshut

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
)

// Long term Quadrant becomes a node.
type Quadrant struct {
	width float64   // the width of the quadrant
	base  []float64 // the "lower-left" point of this quadrant
}

func (q Quadrant) String() string {
	return fmt.Sprintf("Quadrant{width: %v, base: %v}", q.width, q.base)
}

func (q Quadrant) Contains(point []float64) bool {

	log.Debugf("Checking if point(%v) is in quadrant in %s", point, q)

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

	log.Debugf("Checking if q(%s) is equal to oq(%s)", q, oq)

	if len(oq.base) != len(q.base) {
		// Added because of cases where [.5] == [.5, 0]
		return false
	}

	if q.width != oq.width {
		return false
	}

	for i := range q.base {
		if q.base[i] != oq.base[i] {
			return false
		}
	}

	return true
}

func (q Quadrant) Subdivide() []Quadrant {

	log.Debugf("Subdividing quadrant (%s).", q)

	cnt_new_quadrants := Pow(2, len(q.base))

	var quadrants []Quadrant
	var offset []int

	for i := 0; i < cnt_new_quadrants; i++ {
		offset = Index2Offset(i, len(q.base))

		var new_points []float64

		j := len(offset) - 1
		for {
			new_points = append(new_points, q.base[j]+float64(offset[j])*(q.width/2))
			if j == 0 {
				break
			}
			j = j - 1
		}

		quadrants = append(quadrants, Quadrant{base: new_points, width: q.width / 2})
	}

	return quadrants
}
