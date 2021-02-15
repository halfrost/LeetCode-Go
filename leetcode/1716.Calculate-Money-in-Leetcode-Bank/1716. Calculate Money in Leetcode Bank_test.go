package leetcode

import (
	"fmt"
	"testing"
)

type question1716 struct {
	para1716
	ans1716
}

// para 是参数
// one 代表第一个参数
type para1716 struct {
	n int
}

// ans 是答案
// one 代表第一个答案
type ans1716 struct {
	one int
}

func Test_Problem1716(t *testing.T) {

	qs := []question1716{

		{
			para1716{4},
			ans1716{10},
		},

		{
			para1716{10},
			ans1716{37},
		},

		{
			para1716{20},
			ans1716{96},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1716------------------------\n")

	for _, q := range qs {
		_, p := q.ans1716, q.para1716
		fmt.Printf("【input】:%v       【output】:%v\n", p, totalMoney(p.n))
	}
	fmt.Printf("\n\n\n")
}
