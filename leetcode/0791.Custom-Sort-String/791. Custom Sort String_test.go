package leetcode

import (
	"fmt"
	"testing"
)

type question791 struct {
	para791
	ans791
}

// para 是参数
// one 代表第一个参数
type para791 struct {
	order string
	str   string
}

// ans 是答案
// one 代表第一个答案
type ans791 struct {
	one string
}

func Test_Problem791(t *testing.T) {

	qs := []question791{

		{
			para791{"cba", "abcd"},
			ans791{"cbad"},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 791------------------------\n")

	for _, q := range qs {
		_, p := q.ans791, q.para791
		fmt.Printf("【input】:%v       【output】:%v\n", p, customSortString(p.order, p.str))
	}
	fmt.Printf("\n\n\n")
}
