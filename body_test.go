package barneshut

import (
	"math"
	"testing"
)

func TestBodyEquals(t *testing.T) {

	b1 := &Body{
		R:    []float64{0.0, 0.0},
		V:    []float64{0.0, 0.0},
		F:    []float64{0.0, 0.0},
		Mass: 0,
	}

	b2 := &Body{
		R:    []float64{0.0, 0.0},
		V:    []float64{0.0, 0.0},
		F:    []float64{0.0, 0.0},
		Mass: 0,
	}

	if !b1.Equals(b2) {

		t.Errorf("b1(%v) != b2(%v)", b1, b2)

	}

}

func TestDistanceTo(t *testing.T) {
	expected := math.Sqrt(2)
	b1 := Body{
		R: []float64{1.0, 1.0},
	}

	b2 := Body{
		R: []float64{0.0, 0.0},
	}

	actual := b1.DistantceTo(b2)

	if expected != actual {
		t.Errorf("Distance was %f, expected %f", actual, expected)
	}
}

func TestUpdate(t *testing.T) {
	b1 := &Body{
		R:    []float64{1.0, 1.0},
		V:    []float64{1.0, 1.0},
		F:    []float64{1.0, 1.0},
		Mass: 2,
	}

	b1.Update(1)

	b2 := &Body{
		R:    []float64{2.5, 2.5},
		V:    []float64{1.5, 1.5},
		F:    []float64{1.0, 1.0},
		Mass: 2,
	}

	if !b1.Equals(b2) {
		t.Errorf("b1(%v) != b2(%v)", b1, b2)
	}
}

func TestAddForce(t *testing.T) {
	b1 := Body{
		R:    []float64{0.1, 0.1},
		V:    []float64{0.0, 0.0},
		F:    []float64{0.0, 0.0},
		Mass: 1,
	}

	b2 := Body{
		R:    []float64{0.5, 0.5},
		V:    []float64{0.0, 0.0},
		F:    []float64{0.0, 0.0},
		Mass: 2,
	}

	b1.AddForce(b2)

	expectedForce := []float64{2.94775139e-10, 2.94775139e-10}

	for i := 0; i < len(b1.F); i++ {
		if expectedForce[i]-b1.F[i] > 1e-5 {
			t.Errorf("Not Equal: %v != %v", expectedForce[i], b1.F[i])
		}
	}

}

func TestAddBody(t *testing.T) {

	b1 := &Body{
		R:    []float64{0.1, 0.1},
		V:    []float64{0.0, 0.0},
		F:    []float64{0.0, 0.0},
		Mass: 1,
	}

	b2 := &Body{
		R:    []float64{0.5, 0.5},
		V:    []float64{0.0, 0.0},
		F:    []float64{0.0, 0.0},
		Mass: 2,
	}

	b1.AddBody(b2)

	expectedBody := &Body{
		R:    []float64{1.1 / 3, 1.1 / 3},
		V:    []float64{0.0, 0.0},
		F:    []float64{0.0, 0.0},
		Mass: 3,
	}

	if !b1.Equals(expectedBody) {
		t.Errorf("Not Equal: %v != %v", expectedBody, b1)
	}

}

func TestInQuadrant(t *testing.T) {
	b := Body{
		R:    []float64{0.5, 0.5},
		V:    []float64{0.0, 0.0},
		F:    []float64{0.0, 0.0},
		Mass: 2,
	}
	q := Quadrant{
		LL:    []float64{0.0, 0.0},
		Width: 1.0,
	}

	if !b.InQuadrant(q) {
		t.Errorf("body %v not in quadrant %v", b, q)
	}

}
