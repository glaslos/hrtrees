package hrtree

type node struct {
	r        Rect
	parent   *node
	children []*node     // nil for entities
	value    interface{} // nil for internal node
}

func (n *node) addChild(child *node) {
	n.children = append(n.children, child)
	child.parent = n
	n.r = n.r.union(child.r)
}

func (n *node) recomputeRect() {
	if len(n.children) == 0 {
		return
	}
	n.r = n.children[0].r
	for _, c := range n.children[1:] {
		n.r = n.r.union(c.r)
	}
}

func (n *node) removeChild(childtoremove *node) {
	newchildren := make([]*node, 0, 0)
	for _, child := range n.children {
		if child != childtoremove {
			newchildren = append(newchildren, child)
		}
	}
	n.recomputeRect()
	n.children = newchildren
}

func (n *node) split(toadd *node) *node {
	child := pickClosestChild(n, toadd)
	n.removeChild(child)
	newChild := &node{}
	newChild.addChild(child)
	newChild.addChild(toadd)
	return newChild
}

func isLeaf(n *node) bool {
	return len(n.children) == 0 || n.children[0].children == nil
}

func pickClosestChild(parent *node, tosearch *node) *node {
	// Find the child that will lead to the minimum enlargment
	// Initialize with the worst case
	minEnlargment := parent.r.union(tosearch.r).area()
	var chosen *node
	// TODO (glaslos): Do we have to iterate over all children?
	for _, child := range parent.children {
		enlargment := child.r.union(tosearch.r).area()
		if enlargment <= minEnlargment {
			chosen = child
			minEnlargment = enlargment
		}
	}
	return chosen
}

func pickClosestHilbertChild(parent *node, tosearch *node) *node {
	var chosen *node
	for _, child := range parent.children {
		chosen = child
	}
	return chosen
}
