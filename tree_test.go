package barneshut

import (
	log "github.com/Sirupsen/logrus"
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

func TestInsert(t *testing.T) {
	// Tree Doesn't Have a Body
	b := Body{
		r:    []float64{0.6, 0.6},
		v:    []float64{1.0, 0.5},
		f:    []float64{1.0, 0.5},
		mass: .5,
	}

	tree := &Tree{Q: Quadrant{width: 1.0, base: []float64{0.0, 0.0}}}

	tree.Insert(b)
	log.Infof("After insert, tree is: %v, it hasBody: %v", tree, tree.HasBody())

	if !b.Equals(*tree.B) {
		t.Errorf("Body, b, is not the Body in tree even though it was empty.")
	}

	if tree.IsInternal() {
		t.Errorf("Tree is not internal though we just inserted a body.")
	}

	b2 := Body{
		r:    []float64{0.1, 0.1},
		v:    []float64{1.0, 0.1},
		f:    []float64{1.0, 0.1},
		mass: .5,
	}

	log.Infof("Before insert 2, tree is: %v", tree)
	tree.Insert(b2)
	log.Infof("After insert 2, tree is: %v, it hasBody: %v",
		tree, tree.HasBody())

}
