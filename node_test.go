package barneshut

import (
	"testing"
)

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

func TestNodeHasBody(t *testing.T) {
	b := Body{
		R:    []float64{1.0},
		V:    []float64{1.0},
		F:    []float64{1.0},
		Mass: .5,
	}

	node := Node{B: &b}

	if !node.HasBody() {
		t.Errorf("node does not have a body.")
	}
}

func TestIsInternal(t *testing.T) {
	b := Body{
		R:    []float64{1.0},
		V:    []float64{1.0},
		F:    []float64{1.0},
		Mass: .5,
	}

	// A node is internal if at least one subnode has a body
	node := Node{B: &b}
	node2 := Node{B: &b}
	node.Ts = []*Node{&node2}

	if !node.IsInternal() {
		t.Errorf("Node (node) is not internal, but its node's have bodies.")
	}

}

func TestIsExternalInsert(t *testing.T) {

	ab := &Body{
		R:     []float64{-2, -1},
		V:     []float64{0.0, 0.0},
		F:     []float64{0.0, 0.0},
		Mass:  1,
		Label: "AB",
	}
	b := &Body{
		R:     []float64{-1, 2},
		V:     []float64{0.0, 0.0},
		F:     []float64{0.0, 0.0},
		Mass:  1,
		Label: "B",
	}
	ab.AddBody(b)

	a := &Body{
		R:     []float64{-2, -1},
		V:     []float64{0.0, 0.0},
		F:     []float64{0.0, 0.0},
		Mass:  1,
		Label: "A",
	}

	testA := &Body{
		R:     []float64{-2, -1},
		V:     []float64{0.0, 0.0},
		F:     []float64{0.0, 0.0},
		Mass:  1,
		Label: "A",
	}

	q := Quadrant{Width: 6, LL: []float64{-3, -3}}
	subq := q.Subdivide()
	nodes := NodesFromQuadrants(subq)

	testNode := Node{
		Q: q,
	}
	testNode.Insert(testA)
	testNode.Insert(b)

	// TODO(trent): Replace these with explicitly inserted indexes
	// in the expected index location.
	for i := range nodes {
		if nodes[i].Q.ContainsBody(a) {
			nodes[i].Insert(a)
		}

		if nodes[i].Q.ContainsBody(b) {
			nodes[i].Insert(b)
		}
	}

	expectedNode := Node{
		B:  ab,
		Q:  q,
		Ts: nodes,
	}

	if !expectedNode.B.Equals(testNode.B) {
		t.Errorf("%s != %s", expectedNode.B, testNode.B)
	}

	for i := range expectedNode.Ts {
		eN := expectedNode.Ts[i]
		tN := testNode.Ts[i]

		if eN.B == nil && tN.B == nil {
			continue
		}

		if (eN.B != nil && tN.B == nil) || (eN.B == nil && tN.B != nil) || !eN.B.Equals(tN.B) {
			t.Errorf("For index (%d) expected body (%s) got (%s)", i, eN.B, tN.B)
		}
	}

}
