package barneshut

import (
	log "github.com/Sirupsen/logrus"
	//"reflect"
)

type Tree struct {
	B           *Body
	Q           Quadrant
	Ts          []Tree
	HasChildren bool
}

func (t Tree) HasBody() bool {
	return t.B != nil
}

func (t Tree) IsInternal() bool {
	for _, tree := range t.Ts {
		if tree.B == nil {
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

func (t Tree) Insert(b Body) {
	// TODO(refactor)

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
		for _, newTree := range t.Ts {
			if b.InQuadrant(newTree.Q) {
				newTree.Insert(b)
			}
		}

		return
	} else {
		currentBody := t.B
		*t.B = t.B.AddBody(b)

		for _, newTree := range t.Ts {
			if b.InQuadrant(newTree.Q) {
				newTree.Insert(b)
			}
		}

		for _, newTree := range t.Ts {
			if b.InQuadrant(newTree.Q) {
				newTree.Insert(*currentBody)
			}
		}

	}

	return

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

	log.Infof("Returned Index2Offset, got %v", offset)

	return offset
}
