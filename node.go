// Copyright (C) 2016 Trent Hauck - All Rights Reserved

package barneshut

import (
	"encoding/json"
	"fmt"
	log "github.com/Sirupsen/logrus"
)

const (
	// theta is the benchmark value that determines if a set of bodies are sufficiently
	// far enough to treat as one body, or they need to be examined closer.
	theta = 0.5
)

// Node is a single position in space. It contains a quadrant describing that location and optionally
// a Body, B, and an array of children Nodes that are all within the the Node's Quadrant.
type Node struct {
	B     *Body    `json:"Body"`
	Q     Quadrant `json:"Quadrant"`
	Ns    []*Node  `json:"Nodes"`
	Level int      `json:"Level"` // the level of the node in the Tree
}

// Json returns the byte array of the json representation of a Node.
func (n Node) Json() []byte {
	b, _ := json.Marshal(n)
	return b
}

// String returns a string in pretty print form for the Node.
func (t Node) String() string {
	return fmt.Sprintf("Node{B: %s, Q: %s, Ns: %s, Level: %d}", t.B, t.Q, t.Ns, t.Level)
}

// HasBody returns a boolean that is true if the Body has a non-nil Body.
func (t Node) HasBody() bool {
	return t.B != nil
}

// IsInternal determines if a node is internal to the overall Tree.  A Node
// is internal if at least one of its children has a body.
func (t Node) IsInternal() bool {
	for _, node := range t.Ns {
		if node.HasBody() {
			return true
		}
	}

	return false
}

// NodesFromQuadrants retruns a list of Nodes of a list of Quadrants.
func NodesFromQuadrants(qs []Quadrant, l int) []*Node {
	var newNodes []*Node

	for _, q := range qs {
		newNodes = append(newNodes, &Node{Q: q, Level: l})
	}

	return newNodes
}

func (t *Node) isInternalInsert(pB *Body) {
	// isInternalInsert inserts a body into a Node if the Node isInternal.
	for _, newNode := range t.Ns {
		if pB.InQuadrant(newNode.Q) {
			log.Infof("Inserting (internal) Body:%s into Quadrant: %s", pB.Label, newNode.Q)
			newNode.Insert(pB)
		}
	}
	t.B.AddBody(pB)
}

func (t *Node) isExternalInsert(pB *Body) {
	// isExternalInsert inserts a body into a Node if the Node isExternal.
	currentBody := &Body{
		R:     t.B.R,
		V:     t.B.V,
		F:     t.B.F,
		Mass:  t.B.Mass,
		Label: t.B.Label,
	}

	// First inserts the passed body into the array.
	for _, newNode := range t.Ns {
		log.Infof("For pB(%s), checking quadrant %s", pB.Label, newNode.Q)
		if pB.InQuadrant(newNode.Q) {
			log.Infof("Inserting (external) Body:%s into Q: %s", pB.Label, newNode.Q)
			newNode.Insert(pB)
			break
		}
	}

	// Then reinsert the body that was in the Node position in the first place.
	for _, newNode := range t.Ns {
		log.Infof("For currentBody(%s), checking quadrant %s", currentBody.Label, newNode.Q)
		if currentBody.InQuadrant(newNode.Q) {
			log.Infof("Inserting (external) Body:%s into Q: %s", currentBody.Label, newNode.Q)
			newNode.Insert(currentBody)
			break
		}
	}

	// Update the Node's body to the center of mass between the two bodies.
	t.B.AddBody(pB)
}

// Insert adds body pB into the calling Node. Depending on if the Node is internal or
// external the node is inserted differently.
func (t *Node) Insert(pB *Body) {
	log.Infof("Inserting body: %s", pB.Label)
	body := &Body{
		R:     pB.R,
		V:     pB.V,
		F:     pB.F,
		Mass:  pB.Mass,
		Label: pB.Label,
	}

	if !t.HasBody() {
		log.Infof("t doesn't have a body, running t.B = body (%s)", body.Label)
		t.B = body
		return
	}

	if t.IsInternal() {
		log.Infof("t is Internal, running internal insert body (%s)", body.Label)
		t.isInternalInsert(body)
	} else {
		log.Infof("t is External, running external insert, body (%s).", body.Label)
		if t.Ns == nil {
			new_quadrants := t.Q.Subdivide()
			t.Ns = NodesFromQuadrants(new_quadrants, t.Level+1)
		}

		t.isExternalInsert(body)
	}

}

// UseBody returns a bool if the body at the calling Node should be used in the force calculations,
// or if the Tree needs to be traversed further down.
func (n *Node) UseBody(b *Body) bool {
	return ((b.DistanceTo(*n.B) / n.Q.Width) > theta) && !n.IsInternal()
}

// UpdateForce updates the force on the called body, b, based on the Bodies in Node, n, and its children.
func (n *Node) UpdateForce(b *Body) {
	if n.UseBody(b) {
		// We use the body at node n, and don't traverse its children.
		b.AddForce(*n.B)
	} else {
		for i := range n.Ns {
			if (n.Ns[i].B == nil) || (b.Label == n.Ns[i].B.Label) {
				fmt.Printf("Skipping index %d for %s.\n", i, b.Label)
				continue
			}

			n.Ns[i].UpdateForce(b)
		}
	}
}

// NodeFromQuadrant returns a Node given a passed Quadrant.
func NodeFromQuadrant(q Quadrant) Node {
	return Node{
		Q: q,
	}
}

// Pow returns a**b where a and b are integers.
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
