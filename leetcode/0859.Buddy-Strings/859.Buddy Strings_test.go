package leetcode

import (
	"fmt"
	"testing"
)

type question859 struct {
	para859
	ans859
}

// para 是参数
type para859 struct {
	s    string
	goal string
}

// ans 是答案
type ans859 struct {
	ans bool
}

func Test_Problem859(t *testing.T) {

	qs := []question859{

		{
			para859{"ab", "ba"},
			ans859{true},
		},

		{
			para859{"ab", "ab"},
			ans859{false},
		},

		{
			para859{"aa", "aa"},
			ans859{true},
		},

		{
			para859{"aaaaaaabc", "aaaaaaacb"},
			ans859{true},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 859------------------------\n")

	for _, q := range qs {
		_, p := q.ans859, q.para859
		fmt.Printf("【input】:%v    【output】:%v\n", p, buddyStrings(p.s, p.goal))
	}
	fmt.Printf("\n\n\n")
}
