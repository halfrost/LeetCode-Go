package leetcode

import (
	"fmt"
	"testing"
)

type question1154 struct {
	para1154
	ans1154
}

// para 是参数
// one 代表第一个参数
type para1154 struct {
	one string
}

// ans 是答案
// one 代表第一个答案
type ans1154 struct {
	one int
}

func Test_Problem1154(t *testing.T) {

	qs := []question1154{

		question1154{
			para1154{"2019-01-09"},
			ans1154{9},
		},

		question1154{
			para1154{"2019-02-10"},
			ans1154{41},
		},

		question1154{
			para1154{"2003-03-01"},
			ans1154{60},
		},

		question1154{
			para1154{"2004-03-01"},
			ans1154{61},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1154------------------------\n")

	for _, q := range qs {
		_, p := q.ans1154, q.para1154
		fmt.Printf("【input】:%v       【output】:%v\n", p, dayOfYear(p.one))
	}
	fmt.Printf("\n\n\n")
}
