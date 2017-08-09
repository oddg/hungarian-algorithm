package hungarianAlgorithm

type edge struct {
	i, j int
}

type edgeSet struct {
	set []edge
}

func (s *edgeSet) pop() (bool, edge) {
	l := len(s.set)
	if l == 0 {
		return false, edge{}
	}
	e := s.set[l-1]
	s.set = s.set[:l-1]
	return true, e
}

func (s *edgeSet) add(e edge) {
	s.set = append(s.set, e)
}

type tree struct {
	n         int
	root      int
	leftPrec  []int
	rightPrec []int
	edges     edgeSet
}

func makeTree(n int, r int, e []edge) tree {
	left := make([]int, n)
	right := make([]int, n)
	for i := 0; i < n; i++ {
		left[i] = -1
		right[i] = -1
	}
	left[r] = r
	return tree{n, r, left, right, edgeSet{e}}
}

// Returns a couple of values:
// - a boolean: whether the extension has been successful
// - the index of the new element in the tree (when successful)
func (t *tree) extend() (bool, int) {
	b, e := t.edges.pop()
	if b {
		t.rightPrec[e.j] = e.i
		return true, e.j
	}
	return false, -1
}

// Returns two arrays of indices. The first contain the indices of the elements
// in the left side of the tree. The second constains the indices of the
// elements in the right side of the tree. All the indices are sorted in
// increasing order.
func (t *tree) indices() ([]int, []int) {
	l := make([]int, 0, t.n)
	r := make([]int, 0, t.n)

	for i := 0; i < t.n; i++ {
		if t.leftPrec[i] != -1 {
			l = append(l, i)
		}
		if t.rightPrec[i] != -1 {
			r = append(r, i)
		}
	}

	return l, r
}

// Add the edge (i, j) to the tree
func (t *tree) addEdge(i int, j int) {
	t.leftPrec[i] = j
}

// Return the path to the root starting from the given vertex.
func (t *tree) pathToRoot(end int) []int {
	path := make([]int, 0, 2*t.n)

	j := end
	i := t.rightPrec[j]

	for i != t.root {
		path = append(path, j, i)
		j = t.leftPrec[i]
		i = t.rightPrec[j]
	}

	path = append(path, j, i)

	return path
}

func (t *tree) addTightEdges(edges []edge) {
	for _, e := range edges {
		t.edges.add(e)
	}
}
