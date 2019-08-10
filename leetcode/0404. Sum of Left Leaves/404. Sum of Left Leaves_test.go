package leetcode

import (
	"fmt"
	"testing"
)

type question404 struct {
	para404
	ans404
}

// para 是参数
// one 代表第一个参数
type para404 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans404 struct {
	one int
}

func Test_Problem404(t *testing.T) {

	qs := []question404{

		question404{
			para404{[]int{}},
			ans404{0},
		},

		question404{
			para404{[]int{3, 9, 20, NULL, NULL, 15, 7}},
			ans404{24},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 404------------------------\n")

	for _, q := range qs {
		_, p := q.ans404, q.para404
		fmt.Printf("【input】:%v      ", p)
		root := Ints2TreeNode(p.one)
		fmt.Printf("【output】:%v      \n", sumOfLeftLeaves(root))
	}
	fmt.Printf("\n\n\n")
}
