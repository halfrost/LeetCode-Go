package leetcode

import (
	"fmt"
	"testing"
)

type question1649 struct {
	para1649
	ans1649
}

// para 是参数
// one 代表第一个参数
type para1649 struct {
	s string
}

// ans 是答案
// one 代表第一个答案
type ans1649 struct {
	one int
}

func Test_Problem1649(t *testing.T) {

	qs := []question1649{

		{
			para1649{"aababbab"},
			ans1649{2},
		},

		{
			para1649{"bbaaaaabb"},
			ans1649{2},
		},

		{
			para1649{"b"},
			ans1649{0},
		},

		{
			para1649{"ababaaaabbbbbaaababbbbbbaaabbaababbabbbbaabbbbaabbabbabaabbbababaa"},
			ans1649{25},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1649------------------------\n")

	for _, q := range qs {
		_, p := q.ans1649, q.para1649
		fmt.Printf("【input】:%v      【output】:%v      \n", p, minimumDeletions(p.s))
	}
	fmt.Printf("\n\n\n")
}
