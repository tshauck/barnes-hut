package barneshut

import (
	"encoding/json"
	"fmt"
	log "github.com/Sirupsen/logrus"
)

const (
	theta = 0.5
)

type Node struct {
	B  *Body    `json:"Body"`
	Q  Quadrant `json:"Quadrant"`
	Ts []*Node  `json:"Nodes"`
}

func (n Node) Json() []byte {
	b, _ := json.MarshalIndent(n, "", "  ")
	//b, _ := json.Marshal(n)
	return b
}

func (t Node) String() string {
	return fmt.Sprintf("Node{B: %s, Q: %s, Ts: %s}", t.B, t.Q, t.Ts)
}

func (t Node) HasBody() bool {
	return t.B != nil
}

func (t Node) IsInternal() bool {
	// A node is internal if there is at least one sub node
	// that has a body.

	for _, node := range t.Ts {
		if node.HasBody() {
			return true
		}
	}

	return false
}

func NodesFromQuadrants(qs []Quadrant) []*Node {
	var newNodes []*Node

	for _, q := range qs {
		newNodes = append(newNodes, &Node{Q: q})
	}

	return newNodes
}

func (t *Node) isInternalInsert(pB *Body) {
	for _, newNode := range t.Ts {
		if pB.InQuadrant(newNode.Q) {
			newNode.Insert(pB)
		}
	}
	t.B.AddBody(pB)
}

func (t *Node) isExternalInsert(pB *Body) {
	log.Debugf("Entering isExternalInsert.")
	currentBody := &Body{
		R:     t.B.R,
		V:     t.B.V,
		F:     t.B.F,
		Mass:  t.B.Mass,
		Label: t.B.Label,
	}

	for _, newNode := range t.Ts {
		if pB.InQuadrant(newNode.Q) {
			newNode.Insert(pB)
			break
		}
	}

	for _, newNode := range t.Ts {
		if currentBody.InQuadrant(newNode.Q) {
			newNode.Insert(currentBody)
			break
		}
	}

	t.B.AddBody(pB)
}

func (t *Node) Insert(pB *Body) {
	body := &Body{
		R:     pB.R,
		V:     pB.V,
		F:     pB.F,
		Mass:  pB.Mass,
		Label: pB.Label,
	}

	if !t.HasBody() {
		t.B = body
		log.Infof("Node lacks body: inserted Body: %s into Node: %s",
			t.B.Label, t)
		return
	}

	if t.IsInternal() {
		log.Infof("Internal Node, inserting: %s", body)
		t.isInternalInsert(body)
	} else {
		log.Infof("External Node, inserting: %s", body)
		new_quadrants := t.Q.Subdivide()
		t.Ts = NodesFromQuadrants(new_quadrants)

		t.isExternalInsert(body)
	}

}

func (n *Node) UseBody(b *Body) bool {
	// Given that we're on Node, n, do we use the body in the Node or
	// use traverse down the Tree.

	// Also, if the Node is internal there are no possible children to use
	// for a more precise center of mass, therefore, use the body.

	return ((b.DistanceTo(*n.B) / n.Q.Width) > theta) && !n.IsInternal()
}

func (n *Node) UpdateForce(b *Body) {
	if n.UseBody(b) {
		// We use the body at node n, and don't traverse its children.
		b.AddForce(*n.B)
	} else {
		// Body is too close, use its children's bodies.
		for i := range n.Ts {
			// If it's the same point, don't UpdateForece
			if b.Label == n.B.Label {
				continue
			}
			n.Ts[i].UpdateForce(b)
		}
	}
}

func NodeFromQuadrant(q Quadrant) Node {
	return Node{
		Q: q,
	}
}

func Pow(a, b int) int {
	p := 1
	for b > 0 {
		if b&1 != 0 {
			p *= a
		}
		b >>= 1
		a *= a
	}
	return p
}

// TODO(trent): rename
// Offset2Index takes an array of integers that represent an index in N-D
// space and converts it to an integer offset that works in 1-D space.
func Offset2Index(offset []int) int {

	n_dims := len(offset)
	index := 0
	for i := range offset {
		index += Pow(2, i) * offset[n_dims-i-1]
	}

	return index

}

// TODO(trent): rename
// Index2Offset converts an index in 1-D space to N-D Space.
func Index2Offset(index int, n_dims int) []int {

	var offset []int
	for n_dim := 0; n_dim < n_dims; n_dim++ {
		shift := uint64(n_dims - n_dim - 1)
		unsigned_shift := shift

		rem := ((index >> unsigned_shift) << unsigned_shift)

		if rem > 0 {
			offset = append(offset, 1)
		} else {
			offset = append(offset, 0)
		}

		index -= rem
	}

	return offset
}
