package leetcode

import (
	"fmt"
	"testing"
)

type question287 struct {
	para287
	ans287
}

// para 是参数
// one 代表第一个参数
type para287 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans287 struct {
	one int
}

func Test_Problem287(t *testing.T) {

	qs := []question287{

		question287{
			para287{[]int{1, 3, 4, 2, 2}},
			ans287{2},
		},

		question287{
			para287{[]int{3, 1, 3, 4, 2}},
			ans287{3},
		},

		question287{
			para287{[]int{}},
			ans287{0},
		},

		question287{
			para287{[]int{2, 2, 2, 2, 2}},
			ans287{2},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 287------------------------\n")

	for _, q := range qs {
		_, p := q.ans287, q.para287
		fmt.Printf("【input】:%v       【output】:%v\n", p, findDuplicate(p.one))
	}
	fmt.Printf("\n\n\n")
}
