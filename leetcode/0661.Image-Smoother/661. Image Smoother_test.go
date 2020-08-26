package leetcode

import (
	"fmt"
	"testing"
)

type question661 struct {
	para661
	ans661
}

// para 是参数
// one 代表第一个参数
type para661 struct {
	one [][]int
}

// ans 是答案
// one 代表第一个答案
type ans661 struct {
	one [][]int
}

func Test_Problem661(t *testing.T) {

	qs := []question661{

		{
			para661{[][]int{{1, 1, 1}, {1, 1, 2}}},
			ans661{[][]int{{1, 1, 1}, {1, 1, 1}}},
		},

		{
			para661{[][]int{{1, 1, 1}, {1, 1, 2}, {1, 1, 1}}},
			ans661{[][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}}},
		},

		{
			para661{[][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}},
			ans661{[][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 661------------------------\n")

	for _, q := range qs {
		_, p := q.ans661, q.para661
		fmt.Printf("【input】:%v       【output】:%v\n", p, imageSmoother(p.one))
	}
	fmt.Printf("\n\n\n")
}
