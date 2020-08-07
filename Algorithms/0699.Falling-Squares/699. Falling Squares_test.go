package leetcode

import (
	"fmt"
	"testing"
)

type question699 struct {
	para699
	ans699
}

// para 是参数
// one 代表第一个参数
type para699 struct {
	one [][]int
}

// ans 是答案
// one 代表第一个答案
type ans699 struct {
	one []int
}

func Test_Problem699(t *testing.T) {

	qs := []question699{

		question699{
			para699{[][]int{[]int{6, 1}, []int{9, 2}, []int{2, 4}}},
			ans699{[]int{1, 2, 4}},
		},

		question699{
			para699{[][]int{[]int{1, 2}, []int{1, 3}}},
			ans699{[]int{2, 5}},
		},

		question699{
			para699{[][]int{[]int{1, 2}, []int{2, 3}, []int{6, 1}}},
			ans699{[]int{2, 5, 5}},
		},

		question699{
			para699{[][]int{[]int{100, 100}, []int{200, 100}}},
			ans699{[]int{100, 100}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 699------------------------\n")

	for _, q := range qs {
		_, p := q.ans699, q.para699
		fmt.Printf("【input】:%v       【output】:%v\n", p, fallingSquares(p.one))
	}
	fmt.Printf("\n\n\n")
}
