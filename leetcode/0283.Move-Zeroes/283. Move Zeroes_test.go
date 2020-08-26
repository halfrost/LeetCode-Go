package leetcode

import (
	"fmt"
	"testing"
)

type question283 struct {
	para283
	ans283
}

// para 是参数
// one 代表第一个参数
type para283 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans283 struct {
	one []int
}

func Test_Problem283(t *testing.T) {

	qs := []question283{

		{
			para283{[]int{1, 0, 1}},
			ans283{[]int{1, 1, 0}},
		},

		{
			para283{[]int{0, 1, 0, 3, 0, 12}},
			ans283{[]int{1, 3, 12, 0, 0, 0}},
		},

		{
			para283{[]int{0, 1, 0, 3, 0, 0, 0, 0, 1, 12}},
			ans283{[]int{1, 3, 1, 12, 0, 0, 0, 0, 0}},
		},

		{
			para283{[]int{0, 0, 0, 0, 0, 0, 0, 0, 12, 1}},
			ans283{[]int{12, 1, 0, 0, 0, 0, 0, 0, 0, 0}},
		},

		{
			para283{[]int{0, 0, 0, 0, 0}},
			ans283{[]int{0, 0, 0, 0, 0}},
		},

		{
			para283{[]int{1}},
			ans283{[]int{1}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 283------------------------\n")

	for _, q := range qs {
		_, p := q.ans283, q.para283
		fmt.Printf("【input】:%v      ", p.one)
		moveZeroes(p.one)
		fmt.Printf("【output】:%v\n", p.one)
	}
	fmt.Printf("\n\n\n")
}
