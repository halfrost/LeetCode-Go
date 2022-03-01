package leetcode

import (
	"fmt"
	"testing"
)

type question2170 struct {
	para2170
	ans2170
}

// para 是参数
// one 代表第一个参数
type para2170 struct {
	beans []int
}

// ans 是答案
// one 代表第一个答案
type ans2170 struct {
	one int
}

func Test_Problem1(t *testing.T) {

	qs := []question2170{

		{
			para2170{[]int{4, 1, 6, 5}},
			ans2170{4},
		},

		{
			para2170{[]int{2, 10, 3, 2}},
			ans2170{7},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 2170------------------------\n")

	for _, q := range qs {
		_, p := q.ans2170, q.para2170
		fmt.Printf("【input】:%v       【output】:%v\n", p, minimumRemoval(p.beans))
	}
	fmt.Printf("\n\n\n")
}
