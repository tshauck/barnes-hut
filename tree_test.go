package barneshut

import (
	"testing"
)

func init() {

}

func TestPow(t *testing.T) {

	attempts := [][]int{
		[]int{1, 2},
		[]int{2, 4},
		[]int{3, 8},
	}

	for i := range attempts {
		if Pow(2, attempts[i][0]) != attempts[i][1] {
			t.Errorf("Pow(2, %d) != %d", attempts[i][0], attempts[i][1])
		}
	}
}

func TestTreeHasBody(t *testing.T) {
	b := Body{
		r:    []float64{1.0},
		v:    []float64{1.0},
		f:    []float64{1.0},
		mass: .5,
	}

	tree := Tree{B: &b}

	if !tree.HasBody() {
		t.Errorf("tree does not have a body.")
	}
}

func TestIsInternal(t *testing.T) {
	b := Body{
		r:    []float64{1.0},
		v:    []float64{1.0},
		f:    []float64{1.0},
		mass: .5,
	}

	// A tree is internal if at least one subtree has a body
	tree := Tree{B: &b}
	tree2 := Tree{B: &b}
	tree.Ts = []Tree{tree2}

	if !tree.IsInternal() {
		t.Errorf("Tree (tree) is not internal, but its tree's have bodies.")
	}

}
