package hungarianAlgorithm

// Returns the solution as an array `a` such that each row `i` is matched to
// column `a[i]`
func Solve(costs [][]int) ([]int, error) {
	// Validate the input
	if err := validate(costs); err != nil {
		return []int{}, err
	}

	n := len(costs)
	label := makeLabel(n, costs) // labels on the row and columns
	match := makeMatching(n)     // matching using tight edges

	label.initialize()
	match.initialize(label.isTight)

	// loop until the matching is perfect
	for p, r := match.isPerfect(); !p; p, r = match.isPerfect() {
		e := label.initializeSlacks(r) // initializes the min slacks to r
		t := makeTree(n, r, e)         // alternating tree rooted at r

		// loop until the matching is augmented
		for true {
			var j int // new column index in the tree

			// Extend the tree
			if b, k := t.extend(); b {
				j = k
			} else {
				u, v := t.indices()
				e := label.update(u, v)
				t.addTightEdges(e)
				_, j = t.extend()
			}

			if b, i := match.isMatched(j); b {
				// Add the edge (i, j) to the tree
				t.addEdge(i, j)
				e := label.updateSlacks(i)
				t.addTightEdges(e)
			} else {
				// Augment the matching
				path := t.pathToRoot(j)
				match.augment(path)
				break
			}
		}
	}

	return match.format(), nil
}
