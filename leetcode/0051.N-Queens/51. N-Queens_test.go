package leetcode

import (
	"fmt"
	"testing"
)

type question51 struct {
	para51
	ans51
}

// para 是参数
// one 代表第一个参数
type para51 struct {
	one int
}

// ans 是答案
// one 代表第一个答案
type ans51 struct {
	one [][]string
}

func Test_Problem51(t *testing.T) {

	qs := []question51{

		question51{
			para51{4},
			ans51{[][]string{
				[]string{".Q..",
					"...Q",
					"Q...",
					"..Q."},
				[]string{"..Q.",
					"Q...",
					"...Q",
					".Q.."},
			}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 51------------------------\n")
	for _, q := range qs {
		_, p := q.ans51, q.para51
		fmt.Printf("【input】:%v       【output】:%v\n", p, solveNQueens(p.one))
	}
	fmt.Printf("\n\n\n")
}
