package leetcode

import (
	"fmt"
	"testing"
)

type question792 struct {
	para792
	ans792
}

// para 是参数
// one 代表第一个参数
type para792 struct {
	s     string
	words []string
}

// ans 是答案
// one 代表第一个答案
type ans792 struct {
	one int
}

func Test_Problem792(t *testing.T) {

	qs := []question792{

		{
			para792{"abcde", []string{"a", "bb", "acd", "ace"}},
			ans792{3},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 792------------------------\n")

	for _, q := range qs {
		_, p := q.ans792, q.para792
		fmt.Printf("【input】:%v       【output】:%v\n", p, numMatchingSubseq(p.s, p.words))
	}
	fmt.Printf("\n\n\n")
}
