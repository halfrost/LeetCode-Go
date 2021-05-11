package leetcode

import (
	"fmt"
	"testing"
)

type question583 struct {
	para583
	ans583
}

// para 是参数
// one 代表第一个参数
type para583 struct {
	word1 string
	word2 string
}

// ans 是答案
// one 代表第一个答案
type ans583 struct {
	one int
}

func Test_Problem583(t *testing.T) {

	qs := []question583{

		{
			para583{"sea", "eat"},
			ans583{2},
		},

		{
			para583{"leetcode", "etco"},
			ans583{4},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 583------------------------\n")

	for _, q := range qs {
		_, p := q.ans583, q.para583
		fmt.Printf("【input】:%v       【output】:%v\n", p, minDistance(p.word1, p.word2))
	}
	fmt.Printf("\n\n\n")
}
