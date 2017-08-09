package hungarianAlgorithm

type matching struct {
	n  int
	ij []int
	ji []int
}

// Returns an empty matching of the given size
func makeMatching(n int) matching {
	ij := make([]int, n)
	ji := make([]int, n)
	for i := 0; i < n; i++ {
		ij[i] = -1
		ji[i] = -1
	}
	return matching{n, ij, ji}
}

// Greedily build a matching
func (m *matching) initialize(isTight func(int, int) bool) {
	for i := 0; i < m.n; i++ {
		for j := 0; j < m.n; j++ {
			if isTight(i, j) && (m.ji[j] == -1) {
				m.ij[i] = j
				m.ji[j] = i
				break
			}
		}
	}
}

// Returns whether the matching is perfect and a free vertex (when the matching is not perfect)
func (m *matching) isPerfect() (bool, int) {
	for i := 0; i < m.n; i++ {
		if m.ij[i] == -1 {
			return false, i
		}
	}
	return true, -1
}

// Returns whether the vertex of the right set is matched, when true is also
// returns the match
func (m *matching) isMatched(j int) (bool, int) {
	i := m.ji[j]
	return (i != -1), i
}

// Returns an array `a` representing the edges of the matching in the form `(i, a[i])`.
func (m *matching) format() []int {
	return m.ij
}

// Augments the matching using the given augmenting path.
func (m *matching) augment(p []int) {
	for idx, j := range p {
		if idx%2 == 0 {
			i := p[idx+1]
			m.ij[i] = j
			m.ji[j] = i
		}
	}
}
