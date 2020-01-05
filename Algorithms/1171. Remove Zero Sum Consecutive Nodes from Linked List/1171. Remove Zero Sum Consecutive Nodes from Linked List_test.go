package leetcode

import (
	"fmt"
	"testing"
)

type question1171 struct {
	para1171
	ans1171
}

// para 是参数
// one 代表第一个参数
type para1171 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans1171 struct {
	one []int
}

func Test_Problem1171(t *testing.T) {

	qs := []question1171{

		question1171{
			para1171{[]int{1, 2, -3, 3, 1}},
			ans1171{[]int{3, 1}},
		},

		question1171{
			para1171{[]int{1, 2, 3, -3, 4}},
			ans1171{[]int{1, 2, 4}},
		},

		question1171{
			para1171{[]int{1, 2, 3, -3, -2}},
			ans1171{[]int{1}},
		},

		question1171{
			para1171{[]int{1, -1}},
			ans1171{[]int{}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1171------------------------\n")

	for _, q := range qs {
		_, p := q.ans1171, q.para1171
		fmt.Printf("【input】:%v       【output】:%v\n", p, L2s(removeZeroSumSublists(S2l(p.one))))
	}
	fmt.Printf("\n\n\n")
}
