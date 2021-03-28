package leetcode

import (
	"fmt"
	"testing"
)

type question647 struct {
	para647
	ans647
}

// para 是参数
// one 代表第一个参数
type para647 struct {
	s string
}

// ans 是答案
// one 代表第一个答案
type ans647 struct {
	one int
}

func Test_Problem647(t *testing.T) {

	qs := []question647{

		{
			para647{"abc"},
			ans647{3},
		},

		{
			para647{"aaa"},
			ans647{6},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 647------------------------\n")

	for _, q := range qs {
		_, p := q.ans647, q.para647
		fmt.Printf("【input】:%v       【output】:%v\n", p, countSubstrings(p.s))
	}
	fmt.Printf("\n\n\n")
}
