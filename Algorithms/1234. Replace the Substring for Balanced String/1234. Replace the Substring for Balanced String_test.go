package leetcode

import (
	"fmt"
	"testing"
)

type question1234 struct {
	para1234
	ans1234
}

// para 是参数
// one 代表第一个参数
type para1234 struct {
	s string
}

// ans 是答案
// one 代表第一个答案
type ans1234 struct {
	one int
}

func Test_Problem1234(t *testing.T) {

	qs := []question1234{

		question1234{
			para1234{"QWER"},
			ans1234{0},
		},

		question1234{
			para1234{"QQWE"},
			ans1234{1},
		},

		question1234{
			para1234{"QQQW"},
			ans1234{2},
		},

		question1234{
			para1234{"QQQQ"},
			ans1234{3},
		},

		question1234{
			para1234{"WQWRQQQW"},
			ans1234{3},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1234------------------------\n")

	for _, q := range qs {
		_, p := q.ans1234, q.para1234
		fmt.Printf("【input】:%v       【output】:%v\n", p, balancedString(p.s))
	}
	fmt.Printf("\n\n\n")
}
