package leetcode

import (
	"fmt"
	"testing"
)

type question1688 struct {
	para1688
	ans1688
}

// para 是参数
// one 代表第一个参数
type para1688 struct {
	n int
}

// ans 是答案
// one 代表第一个答案
type ans1688 struct {
	one int
}

func Test_Problem1688(t *testing.T) {

	qs := []question1688{

		{
			para1688{7},
			ans1688{6},
		},

		{
			para1688{14},
			ans1688{13},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1688------------------------\n")

	for _, q := range qs {
		_, p := q.ans1688, q.para1688
		fmt.Printf("【input】:%v       【output】:%v\n", p, numberOfMatches(p.n))
	}
	fmt.Printf("\n\n\n")
}
