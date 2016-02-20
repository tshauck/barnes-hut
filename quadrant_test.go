package barneshut

import (
	"testing"
)

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

func TestNewQuadrant(t *testing.T) {
	starting_quad := Quadrant{width: 1, base: []float64{0, 0}}

	expected_quads := map[string]Quadrant{
		"NE": Quadrant{width: 0.5, base: []float64{1, 1}},
		"NW": Quadrant{width: 0.5, base: []float64{0, 1}},
		"SW": Quadrant{width: 0.5, base: []float64{0, 0}},
		"SE": Quadrant{width: 0.5, base: []float64{1, 0}},
	}

	for key := range expected_quads {
		new_quad := starting_quad.NewQuadrant(key)
		if !(new_quad.Equals(expected_quads[key])) {
			t.Errorf("Quads not equal for %s, got %v, expected %v", key, new_quad, expected_quads[key])
		}
	}

}
