package leetcode

import (
	"fmt"
	"testing"
)

type question819 struct {
	para819
	ans819
}

// para 是参数
// one 代表第一个参数
type para819 struct {
	one string
	b   []string
}

// ans 是答案
// one 代表第一个答案
type ans819 struct {
	one string
}

func Test_Problem819(t *testing.T) {

	qs := []question819{

		question819{
			para819{"Bob hit a ball, the hit BALL flew far after it was hit.", []string{"hit"}},
			ans819{"ball"},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 819------------------------\n")

	for _, q := range qs {
		_, p := q.ans819, q.para819
		fmt.Printf("【input】:%v       【output】:%v\n", p, mostCommonWord(p.one, p.b))
	}
	fmt.Printf("\n\n\n")
}
