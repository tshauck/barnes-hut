// Copyright (C) 2016 Trent Hauck - All Rights Reserved

package barneshut

import (
	"testing"
)

// testTree returns a manually created tree for verifiable

func exampleBodiesNoVelocity() map[string]*Body {
	return map[string]*Body{
		"A":    &Body{R: []float64{-2, -1}, F: []float64{0, 0}, V: []float64{0, 0}, Mass: 1, Label: "A"},
		"B":    &Body{R: []float64{-1, 2}, F: []float64{0, 0}, V: []float64{0, 0}, Mass: 100, Label: "B"},
		"C":    &Body{R: []float64{1, -1}, F: []float64{0, 0}, V: []float64{0, 0}, Mass: 0.1, Label: "C"},
		"D":    &Body{R: []float64{2, -2}, F: []float64{0, 0}, V: []float64{0, 0}, Mass: 0.5, Label: "D"},
		"ABCD": &Body{R: []float64{-2.9 / 3.6, 1.9 / 3.6}, Mass: 3.6, Label: "ABCD"},
		"CD":   &Body{R: []float64{1.83333, -1.83333}, Mass: .6, Label: "CD"},
	}
}

func exampleBodies() map[string]*Body {
	return map[string]*Body{
		"A":    &Body{R: []float64{-2, -1}, F: []float64{0, 0}, V: []float64{0, 0}, Mass: 1, Label: "A"},
		"B":    &Body{R: []float64{-1, 2}, F: []float64{0, 0}, V: []float64{0, 1}, Mass: 2, Label: "B"},
		"C":    &Body{R: []float64{1, -1}, F: []float64{0, 0}, V: []float64{1, 0}, Mass: 0.1, Label: "C"},
		"D":    &Body{R: []float64{2, -2}, F: []float64{0, 0}, V: []float64{1, 1}, Mass: 0.5, Label: "D"},
		"ABCD": &Body{R: []float64{-2.9 / 3.6, 1.9 / 3.6}, Mass: 3.6, Label: "ABCD"},
		"CD":   &Body{R: []float64{1.83333, -1.83333}, Mass: .6, Label: "CD"},
	}
}

func insertTestBodies(bodies map[string]*Body) *Tree {
	q := Quadrant{Width: 6, LL: []float64{-3, -3}}
	ns := NodesFromQuadrants(q.Subdivide(), 1)
	ns[0].B = bodies["A"]
	ns[1].B = bodies["B"]
	ns[2].B = bodies["CD"]

	CDNs := NodesFromQuadrants(ns[2].Q.Subdivide(), 2)
	CDNs[1].B = bodies["C"]
	CDNs[2].B = bodies["D"]

	ns[2].Ns = CDNs

	RootNode := &Node{B: bodies["ABCD"], Q: q, Ns: ns}

	return &Tree{Root: RootNode}

}

func testTree(novelocity bool) *Tree {

	if novelocity {
		bodies := exampleBodiesNoVelocity()
		return insertTestBodies(bodies)
	} else {
		bodies := exampleBodies()
		return insertTestBodies(bodies)
	}

}

func TestSave(t *testing.T) {

	tree := testTree(false)
	err := tree.Save("tree.json")
	if err != nil {
		t.Error(err)
	}

	// TODO(trent): Need to remove files during test.
	_, err = TreeFromJsonFile("./tree.json")
	if err != nil {
		t.Error(err)
	}

}

func TestTreeFromBody(t *testing.T) {
	tree, err := TreeFromBodyFile("./bodies-big.json")
	if err != nil {
		panic(err)
	}

	tree.Save("./tree-big.json")
}
