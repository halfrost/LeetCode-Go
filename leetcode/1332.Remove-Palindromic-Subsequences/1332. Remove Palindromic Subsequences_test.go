package leetcode

import (
	"fmt"
	"testing"
)

type question1332 struct {
	para1332
	ans1332
}

// para 是参数
// one 代表第一个参数
type para1332 struct {
	s string
}

// ans 是答案
// one 代表第一个答案
type ans1332 struct {
	one int
}

func Test_Problem1332(t *testing.T) {

	qs := []question1332{

		{
			para1332{"ababa"},
			ans1332{1},
		},

		{
			para1332{"abb"},
			ans1332{2},
		},

		{
			para1332{"baabb"},
			ans1332{2},
		},

		{
			para1332{""},
			ans1332{0},
		},

		{
			para1332{"bbaabaaa"},
			ans1332{2},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1332------------------------\n")

	for _, q := range qs {
		_, p := q.ans1332, q.para1332
		fmt.Printf("【input】:%v       【output】:%v\n", p, removePalindromeSub(p.s))
	}
	fmt.Printf("\n\n\n")
}
