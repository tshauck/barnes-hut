// Copyright (C) 2016 Trent Hauck - All Rights Reserved

package barneshut

import (
	"encoding/json"
	"io/ioutil"
)

// Tree is struct that contains a pointer to the root Node.
type Tree struct {
	Root *Node
}

// Insert calls the Root's insert method.
func (t *Tree) Insert(b *Body) {
	t.Root.Insert(b)
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
