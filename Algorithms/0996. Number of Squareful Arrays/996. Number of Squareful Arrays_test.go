package leetcode

import (
	"fmt"
	"testing"
)

type question996 struct {
	para996
	ans996
}

// para 是参数
// one 代表第一个参数
type para996 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans996 struct {
	one int
}

func Test_Problem996(t *testing.T) {

	qs := []question996{

		question996{
			para996{[]int{1, 17, 8}},
			ans996{2},
		},

		question996{
			para996{[]int{1}},
			ans996{1},
		},

		question996{
			para996{[]int{2, 2, 2}},
			ans996{1},
		},

		question996{
			para996{[]int{51768, 47861, 48143, 33221, 50893, 56758, 39946, 10312, 20276, 40616, 43633}},
			ans996{1},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 996------------------------\n")

	for _, q := range qs {
		_, p := q.ans996, q.para996
		fmt.Printf("【input】:%v       【output】:%v\n", p, numSquarefulPerms(p.one))
	}
	fmt.Printf("\n\n\n")
}
