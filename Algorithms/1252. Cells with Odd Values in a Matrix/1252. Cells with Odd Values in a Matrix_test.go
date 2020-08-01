package leetcode

import (
	"fmt"
	"testing"
)

type question1252 struct {
	para1252
	ans1252
}

// para 是参数
// one 代表第一个参数
type para1252 struct {
	n       int
	m       int
	indices [][]int
}

// ans 是答案
// one 代表第一个答案
type ans1252 struct {
	one int
}

func Test_Problem1252(t *testing.T) {

	qs := []question1252{

		question1252{
			para1252{2, 3, [][]int{[]int{0, 1}, []int{1, 1}}},
			ans1252{6},
		},

		question1252{
			para1252{2, 2, [][]int{[]int{1, 1}, []int{0, 0}}},
			ans1252{0},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1252------------------------\n")

	for _, q := range qs {
		_, p := q.ans1252, q.para1252
		fmt.Printf("【input】:%v       【output】:%v\n", p, oddCells(p.n, p.m, p.indices))
	}
	fmt.Printf("\n\n\n")
}
