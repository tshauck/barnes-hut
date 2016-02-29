package barneshut

import (
	"encoding/json"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"math"
)

const G = 6.67e-11
const EPS = 1e-5

type Body struct {
	R     []float64 `json:"R"` // position vector
	V     []float64 `json:"V"` // velocity vector
	F     []float64 `json:"F"` // force vector
	Mass  float64   `json:"Mass"`
	Label string    `json:"Label"`
}

func (b Body) Json() []byte {
	//by, _ := json.MarshalIndent(b, "", "  ")
	by, _ := json.Marshal(b)
	return by
}

func (b Body) String() string {
	return fmt.Sprintf("Body{Label: %s, r: %v, v: %v, f: %f, mass: %v}",
		b.Label, b.R, b.V, b.F, b.Mass)
}

func (b Body) Update(dt float64) {

	for i := 0; i < len(b.R); i++ {
		b.V[i] += dt * b.F[i] / b.Mass
		b.R[i] += dt * b.V[i]
	}
}

func (b Body) DistanceTo(ob Body) float64 {

	distance := 0.0

	for i := 0; i < len(b.R); i++ {
		distance += math.Pow(ob.R[i]-b.R[i], 2)
	}

	return math.Sqrt(distance)
}

func (b Body) EqualPosition(ob *Body) bool {

	compare_floats := func(a float64, b float64) bool {
		return math.Abs(a-b) < EPS
	}

	if b.Mass != ob.Mass {
		return false
	}

	for i := 0; i < len(b.R); i++ {
		if !(compare_floats(b.R[i], ob.R[i])) {
			return false
		}
	}

	return true
}

func (b Body) Equals(ob *Body) bool {

	compare_floats := func(a float64, b float64) bool {
		return math.Abs(a-b) < EPS
	}

	if b.Mass != ob.Mass {
		return false
	}

	for i := 0; i < len(b.R); i++ {
		if !(compare_floats(b.R[i], ob.R[i]) && compare_floats(b.V[i], ob.V[i]) && compare_floats(b.F[i], ob.F[i])) {
			return false
		}
	}

	return true
}

func (b Body) AddForce(ob Body) {

	// TODO(trent): see test, may need some work.
	overall_distance := b.DistanceTo(ob)
	totalForce := (G * b.Mass * ob.Mass) / math.Pow(overall_distance, 2)

	log.Debugf("Distance between bodies is (%f), and its Force is (%f)", overall_distance, totalForce)

	for i := 0; i < len(b.R); i++ {
		d := ob.R[i] - b.R[i]
		b.F[i] += totalForce * d / overall_distance
	}
}

func (b *Body) AddBody(ob *Body) {

	total_mass := b.Mass + ob.Mass

	dims := len(b.R)

	var R float64
	var newR []float64
	for dim := 0; dim < dims; dim++ {
		R = ((b.R[dim] * b.Mass) + (ob.R[dim] * ob.Mass)) / total_mass
		newR = append(newR, R)
	}

	b.R = newR
	b.Mass = total_mass
	b.Label = fmt.Sprintf("%s,%s", b.Label, ob.Label)
}

func (b Body) InQuadrant(q Quadrant) bool {

	dims := len(b.R)

	for dim := 0; dim < dims; dim++ {
		upper := q.LL[dim] + q.Width
		lower := q.LL[dim]

		if !((lower < b.R[dim]) && (b.R[dim] < upper)) {
			return false
		}
	}

	return true
}
