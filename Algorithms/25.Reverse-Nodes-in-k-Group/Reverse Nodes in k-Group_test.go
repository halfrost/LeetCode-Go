package leetcode

import (
	"fmt"
	"testing"
)

type question25 struct {
	para25
	ans25
}

// para 是参数
// one 代表第一个参数
type para25 struct {
	one []int
	two int
}

// ans 是答案
// one 代表第一个答案
type ans25 struct {
	one []int
}

func Test_Problem25(t *testing.T) {

	qs := []question25{

		question25{
			para25{
				[]int{1, 2, 3, 4, 5},
				3,
			},
			ans25{[]int{3, 2, 1, 4, 5}},
		},

		question25{
			para25{
				[]int{1, 2, 3, 4, 5},
				1,
			},
			ans25{[]int{1, 2, 3, 4, 5}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 25------------------------\n")

	for _, q := range qs {
		_, p := q.ans25, q.para25
		fmt.Printf("【input】:%v       【output】:%v\n", p, L2s(reverseKGroup(S2l(p.one), p.two)))
	}
	fmt.Printf("\n\n\n")
}
