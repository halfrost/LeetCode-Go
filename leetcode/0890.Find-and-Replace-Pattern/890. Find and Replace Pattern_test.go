package leetcode

import (
	"fmt"
	"testing"
)

type question890 struct {
	para890
	ans890
}

// para 是参数
// one 代表第一个参数
type para890 struct {
	words   []string
	pattern string
}

// ans 是答案
// one 代表第一个答案
type ans890 struct {
	one []string
}

func Test_Problem890(t *testing.T) {

	qs := []question890{

		{
			para890{[]string{"abc", "deq", "mee", "aqq", "dkd", "ccc"}, "abb"},
			ans890{[]string{"mee", "aqq"}},
		},

		{
			para890{[]string{"a", "b", "c"}, "a"},
			ans890{[]string{"a", "b", "c"}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 890------------------------\n")

	for _, q := range qs {
		_, p := q.ans890, q.para890
		fmt.Printf("【input】:%v       【output】:%v\n", p, findAndReplacePattern(p.words, p.pattern))
	}
	fmt.Printf("\n\n\n")
}
