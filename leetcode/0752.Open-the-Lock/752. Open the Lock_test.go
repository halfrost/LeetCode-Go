package leetcode

import (
	"fmt"
	"testing"
)

type question752 struct {
	para752
	ans752
}

// para 是参数
// one 代表第一个参数
type para752 struct {
	deadends []string
	target   string
}

// ans 是答案
// one 代表第一个答案
type ans752 struct {
	one int
}

func Test_Problem752(t *testing.T) {

	qs := []question752{

		{
			para752{[]string{"0201", "0101", "0102", "1212", "2002"}, "0202"},
			ans752{6},
		},

		{
			para752{[]string{"8888"}, "0009"},
			ans752{1},
		},

		{
			para752{[]string{"8887", "8889", "8878", "8898", "8788", "8988", "7888", "9888"}, "8888"},
			ans752{-1},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 752------------------------\n")

	for _, q := range qs {
		_, p := q.ans752, q.para752
		fmt.Printf("【input】:%v       【output】:%v\n", p, openLock(p.deadends, p.target))
	}
	fmt.Printf("\n\n\n")
}
