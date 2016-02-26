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
		R:    []float64{-2, -1},
		V:    []float64{0.0, 0.0},
		F:    []float64{0.0, 0.0},
		Mass: 1,
	}
	b := &Body{
		R:    []float64{-1, 2},
		V:    []float64{0.0, 0.0},
		F:    []float64{0.0, 0.0},
		Mass: 1,
	}
	ab.AddBody(b)
	log.Infof("ab: %s", ab)

	a := &Body{
		R:    []float64{-2, -1},
		V:    []float64{0.0, 0.0},
		F:    []float64{0.0, 0.0},
		Mass: 1,
	}

	q := Quadrant{Width: 6, LL: []float64{-3, -3}}
	subq := q.Subdivide()
	nodes := NodesFromQuadrants(subq)

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

	log.Infof("%s", string(expectedNode.Json()))

	return

	testNode := Node{
		B: a,
		Q: q,
	}
	testNode.Insert(b)

	if !expectedNode.B.Equals(testNode.B) {
		t.Errorf("%s != %s", expectedNode.B, testNode.B)
	}

	for i := range expectedNode.Ts {
		eN := expectedNode.Ts[i]
		tN := testNode.Ts[i]

		log.Infof("For index (%d) expected body (%s) got (%s)", i, eN.B, tN.B)
		//if !eN.B.Equals(tN.B) {
		//t.Errorf("For index (%d) expected body (%s) got (%s)", i, eN.B, tN.B)
		//}
	}
}

//func TestInsert(t *testing.T) {
//// Node Doesn't Have a Body
//b := &Body{
//r:    []float64{0.6, 0.6},
//v:    []float64{1.0, 0.5},
//f:    []float64{1.0, 0.5},
//mass: .5,
//}

//node := &Node{Q: Quadrant{width: 1.0, base: []float64{0.0, 0.0}}}

//node.Insert(b)
//log.Infof("After insert, node is: %v, it hasBody: %v", node, node.HasBody())

//if !b.Equals(*node.B) {
//t.Errorf("Body, b, is not the Body in node even though it was empty.")
//}

//if node.IsInternal() {
//t.Errorf("Node is not internal though we just inserted a body.")
//}

//b2 := Body{
//r:    []float64{0.1, 0.1},
//v:    []float64{1.0, 0.1},
//f:    []float64{1.0, 0.1},
//mass: .5,
//}

//log.Infof("Before insert 2, node is: %v", node)
//node.Insert(&b2)
//log.Infof("After insert 2, node is: %v, it hasBody: %v",
//node, node.HasBody())

//}
