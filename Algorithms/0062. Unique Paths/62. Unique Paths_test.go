package leetcode

import (
	"fmt"
	"testing"
)

type question62 struct {
	para62
	ans62
}

// para 是参数
// one 代表第一个参数
type para62 struct {
	m int
	n int
}

// ans 是答案
// one 代表第一个答案
type ans62 struct {
	one int
}

func Test_Problem62(t *testing.T) {

	qs := []question62{

		question62{
			para62{3, 2},
			ans62{3},
		},

		question62{
			para62{7, 3},
			ans62{28},
		},

		question62{
			para62{1, 2},
			ans62{1},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 62------------------------\n")

	for _, q := range qs {
		_, p := q.ans62, q.para62
		fmt.Printf("【input】:%v       【output】:%v\n", p, uniquePaths(p.m, p.n))
	}
	fmt.Printf("\n\n\n")
}
