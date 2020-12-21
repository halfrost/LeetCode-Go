package leetcode

import (
	"fmt"
	"testing"
)

type question910 struct {
	para910
	ans910
}

type para910 struct {
	A []int
	K int
}

type ans910 struct {
	one int
}

// Test_Problem910 ...
func Test_Problem910(t *testing.T) {

	qs := []question910{

		{
			para910{[]int{1}, 0},
			ans910{0},
		},
		{
			para910{[]int{0, 10}, 2},
			ans910{6},
		},
		{
			para910{[]int{1, 3, 6}, 3},
			ans910{3},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 910------------------------\n")

	for _, q := range qs {
		_, p := q.ans910, q.para910
		fmt.Printf("【input】:%v       【output】:%v\n", p, smallestRangeII(p.A, p.K))
	}
	fmt.Printf("\n\n\n")
}
