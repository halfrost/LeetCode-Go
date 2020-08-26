package leetcode

import (
	"fmt"
	"testing"
)

type question862 struct {
	para862
	ans862
}

// para 是参数
// one 代表第一个参数
type para862 struct {
	A []int
	K int
}

// ans 是答案
// one 代表第一个答案
type ans862 struct {
	one int
}

func Test_Problem862(t *testing.T) {

	qs := []question862{
		{
			para862{[]int{1}, 1},
			ans862{1},
		},

		{
			para862{[]int{1, 2}, 4},
			ans862{-1},
		},

		{
			para862{[]int{2, -1, 2}, 3},
			ans862{3},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 862------------------------\n")

	for _, q := range qs {
		_, p := q.ans862, q.para862
		fmt.Printf("【input】:%v       【output】:%v\n", p, shortestSubarray(p.A, p.K))
	}
	fmt.Printf("\n\n\n")
}
