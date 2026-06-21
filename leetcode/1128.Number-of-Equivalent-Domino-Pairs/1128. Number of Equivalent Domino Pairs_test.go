package leetcode

import (
	"fmt"
	"testing"
)

type question1128 struct {
	para1128
	ans1128
}

// para 是参数
// one 代表第一个参数
type para1128 struct {
	dominoes [][]int
}

// ans 是答案
// one 代表第一个答案
type ans1128 struct {
	one int
}

func Test_Problem1128(t *testing.T) {

	qs := []question1128{

		{
			para1128{[][]int{{1, 2}, {2, 1}, {3, 4}, {5, 6}}},
			ans1128{1},
		},
		{
			para1128{[][]int{{1, 2}, {1, 2}, {1, 1}, {1, 2}, {2, 2}}},
			ans1128{3},
		},
		{
			para1128{[][]int{}},
			ans1128{0},
		},
		{
			para1128{nil},
			ans1128{0},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1128------------------------\n")

	for _, q := range qs {
		a, p := q.ans1128, q.para1128
		got := numEquivDominoPairs(p.dominoes)
		if got != a.one {
			t.Fatalf("input: %v, expected: %v, got: %v", p.dominoes, a.one, got)
		}
		fmt.Printf("【input】:%v       【output】:%v\n", p, got)
	}
	fmt.Printf("\n\n\n")
}
