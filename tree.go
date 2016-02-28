package barneshut

type Tree struct {
	Root *Node
}

func (t *Tree) Insert(b *Body) {
	t.Root.Insert(b)
}
