## Hungarian Algorithm

An implementation of the Hungarian Algorithm for solving the assignment problem.
The implementation follows those [notes](http://www.cse.ust.hk/~golin/COMP572/Notes/Matching.pdf).

### Usage

```Go
package main

import (
	"fmt"
	"github.com/oddg/hungarian-algorithm"
)

func main() {
	a := [][]int{{11, 6, 12}, {12, 4, 6}, {8, 12, 11}}
	fmt.Println(hungarianAlgorithm.Solve(a))
}
```
