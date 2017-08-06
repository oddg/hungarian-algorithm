package main

import (
	"errors"
	"fmt"
)

func validate(costs [][]int) error {
	n := len(costs)

	if n == 0 {
		return errors.New("The costs matrix is empty.")
	}

	if m := len(costs[0]); m != n {
		return errors.New("The costs matrix is not square.")
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if costs[i][j] < 0 {
				return fmt.Errorf("The coefficient (%d,%d) is negative.", i, j)
			}
		}
	}

	return nil
}

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
		fmt.Println("matching:", match)
		fmt.Println("root:", r)

		edges := label.initializeSlacks(r) // initializes the min slacks to r
		tre := makeTree(n, r, edges)       // alternating tree rooted at r

		// loop until the matching is augmented
		for true {
			var j int // new column index in the tree

			fmt.Println("before tree:", tre)
			// Extend the tree
			if b, tmp := tre.extend(); b {
				fmt.Println("Tree extended with vertex:", tmp)
				j = tmp
			} else {
				s, t := tre.indices()
				fmt.Println("Tree not extended components:", s, t)
				e := label.update(s, t)
				tre.addTightEdges(e)
				fmt.Println("New labels:", label.left, label.right)
				_, j = tre.extend()
				fmt.Println("Tree extended with vertex:", j)
			}
			fmt.Println("after tree:", tre)

			if b, i := match.isMatched(j); b {
				fmt.Println("new vertex is matched with:", i)
				// Add the edge (i, j) to the tree
				tre.addEdge(i, j)
				e := label.updateSlacks(i)
				tre.addTightEdges(e)
			} else {
				fmt.Println("new vertex is not matched")
				// Augment the matching
				path := tre.pathToRoot(j)
				fmt.Println("path to root", path)
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
