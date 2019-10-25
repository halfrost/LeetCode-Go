package leetcode

import (
	"fmt"
	"testing"
)

type question1111 struct {
	para1111
	ans1111
}

// para 是参数
// one 代表第一个参数
type para1111 struct {
	one string
}

// ans 是答案
// one 代表第一个答案
type ans1111 struct {
	one []int
}

func Test_Problem1111(t *testing.T) {

	qs := []question1111{

		question1111{
			para1111{"(()())"},
			ans1111{[]int{0, 1, 1, 1, 1, 0}},
		},

		question1111{
			para1111{"()(())()"},
			ans1111{[]int{0, 0, 0, 1, 1, 0, 1, 1}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1111------------------------\n")

	for _, q := range qs {
		_, p := q.ans1111, q.para1111
		fmt.Printf("【input】:%v       【output】:%v\n", p, maxDepthAfterSplit(p.one))
	}
	fmt.Printf("\n\n\n")
}
