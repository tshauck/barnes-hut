package barneshut

import (
	"encoding/json"
	"fmt"
	log "github.com/Sirupsen/logrus"
	//"reflect"
)

type Node struct {
	B  *Body    `json:"Body"`
	Q  Quadrant `json:"Quadrant"`
	Ts []*Node  `json:"Nodes"`
}

func (n Node) Json() []byte {
	b, _ := json.MarshalIndent(n, "", "  ")
	return b
}

func (t Node) String() string {
	return fmt.Sprintf("Node{B: %s, Q: %s, Ts: %s}", t.B, t.Q, t.Ts)
}

func (t Node) HasBody() bool {
	//log.Infof("t.B is %v", t.B)
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

}

func (t *Node) isExternalInsert(pB *Body) {
	currentBody := t.B // Kinda need to make a copy here?

	// TODO(trent): Update AddBody to modify the body
	log.Infof("%v", t.B)
	t.B.AddBody(pB)
	log.Infof("%v", t.B)

	for _, newNode := range t.Ts {
		if pB.InQuadrant(newNode.Q) {
			log.Infof("Putting new body in node: %v", newNode)
			newNode.Insert(pB)
			break
		}
	}

	for _, newNode := range t.Ts {
		if currentBody.InQuadrant(newNode.Q) {
			log.Infof("Putting old body in node: %v", newNode)
			newNode.Insert(currentBody)
		}
	}
}

func (t *Node) Insert(pB *Body) {
	// Note: Node is passed as a pointer here so that it can be modified
	if !t.HasBody() {
		t.B = pB
		return
	}

	if t.IsInternal() {
		t.isInternalInsert(pB)
	} else {
		new_quadrants := t.Q.Subdivide()
		t.Ts = NodesFromQuadrants(new_quadrants)

		t.isExternalInsert(pB)
	}

}

func (t Node) UpdateForce(b Body) {
	// logic to update the forces given all
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
