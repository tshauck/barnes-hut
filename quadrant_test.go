package barneshut

import (
	log "github.com/Sirupsen/logrus"
	"testing"
)

func init() {
	//Move this to a different place, like main.go?
	log.SetLevel(log.WarnLevel)
}

func TestContains(t *testing.T) {

	q := Quadrant{
		Width: 1,
		LL:    []float64{0.0, 0.0},
	}

	test_points := [][]float64{
		[]float64{2.2, 2.2},
		[]float64{-2.2, -2.2},
		[]float64{0.5, 0.5},
	}
	expected := []bool{false, false, true}

	for i := range test_points {
		if q.Contains(test_points[i]) != expected[i] {
			t.Errorf("Error for point %v, got: %v", test_points[i], expected[i])
		}
	}
}

func TestEquals(t *testing.T) {

	q := Quadrant{
		Width: 1,
		LL:    []float64{0.0, 0.0},
	}

	q2 := Quadrant{
		Width: 1,
		LL:    []float64{0.0, 0.0},
	}

	if !q.Equals(q2) {
		t.Errorf("Quadrants are not equal!")
	}

	q3 := Quadrant{
		Width: 2, // Different width than q
		LL:    []float64{0.0, 0.0},
	}

	q4 := Quadrant{
		Width: 1,
		LL:    []float64{0.0, 0.0, 0.0}, // Different dimensions
	}

	q5 := Quadrant{
		Width: 1,
		LL:    []float64{0.0, 0.1}, // Different dimensions
	}
	qs := []Quadrant{q3, q4, q5}

	for i := range qs {
		if q.Equals(qs[i]) {
			t.Errorf("Q(%s) != Q(%s)", q, qs[i])
		}
	}

}

func TestQuadrantSubdivide(t *testing.T) {
	LL := Quadrant{Width: 1, LL: []float64{0.0, 0.0}}
	actual := LL.Subdivide()

	expected := []Quadrant{
		Quadrant{Width: .5, LL: []float64{0.0, 0.0}},
		Quadrant{Width: .5, LL: []float64{0.0, 0.5}},
		Quadrant{Width: .5, LL: []float64{0.5, 0.0}},
		Quadrant{Width: .5, LL: []float64{0.5, 0.5}},
	}

	for i := range actual {
		if !actual[i].Equals(expected[i]) {
			t.Errorf("Actual Quadrant (%s) != Expected Quadrant (%s)", actual[i], expected[i])
		}
	}
}
