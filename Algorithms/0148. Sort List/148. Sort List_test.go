package leetcode

import (
	"fmt"
	"testing"
)

type question148 struct {
	para148
	ans148
}

// para 是参数
// one 代表第一个参数
type para148 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans148 struct {
	one []int
}

func Test_Problem148(t *testing.T) {

	qs := []question148{

		question148{
			para148{[]int{1, 2, 3, 4, 5}},
			ans148{[]int{1, 2, 3, 4, 5}},
		},
		question148{
			para148{[]int{1, 1, 2, 5, 5, 4, 10, 0}},
			ans148{[]int{0, 1, 1, 2, 4, 5, 5}},
		},

		question148{
			para148{[]int{1}},
			ans148{[]int{1}},
		},

		question148{
			para148{[]int{}},
			ans148{[]int{}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 148------------------------\n")

	for _, q := range qs {
		_, p := q.ans148, q.para148
		fmt.Printf("【input】:%v       【output】:%v\n", p, L2s(sortList(S2l(p.one))))
	}
	fmt.Printf("\n\n\n")
}
