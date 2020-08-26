package leetcode

import (
	"fmt"
	"testing"
)

type question892 struct {
	para892
	ans892
}

// para 是参数
// one 代表第一个参数
type para892 struct {
	one [][]int
}

// ans 是答案
// one 代表第一个答案
type ans892 struct {
	one int
}

func Test_Problem892(t *testing.T) {

	qs := []question892{

		{
			para892{[][]int{{2}}},
			ans892{10},
		},

		{
			para892{[][]int{{1, 2}, {3, 4}}},
			ans892{34},
		},

		{
			para892{[][]int{{1, 0}, {0, 2}}},
			ans892{16},
		},

		{
			para892{[][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}},
			ans892{32},
		},

		{
			para892{[][]int{{2, 2, 2}, {2, 1, 2}, {2, 2, 2}}},
			ans892{46},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 892------------------------\n")

	for _, q := range qs {
		_, p := q.ans892, q.para892
		fmt.Printf("【input】:%v       【output】:%v\n", p, surfaceArea(p.one))
	}
	fmt.Printf("\n\n\n")
}
