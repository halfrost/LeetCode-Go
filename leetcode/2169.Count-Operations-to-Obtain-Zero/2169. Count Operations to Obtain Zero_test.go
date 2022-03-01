package leetcode

import (
	"fmt"
	"testing"
)

type question2169 struct {
	para2169
	ans2169
}

// para 是参数
// one 代表第一个参数
type para2169 struct {
	num1 int
	num2 int
}

// ans 是答案
// one 代表第一个答案
type ans2169 struct {
	one int
}

func Test_Problem2169(t *testing.T) {

	qs := []question2169{

		{
			para2169{2, 3},
			ans2169{3},
		},

		{
			para2169{10, 10},
			ans2169{1},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 2169------------------------\n")

	for _, q := range qs {
		_, p := q.ans2169, q.para2169
		fmt.Printf("【input】:%v       【output】:%v\n", p, countOperations(p.num1, p.num2))
	}
	fmt.Printf("\n\n\n")
}
