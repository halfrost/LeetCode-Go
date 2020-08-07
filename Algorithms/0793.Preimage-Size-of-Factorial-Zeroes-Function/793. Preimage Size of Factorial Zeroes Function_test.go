package leetcode

import (
	"fmt"
	"testing"
)

type question793 struct {
	para793
	ans793
}

// para 是参数
// one 代表第一个参数
type para793 struct {
	one int
}

// ans 是答案
// one 代表第一个答案
type ans793 struct {
	one int
}

func Test_Problem793(t *testing.T) {

	qs := []question793{

		question793{
			para793{0},
			ans793{5},
		},

		question793{
			para793{5},
			ans793{0},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 793------------------------\n")

	for _, q := range qs {
		_, p := q.ans793, q.para793
		fmt.Printf("【input】:%v       【output】:%v\n", p, preimageSizeFZF(p.one))
	}
	fmt.Printf("\n\n\n")
}
