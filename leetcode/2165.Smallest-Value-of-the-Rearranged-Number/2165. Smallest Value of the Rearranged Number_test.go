package leetcode

import (
	"fmt"
	"testing"
)

type question2165 struct {
	para2165
	ans2165
}

// para 是参数
// one 代表第一个参数
type para2165 struct {
	nums int64
}

// ans 是答案
// one 代表第一个答案
type ans2165 struct {
	one int64
}

func Test_Problem1(t *testing.T) {

	qs := []question2165{

		{
			para2165{310},
			ans2165{103},
		},

		{
			para2165{5059},
			ans2165{5059},
		},

		{
			para2165{-7605},
			ans2165{-7650},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 2165------------------------\n")

	for _, q := range qs {
		_, p := q.ans2165, q.para2165
		fmt.Printf("【input】:%v       【output】:%v\n", p, smallestNumber(p.nums))
	}
	fmt.Printf("\n\n\n")
}
