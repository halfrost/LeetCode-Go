package leetcode

import (
	"fmt"
	"testing"
)

type question1573 struct {
	para1573
	ans1573
}

// para 是参数
// one 代表第一个参数
type para1573 struct {
	s string
}

// ans 是答案
// one 代表第一个答案
type ans1573 struct {
	one int
}

func Test_Problem1573(t *testing.T) {

	qs := []question1573{

		{
			para1573{"10101"},
			ans1573{4},
		},

		{
			para1573{"1001"},
			ans1573{0},
		},

		{
			para1573{"0000"},
			ans1573{3},
		},

		{
			para1573{"100100010100110"},
			ans1573{12},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1573------------------------\n")

	for _, q := range qs {
		_, p := q.ans1573, q.para1573
		fmt.Printf("【input】:%v      【output】:%v      \n", p, numWays(p.s))
	}
	fmt.Printf("\n\n\n")
}
