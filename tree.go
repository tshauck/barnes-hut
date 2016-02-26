package barneshut

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	//"reflect"
)

type Tree struct {
	B           *Body
	Q           Quadrant
	Ts          []Tree
	HasChildren bool
}

func (t Tree) String() string {
	return fmt.Sprintf("Tree{B: %s, Q: %s, Ts: %s}", t.B, t.Q, t.Ts)
}

func (t Tree) HasBody() bool {
	//log.Infof("t.B is %v", t.B)
	return t.B != nil
}

func (t Tree) IsInternal() bool {
	// A tree is internal if there is at least one sub tree
	// that has a body.

	for _, tree := range t.Ts {
		if tree.HasBody() {
			return true
		}
	}

	return false
}

func TreesFromQuadrants(qs []Quadrant) []Tree {
	var newTrees []Tree

	for _, q := range qs {
		newTrees = append(newTrees, Tree{Q: q})
	}

	return newTrees
}

func (t *Tree) isInternalInsert(b Body) {
	for _, newTree := range t.Ts {
		if b.InQuadrant(newTree.Q) {
			newTree.Insert(b)
		}
	}

}

func (t *Tree) isExternalInsert(b Body) {
	currentBody := t.B

	// TODO(trent): Update AddBody to modify the body
	//t.B = t.B.AddBody(b)

	for _, newTree := range t.Ts {
		if b.InQuadrant(newTree.Q) {
			log.Infof("Putting new body in tree: %v", newTree)
			newTree.Insert(b)
			break
		}
	}

	for _, newTree := range t.Ts {
		if currentBody.InQuadrant(newTree.Q) {
			log.Infof("Putting old body in tree: %v", newTree)
			newTree.Insert(*currentBody)
		}
	}
}

func (t *Tree) Insert(b Body) {
	// Note: Tree is passed as a pointer here so that it can be modified

	if !t.HasChildren {
		new_quadrants := t.Q.Subdivide()
		t.Ts = TreesFromQuadrants(new_quadrants)
		t.HasChildren = true
	}

	if !t.HasBody() {
		t.B = &b
		return
	}

	if t.IsInternal() {
		log.Infof("t is internal, inserting.")
		t.isInternalInsert(b)
	} else {
		log.Infof("t is external, inserting.")
		t.isExternalInsert(b)
	}

}

func (t Tree) UpdateForce(b Body) {
	// logic to update the forces given all
}

func TreeFromQuadrant(q Quadrant) Tree {
	return Tree{
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
