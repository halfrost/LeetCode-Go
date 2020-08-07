package leetcode

import (
	"fmt"
	"testing"
)

type question393 struct {
	para393
	ans393
}

// para 是参数
// one 代表第一个参数
type para393 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans393 struct {
	one bool
}

func Test_Problem393(t *testing.T) {

	qs := []question393{

		question393{
			para393{[]int{197, 130, 1}},
			ans393{true},
		},

		question393{
			para393{[]int{235, 140, 4}},
			ans393{false},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 393------------------------\n")

	for _, q := range qs {
		_, p := q.ans393, q.para393
		fmt.Printf("【input】:%v       【output】:%v\n", p, validUtf8(p.one))
	}
	fmt.Printf("\n\n\n")
}
