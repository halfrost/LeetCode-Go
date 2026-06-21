package leetcode

import (
	"fmt"
	"testing"
)

type question60 struct {
	para60
	ans60
}

// para 是参数
// one 代表第一个参数
type para60 struct {
	n int
	k int
}

// ans 是答案
// one 代表第一个答案
type ans60 struct {
	one string
}

func Test_Problem60(t *testing.T) {

	qs := []question60{

		{
			para60{3, 3},
			ans60{"213"},
		},

		{
			para60{4, 9},
			ans60{"2314"},
		},

		{
			para60{3, 0},
			ans60{""},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 60------------------------\n")

	for _, q := range qs {
		a, p := q.ans60, q.para60
		got := getPermutation(p.n, p.k)
		if got != a.one {
			t.Fatalf("input %v expected %v got %v", p, a.one, got)
		}
		fmt.Printf("【input】:%v       【output】:%v\n", p, got)
	}
	fmt.Printf("\n\n\n")
}
