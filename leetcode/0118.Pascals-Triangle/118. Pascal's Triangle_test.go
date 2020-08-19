package leetcode

import (
	"fmt"
	"testing"
)

type question118 struct {
	para118
	ans118
}

// para 是参数
// one 代表第一个参数
type para118 struct {
	numRows int
}

// ans 是答案
// one 代表第一个答案
type ans118 struct {
	one [][]int
}

func Test_Problem118(t *testing.T) {

	qs := []question118{

		question118{
			para118{2},
			ans118{[][]int{{1}, {1, 1}}},
		},

		question118{
			para118{5},
			ans118{[][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}}},
		},

		question118{
			para118{10},
			ans118{[][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}, {1, 5, 10, 10, 5, 1}, {1, 6, 15, 20, 15, 6, 1}, {1, 7, 21, 35, 35, 21, 7, 1}, {1, 8, 28, 56, 70, 56, 28, 8, 1}, {1, 9, 36, 84, 126, 126, 84, 36, 9, 1}}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 118------------------------\n")

	for _, q := range qs {
		_, p := q.ans118, q.para118
		fmt.Printf("【input】:%v    【output】:%v\n", p, generate(p.numRows))
	}
	fmt.Printf("\n\n\n")
}
