package leetcode

import (
	"fmt"
	"testing"
)

type question867 struct {
	para867
	ans867
}

// para 是参数
// one 代表第一个参数
type para867 struct {
	A [][]int
}

// ans 是答案
// one 代表第一个答案
type ans867 struct {
	B [][]int
}

func Test_Problem867(t *testing.T) {

	qs := []question867{

		{
			para867{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}},
			ans867{[][]int{{1, 4, 7}, {2, 5, 8}, {3, 6, 9}}},
		},

		{
			para867{[][]int{{1, 2, 3}, {4, 5, 6}}},
			ans867{[][]int{{1, 4}, {2, 5}, {3, 6}}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 867------------------------\n")

	for _, q := range qs {
		_, p := q.ans867, q.para867
		fmt.Printf("【input】:%v       【output】:%v\n", p, transpose(p.A))
	}
	fmt.Printf("\n\n\n")
}
