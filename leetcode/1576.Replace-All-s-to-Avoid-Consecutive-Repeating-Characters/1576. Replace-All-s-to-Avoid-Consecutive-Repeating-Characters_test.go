package leetcode

import (
	"fmt"
	"testing"
)

type question1576 struct {
	para1576
	ans1576
}

// para 是参数
// one 代表第一个参数
type para1576 struct {
	s string
}

// ans 是答案
// one 代表第一个答案
type ans1576 struct {
	one string
}

func Test_Problem1576(t *testing.T) {

	qs := []question1576{

		{
			para1576{"?zs"},
			ans1576{"azs"},
		},

		{
			para1576{"ubv?w"},
			ans1576{"ubvaw"},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1576------------------------\n")

	for _, q := range qs {
		_, p := q.ans1576, q.para1576
		fmt.Printf("【input】:%v      【output】:%v      \n", p, modifyString(p.s))
	}
	fmt.Printf("\n\n\n")
}
