package leetcode

import (
	"fmt"
	"testing"
)

type question124 struct {
	para124
	ans124
}

// para 是参数
// one 代表第一个参数
type para124 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans124 struct {
	one int
}

func Test_Problem124(t *testing.T) {

	qs := []question124{

		question124{
			para124{[]int{}},
			ans124{0},
		},

		question124{
			para124{[]int{1}},
			ans124{1},
		},

		question124{
			para124{[]int{1, 2, 3}},
			ans124{6},
		},

		question124{
			para124{[]int{-10, 9, 20, NULL, NULL, 15, 7}},
			ans124{42},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 124------------------------\n")

	for _, q := range qs {
		_, p := q.ans124, q.para124
		fmt.Printf("【input】:%v      ", p)
		root := Ints2TreeNode(p.one)
		fmt.Printf("【output】:%v      \n", maxPathSum(root))
	}
	fmt.Printf("\n\n\n")
}
