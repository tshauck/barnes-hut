package barneshut

import (
	"math"
	"testing"
)

func TestEquals(t *testing.T) {

	b1 := Body{
		r:    []float64{0.0, 0.0},
		v:    []float64{0.0, 0.0},
		f:    []float64{0.0, 0.0},
		mass: 0,
	}

	b2 := Body{
		r:    []float64{0.0, 0.0},
		v:    []float64{0.0, 0.0},
		f:    []float64{0.0, 0.0},
		mass: 0,
	}

	if !b1.Equals(b2) {

		t.Errorf("b1(%v) != b2(%v)", b1, b2)

	}

}

func TestDistanceTo(t *testing.T) {
	expected := math.Sqrt(2)
	b1 := Body{
		r: []float64{1.0, 1.0},
	}

	b2 := Body{
		r: []float64{0.0, 0.0},
	}

	actual := b1.DistantceTo(b2)

	if expected != actual {
		t.Errorf("Distance was %f, expected %f", actual, expected)
	}
}

func TestUpdate(t *testing.T) {
	b1 := Body{
		r:    []float64{1.0, 1.0},
		v:    []float64{1.0, 1.0},
		f:    []float64{1.0, 1.0},
		mass: 2,
	}

	b1.Update(1)

	b2 := Body{
		r:    []float64{2.5, 2.5},
		v:    []float64{1.5, 1.5},
		f:    []float64{1.0, 1.0},
		mass: 2,
	}

	if !b1.Equals(b2) {
		t.Errorf("b1(%v) != b2(%v)", b1, b2)
	}
}

func TestAddForce(t *testing.T) {
	b1 := Body{
		r:    []float64{0.1, 0.1},
		v:    []float64{0.0, 0.0},
		f:    []float64{0.0, 0.0},
		mass: 1,
	}

	b2 := Body{
		r:    []float64{0.5, 0.5},
		v:    []float64{0.0, 0.0},
		f:    []float64{0.0, 0.0},
		mass: 2,
	}

	b1.AddForce(b2)

	expectedForce := []float64{2.94775139e-10, 2.94775139e-10}

	for i := 0; i < len(b1.f); i++ {
		if expectedForce[i]-b1.f[i] > 1e-5 {
			t.Errorf("Not Equal: %v != %v", expectedForce[i], b1.f[i])
		}
	}

}
