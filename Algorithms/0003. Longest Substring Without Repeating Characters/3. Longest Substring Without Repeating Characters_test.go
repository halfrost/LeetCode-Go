package leetcode

import (
	"fmt"
	"testing"
)

type question3 struct {
	para3
	ans3
}

// para 是参数
// one 代表第一个参数
type para3 struct {
	s string
}

// ans 是答案
// one 代表第一个答案
type ans3 struct {
	one int
}

func Test_Problem3(t *testing.T) {

	qs := []question3{

		question3{
			para3{"abcabcbb"},
			ans3{3},
		},

		question3{
			para3{"bbbbb"},
			ans3{1},
		},

		question3{
			para3{"pwwkew"},
			ans3{3},
		},

		question3{
			para3{""},
			ans3{0},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 3------------------------\n")

	for _, q := range qs {
		_, p := q.ans3, q.para3
		fmt.Printf("【input】:%v       【output】:%v\n", p, lengthOfLongestSubstring(p.s))
	}
	fmt.Printf("\n\n\n")
}
