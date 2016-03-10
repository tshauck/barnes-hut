// Copyright (C) 2016 Trent Hauck - All Rights Reserved

package barneshut

import (
	"bufio"
	"encoding/json"
	"io/ioutil"
	"os"
)

// Tree is struct that contains a pointer to the root Node.
type Tree struct {
	Root *Node
}

// Insert calls the Root's insert method.
func (t *Tree) Insert(b *Body) {
	t.Root.Insert(b)
}

// TODO(trent): This is an odd function all alone.
func updateLocations(n *Node) {
	for _, ns := range n.Ns {
		if !ns.IsInternal() && ns.B != nil {
			ns.B.Update(1)
		} else {
			updateLocations(ns)
		}
	}
}

func (t *Tree) UpdateBodies() {
	updateLocations(t.Root)
}

// Save persists the json representation of a Tree into file, f.
func (t Tree) Save(f string) error {
	b, err := json.Marshal(t)

	if err != nil {
		return err
	}

	return ioutil.WriteFile(f, b, 0644)
}

// TreeFromJsonFile loads a json file into a Tree pointer.
func TreeFromJsonFile(f string) (*Tree, error) {
	opened_f, err := ioutil.ReadFile(f)
	if err != nil {
		return nil, err
	}

	var t Tree
	err = json.Unmarshal(opened_f, &t)
	if err != nil {
		return nil, err
	}

	return &t, nil

}

// TreeFromBodies takes an array of Bodies and returns
// a populated Tree.
func TreeFromBodies(bs []*Body) (*Tree, error) {

	// TODO(trent): need to infer q somehow, maybe this is
	// set as a config option
	q := Quadrant{Width: 6, LL: []float64{-3, -3}}

	t := &Tree{Root: &Node{Q: q}}

	for _, b := range bs {
		t.Insert(b)
	}

	return t, nil
}

// TreeFromBodyFile takes a file that contains JSON rows of Bodies
// and returns a populated Tree.
func TreeFromBodyFile(f string) (*Tree, error) {
	file, err := os.Open(f)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var bs []*Body
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		var b *Body
		err = json.Unmarshal(scanner.Bytes(), &b)
		if err != nil {
			return nil, err
		}

		bs = append(bs, b)
	}

	return TreeFromBodies(bs)
}
