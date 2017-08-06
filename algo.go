package main

import (
	"fmt"
)

// Returns the solution as an array `a` such that each row `i` is matched to
// column `a[i]`
func algo(costs [][]int) ([]int, error) {
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

func test(costs [][]int) {
	r, err := algo(costs)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("everything went fine")
		fmt.Println(r)
	}
}

func main() {
	//test([][]int{ {11, 6, 12}, {12, 4, 6}, {8, 12, 11}, })
	test([][]int{
		{13, 13, 19, 50, 33, 38},
		{73, 33, 71, 77, 97, 95},
		{20, 8, 56, 55, 64, 35},
		{26, 25, 72, 32, 55, 77},
		{83, 40, 69, 3, 53, 49},
		{67, 20, 44, 29, 86, 61},
	})
}
