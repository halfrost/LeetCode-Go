package leetcode

import (
	"fmt"
	"testing"
)

type question476 struct {
	para476
	ans476
}

// para 是参数
// one 代表第一个参数
type para476 struct {
	one int
}

// ans 是答案
// one 代表第一个答案
type ans476 struct {
	one int
}

func Test_Problem476(t *testing.T) {

	qs := []question476{

		question476{
			para476{5},
			ans476{2},
		},

		question476{
			para476{1},
			ans476{0},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 476------------------------\n")

	for _, q := range qs {
		_, p := q.ans476, q.para476
		fmt.Printf("【input】:%v       【output】:%v\n", p, findComplement(p.one))
	}
	fmt.Printf("\n\n\n")
}
