package leetcode

import (
	"fmt"
	"testing"
)

type question815 struct {
	para815
	ans815
}

// para 是参数
// one 代表第一个参数
type para815 struct {
	r [][]int
	s int
	t int
}

// ans 是答案
// one 代表第一个答案
type ans815 struct {
	one int
}

func Test_Problem815(t *testing.T) {

	qs := []question815{

		{
			para815{[][]int{{1, 2, 7}, {3, 6, 7}}, 1, 6},
			ans815{2},
		},

		{
			para815{[][]int{{1, 2, 7}, {3, 6, 7}}, 5, 5},
			ans815{0},
		},

		{
			para815{[][]int{{1, 2, 7}, {3, 6, 7}}, 1, 100},
			ans815{-1},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 815------------------------\n")

	for _, q := range qs {
		a, p := q.ans815, q.para815
		fmt.Printf("【input】:%v       【output】:%v\n", p, numBusesToDestination(p.r, p.s, p.t))
		if got := numBusesToDestination(p.r, p.s, p.t); got != a.one {
			t.Fatalf("input: %v, expected: %v, got: %v", p, a.one, got)
		}
	}
	fmt.Printf("\n\n\n")
}
