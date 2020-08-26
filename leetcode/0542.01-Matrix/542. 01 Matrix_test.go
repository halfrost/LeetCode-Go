package leetcode

import (
	"fmt"
	"testing"
)

type question542 struct {
	para542
	ans542
}

// para 是参数
// one 代表第一个参数
type para542 struct {
	one [][]int
}

// ans 是答案
// one 代表第一个答案
type ans542 struct {
	one [][]int
}

func Test_Problem542(t *testing.T) {

	qs := []question542{

		{
			para542{[][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}},
			ans542{[][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}},
		},

		{
			para542{[][]int{{0, 0, 0}, {0, 1, 0}, {1, 1, 1}}},
			ans542{[][]int{{0, 0, 0}, {0, 1, 0}, {1, 2, 1}}},
		},

		{
			para542{[][]int{{1, 1, 1, 1, 1, 1}, {1, 1, 1, 1, 1, 1}, {1, 1, 0, 0, 1, 1}, {1, 1, 0, 0, 1, 1}, {1, 1, 1, 1, 1, 1}, {1, 1, 1, 1, 1, 1}}},
			ans542{[][]int{{4, 3, 2, 2, 3, 4}, {3, 2, 1, 1, 2, 3}, {2, 1, 0, 0, 1, 2}, {2, 1, 0, 0, 1, 2}, {3, 2, 1, 1, 2, 3}, {4, 3, 2, 2, 3, 4}}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 542------------------------\n")

	for _, q := range qs {
		_, p := q.ans542, q.para542
		fmt.Printf("【input】:%v       【output】:%v\n", p, updateMatrixDP(p.one))
	}
	fmt.Printf("\n\n\n")
}
