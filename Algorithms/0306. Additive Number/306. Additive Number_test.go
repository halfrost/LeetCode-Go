package leetcode

import (
	"fmt"
	"testing"
)

type question306 struct {
	para306
	ans306
}

// para 是参数
// one 代表第一个参数
type para306 struct {
	one string
}

// ans 是答案
// one 代表第一个答案
type ans306 struct {
	one bool
}

func Test_Problem306(t *testing.T) {

	qs := []question306{

		question306{
			para306{"112358"},
			ans306{true},
		},

		question306{
			para306{"199100199"},
			ans306{true},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 306------------------------\n")

	for _, q := range qs {
		_, p := q.ans306, q.para306
		fmt.Printf("【input】:%v       【output】:%v\n", p, isAdditiveNumber(p.one))
	}
	fmt.Printf("\n\n\n")
}
