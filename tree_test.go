package barneshut

import (
	"fmt"
	"testing"
)

func TestSave(t *testing.T) {

	testA := &Body{
		R:     []float64{-2, -1},
		V:     []float64{0.0, 0.0},
		F:     []float64{0.0, 0.0},
		Mass:  1,
		Label: "A",
	}

	b := &Body{
		R:     []float64{-1, 2},
		V:     []float64{0.0, 0.0},
		F:     []float64{0.0, 0.0},
		Mass:  1,
		Label: "B",
	}

	q := Quadrant{Width: 6, LL: []float64{-3, -3}}

	testNode := Node{
		Q: q,
	}
	testNode.Insert(testA)
	testNode.Insert(b)

	tree := Tree{Root: &testNode}
	err := tree.Save("tree.json")
	if err != nil {
		t.Error(err)
	}

	// TODO(trent): Need to remove files during test.
	newTree, err := TreeFromJsonFile("tree.json")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(newTree)

}
