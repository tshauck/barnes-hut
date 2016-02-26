package barneshut

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"math"
)

const G = 6.67e-11
const EPS = 1e-5

type Body struct {
	r    []float64 // position vector
	v    []float64 // velocity vector
	f    []float64 // force vector
	mass float64
}

func (b Body) String() string {
	return fmt.Sprintf("Body{r: %v, v: %v, f: %f, mass: %v}", b.r, b.v, b.f, b.mass)
}

func (b Body) Update(dt float64) {

	for i := 0; i < len(b.r); i++ {
		b.v[i] += dt * b.f[i] / b.mass
		b.r[i] += dt * b.v[i]
	}
}

func (b Body) DistantceTo(ob Body) float64 {

	distance := 0.0

	for i := 0; i < len(b.r); i++ {
		distance += math.Pow(ob.r[i]-b.r[i], 2)
	}

	return math.Sqrt(distance)
}

func (b Body) Equals(ob Body) bool {

	compare_floats := func(a float64, b float64) bool {
		return math.Abs(a-b) < EPS
	}

	for i := 0; i < len(b.r); i++ {
		if !(compare_floats(b.r[i], ob.r[i]) && compare_floats(b.v[i], ob.v[i]) && compare_floats(b.f[i], ob.f[i])) {
			return false
		}
	}

	return true
}

func (b Body) AddForce(ob Body) {

	// TODO(trent): see test, may need some work.
	overall_distance := b.DistantceTo(ob)
	totalForce := (G * b.mass * ob.mass) / math.Pow(overall_distance, 2)

	log.Debugf("Distance between bodies is (%f), and its Force is (%f)", overall_distance, totalForce)

	for i := 0; i < len(b.r); i++ {
		d := ob.r[i] - b.r[i]
		b.f[i] += totalForce * d / overall_distance
	}
}

func (b Body) AddBody(ob Body) Body {

	log.Debugf("Adding other body (%v) to (%v)", ob, b)
	total_mass := b.mass + ob.mass

	dims := len(b.r)
	var new_rs []float64
	var new_r float64

	for dim := 0; dim < dims; dim++ {
		log.Debugf("On dim (%d) adding b.r %v to ob.r %v", dim, b.r[dim], ob.r[dim])
		new_r = ((b.r[dim] * b.mass) + (ob.r[dim] * ob.mass)) / total_mass
		new_rs = append(new_rs, new_r)
	}

	return Body{
		r:    new_rs,
		v:    b.v,
		f:    b.f,
		mass: total_mass,
	}
}

func (b Body) InQuadrant(q Quadrant) bool {

	dims := len(b.r)

	for dim := 0; dim < dims; dim++ {
		upper := q.base[dim] + q.width
		lower := q.base[dim]

		if !((lower < b.r[dim]) && (b.r[dim] < upper)) {
			return false
		}
	}

	return true
}
