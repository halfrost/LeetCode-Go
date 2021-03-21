package leetcode

import (
	"fmt"
	"testing"
)

type question1689 struct {
	para1689
	ans1689
}

// para 是参数
// one 代表第一个参数
type para1689 struct {
	n string
}

// ans 是答案
// one 代表第一个答案
type ans1689 struct {
	one int
}

func Test_Problem1689(t *testing.T) {

	qs := []question1689{

		{
			para1689{"32"},
			ans1689{3},
		},

		{
			para1689{"82734"},
			ans1689{8},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1689------------------------\n")

	for _, q := range qs {
		_, p := q.ans1689, q.para1689
		fmt.Printf("【input】:%v       【output】:%v\n", p, minPartitions(p.n))
	}
	fmt.Printf("\n\n\n")
}
