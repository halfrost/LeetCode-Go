package leetcode

import (
	"fmt"
	"testing"
)

type question201 struct {
	para201
	ans201
}

// para 是参数
// one 代表第一个参数
type para201 struct {
	m int
	n int
}

// ans 是答案
// one 代表第一个答案
type ans201 struct {
	one int
}

func Test_Problem201(t *testing.T) {

	qs := []question201{

		question201{
			para201{5, 7},
			ans201{4},
		},

		question201{
			para201{0, 1},
			ans201{0},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 201------------------------\n")

	for _, q := range qs {
		_, p := q.ans201, q.para201
		fmt.Printf("【input】:%v       【output】:%v\n", p, rangeBitwiseAnd(p.m, p.n))
	}
	fmt.Printf("\n\n\n")
}
