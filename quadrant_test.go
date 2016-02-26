package barneshut

import (
	log "github.com/Sirupsen/logrus"
	"testing"
)

func init() {
	//Move this to a different place, like main.go?
	log.SetLevel(log.InfoLevel)
}

func TestContains(t *testing.T) {

	q := Quadrant{
		width: 1,
		base:  []float64{0.0, 0.0},
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
		width: 1,
		base:  []float64{0.0, 0.0},
	}

	q2 := Quadrant{
		width: 1,
		base:  []float64{0.0, 0.0},
	}

	if !q.Equals(q2) {
		t.Errorf("Quadrants are not equal!")
	}

}

func TestQuadrantSubdivide(t *testing.T) {
	base := Quadrant{width: 1, base: []float64{0.0, 0.0}}
	actual := base.Subdivide()

	expected := []Quadrant{
		Quadrant{width: .5, base: []float64{0.0, 0.0}},
		Quadrant{width: .5, base: []float64{0.5, 0.0}},
		Quadrant{width: .5, base: []float64{0.0, 0.5}},
		Quadrant{width: .5, base: []float64{0.5, 0.5}},
	}

	for i := range actual {
		if !actual[i].Equals(expected[i]) {
			t.Errorf("Actual Quadrant (%s) != Expected Quadrant (%s)", actual[i], expected[i])
		}
	}
}
