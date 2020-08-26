package leetcode

import (
	"fmt"
	"testing"
)

type question9 struct {
	para9
	ans9
}

// para 是参数
// one 代表第一个参数
type para9 struct {
	one int
}

// ans 是答案
// one 代表第一个答案
type ans9 struct {
	one bool
}

func Test_Problem9(t *testing.T) {

	qs := []question9{

		{
			para9{121},
			ans9{true},
		},

		{
			para9{-121},
			ans9{false},
		},

		{
			para9{10},
			ans9{false},
		},

		{
			para9{321},
			ans9{false},
		},

		{
			para9{-123},
			ans9{false},
		},

		{
			para9{120},
			ans9{false},
		},

		{
			para9{1534236469},
			ans9{false},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 9------------------------\n")

	for _, q := range qs {
		_, p := q.ans9, q.para9
		fmt.Printf("【input】:%v    【output】:%v\n", p.one, isPalindrome(p.one))
	}
	fmt.Printf("\n\n\n")
}
