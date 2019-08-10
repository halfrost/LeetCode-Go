package leetcode

import (
	"fmt"
	"testing"
)

type question60 struct {
	para60
	ans60
}

// para 是参数
// one 代表第一个参数
type para60 struct {
	n int
	k int
}

// ans 是答案
// one 代表第一个答案
type ans60 struct {
	one string
}

func Test_Problem60(t *testing.T) {

	qs := []question60{

		question60{
			para60{3, 3},
			ans60{"213"},
		},

		question60{
			para60{4, 9},
			ans60{"2314"},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 60------------------------\n")

	for _, q := range qs {
		_, p := q.ans60, q.para60
		fmt.Printf("【input】:%v       【output】:%v\n", p, getPermutation(p.n, p.k))
	}
	fmt.Printf("\n\n\n")
}
