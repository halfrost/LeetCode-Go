package leetcode

import (
	"fmt"
	"testing"
)

type question947 struct {
	para947
	ans947
}

// para 是参数
// one 代表第一个参数
type para947 struct {
	stones [][]int
}

// ans 是答案
// one 代表第一个答案
type ans947 struct {
	one int
}

func Test_Problem947(t *testing.T) {

	qs := []question947{
		{
			para947{[][]int{{0, 0}, {0, 1}, {1, 0}, {1, 2}, {2, 1}, {2, 2}}},
			ans947{5},
		},

		{
			para947{[][]int{{0, 0}, {0, 2}, {1, 1}, {2, 0}, {2, 2}}},
			ans947{3},
		},

		{
			para947{[][]int{{0, 0}}},
			ans947{0},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 947------------------------\n")

	for _, q := range qs {
		_, p := q.ans947, q.para947
		fmt.Printf("【input】:%v       【output】:%v\n", p, removeStones(p.stones))
	}
	fmt.Printf("\n\n\n")
}
