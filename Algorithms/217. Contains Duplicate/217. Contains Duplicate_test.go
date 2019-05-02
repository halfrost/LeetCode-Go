package leetcode

import (
	"fmt"
	"testing"
)

type question217 struct {
	para217
	ans217
}

// para 是参数
// one 代表第一个参数
type para217 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans217 struct {
	one bool
}

func Test_Problem217(t *testing.T) {

	qs := []question217{

		question217{
			para217{[]int{1, 2, 3, 1}},
			ans217{true},
		},

		question217{
			para217{[]int{1, 2, 3, 4}},
			ans217{false},
		},

		question217{
			para217{[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}},
			ans217{true},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 217------------------------\n")

	for _, q := range qs {
		_, p := q.ans217, q.para217
		fmt.Printf("【input】:%v       【output】:%v\n", p, containsDuplicate(p.one))
	}
	fmt.Printf("\n\n\n")
}
