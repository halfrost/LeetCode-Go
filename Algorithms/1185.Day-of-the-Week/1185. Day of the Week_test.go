package leetcode

import (
	"fmt"
	"testing"
)

type question1185 struct {
	para1185
	ans1185
}

// para 是参数
// one 代表第一个参数
type para1185 struct {
	day   int
	month int
	year  int
}

// ans 是答案
// one 代表第一个答案
type ans1185 struct {
	one string
}

func Test_Problem1185(t *testing.T) {

	qs := []question1185{

		question1185{
			para1185{31, 8, 2019},
			ans1185{"Saturday"},
		},

		question1185{
			para1185{18, 7, 1999},
			ans1185{"Sunday"},
		},

		question1185{
			para1185{15, 8, 1993},
			ans1185{"Sunday"},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1185------------------------\n")

	for _, q := range qs {
		_, p := q.ans1185, q.para1185
		fmt.Printf("【input】:%v       【output】:%v\n", p, dayOfTheWeek(p.day, p.month, p.year))
	}
	fmt.Printf("\n\n\n")
}
