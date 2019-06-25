package leetcode

import (
	"fmt"
	"testing"
)

type question64 struct {
	para64
	ans64
}

// para 是参数
// one 代表第一个参数
type para64 struct {
	og [][]int
}

// ans 是答案
// one 代表第一个答案
type ans64 struct {
	one int
}

func Test_Problem64(t *testing.T) {

	qs := []question64{

		question64{
			para64{[][]int{
				[]int{1, 3, 1},
				[]int{1, 5, 1},
				[]int{4, 2, 1},
			}},
			ans64{7},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 64------------------------\n")

	for _, q := range qs {
		_, p := q.ans64, q.para64
		fmt.Printf("【input】:%v       【output】:%v\n", p, minPathSum(p.og))
	}
	fmt.Printf("\n\n\n")
}
