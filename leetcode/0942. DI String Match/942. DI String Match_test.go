package leetcode

import (
	"fmt"
	"testing"
)

type question942 struct {
	para942
	ans942
}

// para 是参数
// one 代表第一个参数
type para942 struct {
	S string
}

// ans 是答案
// one 代表第一个答案
type ans942 struct {
	one []int
}

func Test_Problem942(t *testing.T) {

	qs := []question942{

		question942{
			para942{"IDID"},
			ans942{[]int{0, 4, 1, 3, 2}},
		},

		question942{
			para942{"III"},
			ans942{[]int{0, 1, 2, 3}},
		},

		question942{
			para942{"DDI"},
			ans942{[]int{3, 2, 0, 1}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 942------------------------\n")

	for _, q := range qs {
		_, p := q.ans942, q.para942
		fmt.Printf("【input】:%v       【output】:%v\n", p, diStringMatch(p.S))
	}
	fmt.Printf("\n\n\n")
}
