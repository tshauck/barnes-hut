package barneshut

import (
	"log"
	"math"
)

const G = 6.67e-11

type Color struct{}

type Body struct {
	r     []float64 // position vector
	v     []float64 // velocity vector
	f     []float64 // force vector
	mass  float64
	color Color
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

	for i := 0; i < len(b.r); i++ {
		if !(b.r[i] == ob.r[i] && b.v[i] == ob.v[i] && b.f[i] == ob.f[i]) {
			return false
		}
	}

	return true
}

func (b Body) AddForce(ob Body) {

	// TODO(trent): see test, may need some work.
	overall_distance := b.DistantceTo(ob)
	totalForce := (G * b.mass * ob.mass) / math.Pow(overall_distance, 2)

	log.Printf("Distance between bodies is (%f), and its Force is (%f)", overall_distance, totalForce)

	for i := 0; i < len(b.r); i++ {
		d := ob.r[i] - b.r[i]
		b.f[i] += totalForce * d / overall_distance
	}
}
