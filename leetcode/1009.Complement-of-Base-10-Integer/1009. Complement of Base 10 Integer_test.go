package leetcode

import (
	"fmt"
	"testing"
)

type question1009 struct {
	para1009
	ans1009
}

// para 是参数
// one 代表第一个参数
type para1009 struct {
	n int
}

// ans 是答案
// one 代表第一个答案
type ans1009 struct {
	one int
}

func Test_Problem1009(t *testing.T) {

	qs := []question1009{

		{
			para1009{5},
			ans1009{2},
		},

		{
			para1009{7},
			ans1009{0},
		},

		{
			para1009{10},
			ans1009{5},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1009------------------------\n")

	for _, q := range qs {
		_, p := q.ans1009, q.para1009
		fmt.Printf("【input】:%v       【output】:%v\n", p, bitwiseComplement(p.n))
	}
	fmt.Printf("\n\n\n")
}
