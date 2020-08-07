package leetcode

import (
	"fmt"
	"testing"
)

type question756 struct {
	para756
	ans756
}

// para 是参数
// one 代表第一个参数
type para756 struct {
	b string
	a []string
}

// ans 是答案
// one 代表第一个答案
type ans756 struct {
	one bool
}

func Test_Problem756(t *testing.T) {

	qs := []question756{

		question756{
			para756{"BCD", []string{"BCG", "CDE", "GEA", "FFF"}},
			ans756{true},
		},

		question756{
			para756{"AABA", []string{"AAA", "AAB", "ABA", "ABB", "BAC"}},
			ans756{false},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 756------------------------\n")

	for _, q := range qs {
		_, p := q.ans756, q.para756
		fmt.Printf("【input】:%v       【output】:%v\n", p, pyramidTransition(p.b, p.a))
	}
	fmt.Printf("\n\n\n")
}
