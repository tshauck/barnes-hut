package barneshut

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Tree struct {
	Root *Node
}

func (t *Tree) Insert(b *Body) {
	t.Root.Insert(b)
}

func (t Tree) Save(f string) error {
	b, err := json.Marshal(t)

	fmt.Println(string(b))
	if err != nil {
		return err
	}

	return ioutil.WriteFile(f, b, 0644)
}

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
