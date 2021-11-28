package leetcode

import (
	"fmt"
	"testing"
)

type question58 struct {
	para58
	ans58
}

// para 是参数
type para58 struct {
	s string
}

// ans 是答案
type ans58 struct {
	ans int
}

func Test_Problem58(t *testing.T) {

	qs := []question58{

		{
			para58{"Hello World"},
			ans58{5},
		},

		{
			para58{"   fly me   to   the moon  "},
			ans58{4},
		},

		{
			para58{"luffy is still joyboy"},
			ans58{6},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 58------------------------\n")

	for _, q := range qs {
		_, p := q.ans58, q.para58
		fmt.Printf("【input】:%v       【output】:%v\n", p, lengthOfLastWord(p.s))
	}
	fmt.Printf("\n\n\n")
}
