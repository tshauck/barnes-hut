// Copyright (C) 2016 Trent Hauck - All Rights Reserved

package barneshut

import (
	"testing"
)

func TestIndex2Offset(t *testing.T) {

	indexes := []int{0, 1, 2, 3}
	ndims := 2
	expectedOffset := [][]int{
		[]int{0, 0},
		[]int{0, 1},
		[]int{1, 0},
		[]int{1, 1},
	}

	for i := range indexes {
		actual := Index2Offset(indexes[i], ndims)

		for idx := 0; idx < ndims; idx++ {
			if actual[idx] != expectedOffset[i][idx] {
				t.Errorf("Expected (%v) got (%v)", expectedOffset[i],
					Index2Offset(indexes[i], ndims))
			}

		}
	}

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

func TestUpdateForce(t *testing.T) {
	bodies := exampleBodies()
	tt := testTree()

	body := &Body{
		Label: bodies["A"].Label,
		V:     bodies["A"].V,
		F:     bodies["A"].F,
		R:     bodies["A"].R,
		Mass:  bodies["A"].Mass,
	}

	tt.Root.UpdateForce(body)
}

func TestLevel(t *testing.T) {

	bodies := exampleBodies()

	testNode := Node{
		Q:     Quadrant{Width: 6, LL: []float64{-3, -3}},
		Level: 1,
	}

	testNode.Insert(bodies["A"])
	testNode.Insert(bodies["B"])
	testNode.Insert(bodies["C"])
	testNode.Insert(bodies["D"])

	if testNode.Level != 1 {
		t.Errorf("testNode has level %d, but should have level %d", testNode.Level, 1)
	}

	if testNode.Ns[0].Level != 2 {
		t.Errorf("testNode has level %d, but should have level %d", testNode.Ns[0].Level, 2)
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
	node.Ns = []*Node{&node2}

	if !node.IsInternal() {
		t.Errorf("Node (node) is not internal, but its node's have bodies.")
	}

}

func TestInsert(t *testing.T) {

	bodies := exampleBodies()

	testNode := Node{
		Q: Quadrant{Width: 6, LL: []float64{-3, -3}},
	}

	testNode.Insert(bodies["A"])
	testNode.Insert(bodies["B"])
	testNode.Insert(bodies["C"])
	testNode.Insert(bodies["D"])

	expected := testTree()
	expectedNode := expected.Root

	if !expectedNode.B.EqualPosition(testNode.B) {
		t.Errorf("Root bodies not equal: %s != %s", expectedNode.B, testNode.B)
	}

	for i := range expectedNode.Ns {
		eN := expectedNode.Ns[i]
		tN := testNode.Ns[i]

		if eN.B == nil && tN.B == nil {
			continue
		}

		if (eN.B != nil && tN.B == nil) || (eN.B == nil && tN.B != nil) || !eN.B.EqualPosition(tN.B) {
			t.Errorf("Expected Body: %s, Actual Body: %s", eN.B, tN.B)
		}
	}

}
