package leetcode

import (
	"fmt"
	"testing"
)

type question1647 struct {
	para1647
	ans1647
}

// para 是参数
// one 代表第一个参数
type para1647 struct {
	s string
}

// ans 是答案
// one 代表第一个答案
type ans1647 struct {
	one int
}

func Test_Problem1647(t *testing.T) {

	qs := []question1647{

		{
			para1647{"aab"},
			ans1647{0},
		},

		{
			para1647{"aaabbbcc"},
			ans1647{2},
		},

		{
			para1647{"ceabaacb"},
			ans1647{2},
		},

		{
			para1647{""},
			ans1647{0},
		},

		{
			para1647{"abcabc"},
			ans1647{3},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1647------------------------\n")

	for _, q := range qs {
		_, p := q.ans1647, q.para1647
		fmt.Printf("【input】:%v      【output】:%v      \n", p, minDeletions(p.s))
	}
	fmt.Printf("\n\n\n")
}
