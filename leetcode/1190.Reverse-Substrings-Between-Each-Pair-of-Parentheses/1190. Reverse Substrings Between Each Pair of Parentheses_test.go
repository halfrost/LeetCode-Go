package leetcode

import (
	"fmt"
	"testing"
)

type question1190 struct {
	para1190
	ans1190
}

// para 是参数
// one 代表第一个参数
type para1190 struct {
	s string
}

// ans 是答案
// one 代表第一个答案
type ans1190 struct {
	one string
}

func Test_Problem1190(t *testing.T) {

	qs := []question1190{

		{
			para1190{"(abcd)"},
			ans1190{"dcba"},
		},

		{
			para1190{"(u(love)i)"},
			ans1190{"iloveu"},
		},

		{
			para1190{"(ed(et(oc))el)"},
			ans1190{"leetcode"},
		},

		{
			para1190{"a(bcdefghijkl(mno)p)q"},
			ans1190{"apmnolkjihgfedcbq"},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1190------------------------\n")

	for _, q := range qs {
		_, p := q.ans1190, q.para1190
		fmt.Printf("【input】:%v       【output】:%v\n", p, reverseParentheses(p.s))
	}
	fmt.Printf("\n\n\n")
}
