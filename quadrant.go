package barneshut

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
)

// Long term Quadrant becomes a node.
type Quadrant struct {
	Width float64   `json:"Width"` // the Width of the quadrant
	LL    []float64 `LL:"Base"`    // the "lower-left" point of this quadrant
}

func (q Quadrant) String() string {
	return fmt.Sprintf("Quadrant{Width: %v, LL: %v}", q.Width, q.LL)
}

func (q Quadrant) ContainsBody(body *Body) bool {
	return q.Contains(body.R)
}

func (q Quadrant) Contains(point []float64) bool {

	for dim := 0; dim < len(point); dim++ {
		LL_open := q.LL[dim]
		LL_close := LL_open + q.Width

		if !((LL_open < point[dim]) && (point[dim] < LL_close)) {
			return false
		}
	}

	return true
}

func (q Quadrant) Equals(oq Quadrant) bool {

	if len(oq.LL) != len(q.LL) {
		// Added because of cases where [.5] == [.5, 0]
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

func (q Quadrant) Subdivide() []Quadrant {
	log.Infof("Subdividing quadrant: %s", q)

	cnt_new_quadrants := Pow(2, len(q.LL))

	var quadrants []Quadrant
	var offset []int

	for i := 0; i < cnt_new_quadrants; i++ {
		offset = Index2Offset(i, len(q.LL))
		log.Infof("Working on offset %v", offset)

		var new_points []float64

		j := 0
		for {
			new_points = append(new_points, q.LL[j]+float64(offset[j])*(q.Width/2))
			if j == len(offset)-1 {
				break
			}
			j++
		}

		log.Infof("Got new points: %v", new_points)
		quadrants = append(quadrants, Quadrant{LL: new_points, Width: q.Width / 2})
	}

	return quadrants
}
