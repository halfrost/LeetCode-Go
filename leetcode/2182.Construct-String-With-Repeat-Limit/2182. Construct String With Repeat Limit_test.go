package leetcode

import (
	"fmt"
	"testing"
)

type question2182 struct {
	para2182
	ans2182
}

// para 是参数
// one 代表第一个参数
type para2182 struct {
	one   string
	limit int
}

// ans 是答案
// one 代表第一个答案
type ans2182 struct {
	one string
}

func Test_Problem2182(t *testing.T) {

	qs := []question2182{

		{
			para2182{"cczazcc", 3},
			ans2182{"zzcccac"},
		},

		{
			para2182{"aababab", 2},
			ans2182{"bbabaa"},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 2182------------------------\n")

	for _, q := range qs {
		_, p := q.ans2182, q.para2182
		fmt.Printf("【input】:%v       【output】:%v\n", p, repeatLimitedString(p.one, p.limit))
	}
	fmt.Printf("\n\n\n")
}
