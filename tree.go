package hrtree

var capacity = 3

// RegularRTree is the most basic kind of RTrees
type RegularRTree struct {
	Root     *node
	capacity int
}

type searchres struct {
	results []interface{}
}

// NewRegularRTree returns new tree with capacity 3
func NewRegularRTree() *RegularRTree {
	return &RegularRTree{&node{}, 3}
}

// WithCapacity sets the RegularRTree capacity
func (r *RegularRTree) WithCapacity(capacity int) *RegularRTree {
	r.capacity = capacity
	return r
}

// Insert a new entry in the RTree, the entry has a surface (Rect) and
// a value
func (r *RegularRTree) Insert(re Rect, v interface{}) {
	n := &node{re, nil, nil, v}
	leaf := pickLeaf(r, n)
	var splitResult *node
	if len(leaf.children) == r.capacity {
		splitResult = leaf.split(n)
	} else {
		leaf.addChild(n)
	}
	rootSplitted := r.adjust(leaf, splitResult)
	if rootSplitted != nil {
		oldRoot := r.Root
		r.Root = &node{}
		r.Root.addChild(oldRoot)
		r.Root.addChild(rootSplitted)
	}
}

func (n *node) searchEntries(re Rect, results *searchres) {
	if n.children == nil && n.r.intersect(re) {
		results.results = append(results.results, n.value)
		return
	}
	for _, c := range n.children {
		if c.r.intersect(re) {
			c.searchEntries(re, results)
		}
	}

}

// Search for an entry in the RTree
func (r *RegularRTree) Search(re Rect) []interface{} {
	results := make([]interface{}, 0, 0)
	res := searchres{results}
	r.Root.searchEntries(re, &res)
	return res.results
}

func pickLeaf(r *RegularRTree, n *node) *node {
	current := r.Root
	for !isLeaf(current) {
		current = pickClosestChild(current, n)
	}
	return current
}

func (r *RegularRTree) pickHilbertLeaf(n *node) *node {
}

func (r *RegularRTree) adjust(leaf *node, splitResult *node) *node {
	if leaf == r.Root {
		return splitResult
	}
	leaf.parent.recomputeRect()
	if splitResult == nil {
		return r.adjust(leaf.parent, nil)
	}
	if len(leaf.children) == r.capacity {
		toSplit := leaf.parent
		splitted := toSplit.split(splitResult)
		return r.adjust(toSplit, splitted)
	}
	leaf.parent.addChild(splitResult)
	return r.adjust(leaf.parent, nil)
}
