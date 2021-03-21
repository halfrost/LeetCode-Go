package leetcode

import (
	"fmt"
	"testing"
)

type question1614 struct {
	para1614
	ans1614
}

// para 是参数
// one 代表第一个参数
type para1614 struct {
	p string
}

// ans 是答案
// one 代表第一个答案
type ans1614 struct {
	one int
}

func Test_Problem1614(t *testing.T) {

	qs := []question1614{

		{
			para1614{"(1+(2*3)+((8)/4))+1"},
			ans1614{3},
		},

		{
			para1614{"(1)+((2))+(((3)))"},
			ans1614{3},
		},

		{
			para1614{"1+(2*3)/(2-1)"},
			ans1614{1},
		},

		{
			para1614{"1"},
			ans1614{0},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1614------------------------\n")

	for _, q := range qs {
		_, p := q.ans1614, q.para1614
		fmt.Printf("【input】:%v      【output】:%v      \n", p, maxDepth(p.p))
	}
	fmt.Printf("\n\n\n")
}
