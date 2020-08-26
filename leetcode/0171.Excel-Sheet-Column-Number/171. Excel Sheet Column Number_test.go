package leetcode

import (
	"fmt"
	"testing"
)

type question171 struct {
	para171
	ans171
}

// para 是参数
// one 代表第一个参数
type para171 struct {
	s string
}

// ans 是答案
// one 代表第一个答案
type ans171 struct {
	one int
}

func Test_Problem171(t *testing.T) {

	qs := []question171{

		{
			para171{"A"},
			ans171{1},
		},

		{
			para171{"AB"},
			ans171{28},
		},

		{
			para171{"ZY"},
			ans171{701},
		},

		{
			para171{"ABC"},
			ans171{731},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 171------------------------\n")

	for _, q := range qs {
		_, p := q.ans171, q.para171
		fmt.Printf("【input】:%v       【output】:%v\n", p, titleToNumber(p.s))
	}
	fmt.Printf("\n\n\n")
}
