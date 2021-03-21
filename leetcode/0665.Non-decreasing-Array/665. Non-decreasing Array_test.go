package leetcode

import (
	"fmt"
	"testing"
)

type question665 struct {
	para665
	ans665
}

// para 是参数
// one 代表第一个参数
type para665 struct {
	nums []int
}

// ans 是答案
// one 代表第一个答案
type ans665 struct {
	one bool
}

func Test_Problem665(t *testing.T) {

	qs := []question665{

		{
			para665{[]int{4, 2, 3}},
			ans665{true},
		},

		{
			para665{[]int{4, 2, 1}},
			ans665{false},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 665------------------------\n")

	for _, q := range qs {
		_, p := q.ans665, q.para665
		fmt.Printf("【input】:%v       【output】:%v\n", p, checkPossibility(p.nums))
	}
	fmt.Printf("\n\n\n")
}
