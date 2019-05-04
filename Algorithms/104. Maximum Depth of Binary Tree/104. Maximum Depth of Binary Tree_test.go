package leetcode

import (
	"fmt"
	"testing"
)

type question104 struct {
	para104
	ans104
}

// para 是参数
// one 代表第一个参数
type para104 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans104 struct {
	one int
}

func Test_Problem104(t *testing.T) {

	qs := []question104{

		question104{
			para104{[]int{}},
			ans104{0},
		},

		question104{
			para104{[]int{3, 9, 20, NULL, NULL, 15, 7}},
			ans104{3},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 104------------------------\n")

	for _, q := range qs {
		_, p := q.ans104, q.para104
		fmt.Printf("【input】:%v      ", p)
		root := Ints2TreeNode(p.one)
		fmt.Printf("【output】:%v      \n", maxDepth(root))
	}
	fmt.Printf("\n\n\n")
}
