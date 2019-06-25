package leetcode

import (
	"fmt"
	"testing"
)

type question343 struct {
	para343
	ans343
}

// para 是参数
// one 代表第一个参数
type para343 struct {
	one int
}

// ans 是答案
// one 代表第一个答案
type ans343 struct {
	one int
}

func Test_Problem343(t *testing.T) {

	qs := []question343{

		question343{
			para343{2},
			ans343{1},
		},

		question343{
			para343{10},
			ans343{36},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 343------------------------\n")

	for _, q := range qs {
		_, p := q.ans343, q.para343
		fmt.Printf("【input】:%v       【output】:%v\n", p, integerBreak(p.one))
	}
	fmt.Printf("\n\n\n")
}
