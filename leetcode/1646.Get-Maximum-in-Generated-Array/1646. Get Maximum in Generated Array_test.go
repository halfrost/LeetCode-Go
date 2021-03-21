package leetcode

import (
	"fmt"
	"testing"
)

type question1646 struct {
	para1646
	ans1646
}

// para 是参数
// one 代表第一个参数
type para1646 struct {
	n int
}

// ans 是答案
// one 代表第一个答案
type ans1646 struct {
	one int
}

func Test_Problem1646(t *testing.T) {

	qs := []question1646{

		{
			para1646{7},
			ans1646{3},
		},

		{
			para1646{2},
			ans1646{1},
		},

		{
			para1646{3},
			ans1646{2},
		},

		{
			para1646{0},
			ans1646{0},
		},

		{
			para1646{1},
			ans1646{1},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1646------------------------\n")

	for _, q := range qs {
		_, p := q.ans1646, q.para1646
		fmt.Printf("【input】:%v      【output】:%v      \n", p, getMaximumGenerated(p.n))
	}
	fmt.Printf("\n\n\n")
}
