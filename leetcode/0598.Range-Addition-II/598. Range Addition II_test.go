package leetcode

import (
	"fmt"
	"testing"
)

type question598 struct {
	para598
	ans598
}

// para 是参数
// one 代表第一个参数
type para598 struct {
	m   int
	n   int
	ops [][]int
}

// ans 是答案
// one 代表第一个答案
type ans598 struct {
	one int
}

func Test_Problem598(t *testing.T) {

	qs := []question598{

		{
			para598{3, 3, [][]int{{2, 2}, {3, 3}}},
			ans598{4},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 598------------------------\n")

	for _, q := range qs {
		_, p := q.ans598, q.para598
		fmt.Printf("【input】:%v       【output】:%v\n", p, maxCount(p.m, p.n, p.ops))
	}
	fmt.Printf("\n\n\n")
}
