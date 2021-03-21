package leetcode

import (
	"fmt"
	"testing"
)

type question1704 struct {
	para1704
	ans1704
}

// para 是参数
// one 代表第一个参数
type para1704 struct {
	s string
}

// ans 是答案
// one 代表第一个答案
type ans1704 struct {
	one bool
}

func Test_Problem1704(t *testing.T) {

	qs := []question1704{

		{
			para1704{"book"},
			ans1704{true},
		},

		{
			para1704{"textbook"},
			ans1704{false},
		},

		{
			para1704{"MerryChristmas"},
			ans1704{false},
		},

		{
			para1704{"AbCdEfGh"},
			ans1704{true},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1704------------------------\n")

	for _, q := range qs {
		_, p := q.ans1704, q.para1704
		fmt.Printf("【input】:%v       【output】:%v\n", p, halvesAreAlike(p.s))
	}
	fmt.Printf("\n\n\n")
}
