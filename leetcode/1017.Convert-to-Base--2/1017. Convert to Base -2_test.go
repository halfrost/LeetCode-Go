package leetcode

import (
	"fmt"
	"testing"
)

type question1017 struct {
	para1017
	ans1017
}

// para 是参数
// one 代表第一个参数
type para1017 struct {
	one int
}

// ans 是答案
// one 代表第一个答案
type ans1017 struct {
	one string
}

func Test_Problem1017(t *testing.T) {

	qs := []question1017{

		question1017{
			para1017{2},
			ans1017{"110"},
		},

		question1017{
			para1017{3},
			ans1017{"111"},
		},

		question1017{
			para1017{4},
			ans1017{"110"},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1017------------------------\n")

	for _, q := range qs {
		_, p := q.ans1017, q.para1017
		fmt.Printf("【input】:%v       【output】:%v\n", p, baseNeg2(p.one))
	}
	fmt.Printf("\n\n\n")
}
