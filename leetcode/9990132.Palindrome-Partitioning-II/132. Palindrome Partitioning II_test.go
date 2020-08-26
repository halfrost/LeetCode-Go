package leetcode

import (
	"fmt"
	"testing"
)

type question132 struct {
	para132
	ans132
}

// para 是参数
// one 代表第一个参数
type para132 struct {
	s string
}

// ans 是答案
// one 代表第一个答案
type ans132 struct {
	one int
}

func Test_Problem132(t *testing.T) {

	qs := []question132{

		{
			para132{"aab"},
			ans132{1},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 132------------------------\n")

	for _, q := range qs {
		_, p := q.ans132, q.para132
		fmt.Printf("【input】:%v       【output】:%v\n", p, minCut(p.s))
	}
	fmt.Printf("\n\n\n")
}
