package leetcode

import (
	"fmt"
	"testing"
)

type question1221 struct {
	para1221
	ans1221
}

// para 是参数
// one 代表第一个参数
type para1221 struct {
	s string
}

// ans 是答案
// one 代表第一个答案
type ans1221 struct {
	one int
}

func Test_Problem1221(t *testing.T) {

	qs := []question1221{

		question1221{
			para1221{"RLRRLLRLRL"},
			ans1221{4},
		},

		question1221{
			para1221{"RLLLLRRRLR"},
			ans1221{3},
		},

		question1221{
			para1221{"LLLLRRRR"},
			ans1221{1},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1221------------------------\n")

	for _, q := range qs {
		_, p := q.ans1221, q.para1221
		fmt.Printf("【input】:%v       【output】:%v\n", p, balancedStringSplit(p.s))
	}
	fmt.Printf("\n\n\n")
}
