package leetcode

import (
	"fmt"
	"testing"
)

type question709 struct {
	para709
	ans709
}

// para 是参数
// one 代表第一个参数
type para709 struct {
	one string
}

// ans 是答案
// one 代表第一个答案
type ans709 struct {
	one string
}

func Test_Problem709(t *testing.T) {

	qs := []question709{

		{
			para709{"Hello"},
			ans709{"hello"},
		},

		{
			para709{"here"},
			ans709{"here"},
		},

		{
			para709{"LOVELY"},
			ans709{"lovely"},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 709------------------------\n")

	for _, q := range qs {
		_, p := q.ans709, q.para709
		fmt.Printf("【input】:%v       【output】:%v\n", p, toLowerCase(p.one))
	}
	fmt.Printf("\n\n\n")
}
