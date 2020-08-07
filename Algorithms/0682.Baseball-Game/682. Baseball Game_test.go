package leetcode

import (
	"fmt"
	"testing"
)

type question682 struct {
	para682
	ans682
}

// para 是参数
// one 代表第一个参数
type para682 struct {
	one []string
}

// ans 是答案
// one 代表第一个答案
type ans682 struct {
	one int
}

func Test_Problem682(t *testing.T) {

	qs := []question682{

		question682{
			para682{[]string{"5", "2", "C", "D", "+"}},
			ans682{30},
		},

		question682{
			para682{[]string{"5", "-2", "4", "C", "D", "9", "+", "+"}},
			ans682{27},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 682------------------------\n")

	for _, q := range qs {
		_, p := q.ans682, q.para682
		fmt.Printf("【input】:%v       【output】:%v\n", p, calPoints(p.one))
	}
	fmt.Printf("\n\n\n")
}
