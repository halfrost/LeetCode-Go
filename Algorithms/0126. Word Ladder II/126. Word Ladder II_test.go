package leetcode

import (
	"fmt"
	"testing"
)

type question126 struct {
	para126
	ans126
}

// para 是参数
// one 代表第一个参数
type para126 struct {
	b string
	e string
	w []string
}

// ans 是答案
// one 代表第一个答案
type ans126 struct {
	one [][]string
}

func Test_Problem126(t *testing.T) {

	qs := []question126{
		question126{
			para126{"hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}},
			ans126{[][]string{[]string{"hit", "hot", "dot", "dog", "cog"}, []string{"hit", "hot", "lot", "log", "cog"}}},
		},

		question126{
			para126{"hit", "cog", []string{"hot", "dot", "dog", "lot", "log"}},
			ans126{[][]string{}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 126------------------------\n")

	for _, q := range qs {
		_, p := q.ans126, q.para126
		fmt.Printf("【input】:%v       【output】:%v\n", p, findLadders(p.b, p.e, p.w))
	}
	fmt.Printf("\n\n\n")
}
