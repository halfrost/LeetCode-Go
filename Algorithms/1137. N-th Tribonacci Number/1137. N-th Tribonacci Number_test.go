package leetcode

import (
	"fmt"
	"testing"
)

type question1137 struct {
	para1137
	ans1137
}

// para 是参数
// one 代表第一个参数
type para1137 struct {
	one int
}

// ans 是答案
// one 代表第一个答案
type ans1137 struct {
	one int
}

func Test_Problem1137(t *testing.T) {

	qs := []question1137{

		question1137{
			para1137{1},
			ans1137{1},
		},

		question1137{
			para1137{2},
			ans1137{1},
		},

		question1137{
			para1137{3},
			ans1137{2},
		},

		question1137{
			para1137{4},
			ans1137{4},
		},

		question1137{
			para1137{25},
			ans1137{1389537},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1137------------------------\n")

	for _, q := range qs {
		_, p := q.ans1137, q.para1137
		fmt.Printf("【input】:%v       【output】:%v\n", p, tribonacci(p.one))
	}
	fmt.Printf("\n\n\n")
}
