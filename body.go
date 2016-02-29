// Copyright (C) 2016 Trent Hauck - All Rights Reserved

package barneshut

import (
	"encoding/json"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"math"
)

// G is the gravitational constant, https://en.wikipedia.org/wiki/Gravitational_constant
const G = 6.67e-11

// EPS is the allowable epsilon for checking equal floating point values.
const EPS = 1e-5

// Body holds information for a "body" in space.  This includes the position vector, R,
// the velocity vector, V, the force vector, F, the mass, Mass, and a label, Label, for the
// name of the body for bookkeeping.
type Body struct {
	R     []float64 `json:"R"` // position vector
	V     []float64 `json:"V"` // velocity vector
	F     []float64 `json:"F"` // force vector
	Mass  float64   `json:"Mass"`
	Label string    `json:"Label"`
}

// Json returns the byte array of the json representation of a Body.
func (b Body) Json() []byte {
	by, err := json.Marshal(b)
	if err != nil {
		panic(err)
	}
	return by
}

// String returns a string in pretty print form for the Body.
func (b Body) String() string {
	return fmt.Sprintf("Body{Label: %s, R: %v, V: %v, F: %f, Mass: %v}",
		b.Label, b.R, b.V, b.F, b.Mass)
}

// Update moves the Velocity and Position "forward" one time step based on the
// current Velocity, Force, and Mass.
func (b Body) Update(dt float64) {
	for i := 0; i < len(b.R); i++ {
		b.V[i] += dt * b.F[i] / b.Mass
		b.R[i] += dt * b.V[i]
	}
}

// DistanceTo calculates the euclidian distance between Body, b, and Body, ob.
func (b Body) DistanceTo(ob Body) float64 {
	distance := 0.0

	for i := 0; i < len(b.R); i++ {
		distance += math.Pow(ob.R[i]-b.R[i], 2)
	}

	return math.Sqrt(distance)
}

// EqualPosition returns a bool that's true if the Position and
// Mass for Body, b, and Body, ob, are equal.
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

// Equals returns a bool that's true if the Position, Mass, Force
// and Velocity for Body, b, and Body, ob, are equal.
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

// AddForce updates the Force vector for the calling struct based on the
// metrics associated with Body, ob.
func (b *Body) AddForce(ob Body) {
	fmt.Printf("Adding force from %s to %s.\n", ob.Label, b.Label)
	overall_distance := b.DistanceTo(ob)
	totalForce := (G * b.Mass * ob.Mass) / math.Pow(overall_distance, 2)

	log.Debugf("Distance between bodies is (%f), and its Force is (%f)", overall_distance, totalForce)

	var newForce []float64
	for i := 0; i < len(b.R); i++ {
		d := ob.R[i] - b.R[i]
		newForce = append(newForce, b.F[i]+(totalForce*d/overall_distance))
	}

	b.F = newForce
	fmt.Printf("%s's new force: %v\n", b.Label, b.F)
}

// AddBody updates body, b, with a new mass and position.  This is effectively
// updating body, b, so its new mass and position are at the center of mass between
// b and ob.
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
	b.Label = fmt.Sprintf("%s%s", b.Label, ob.Label)
}

// InQuadrant returns a boolean if the Body, b, is in Quadrant, q.
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
