package hungarianAlgorithm

import (
	"reflect"
	"testing"
)

type testCase struct {
	in   [][]int
	want []int
}

func TestSolve(t *testing.T) {
	cases := []testCase{
		{
			in: [][]int{
				{11, 6, 12},
				{12, 4, 6},
				{8, 12, 11},
			},
			want: []int{1, 2, 0},
		},
		{
			in: [][]int{
				{13, 13, 19, 50, 33, 38},
				{73, 33, 71, 77, 97, 95},
				{20, 8, 56, 55, 64, 35},
				{26, 25, 72, 32, 55, 77},
				{83, 40, 69, 3, 53, 49},
				{67, 20, 44, 29, 86, 61},
			},
			want: []int{4, 1, 5, 0, 3, 2},
		},
	}

	for _, c := range cases {
		got, err := Solve(c.in)
		if err != nil {
			t.Errorf(err.Error())
		}
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("Algo(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
