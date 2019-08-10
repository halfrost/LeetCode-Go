package leetcode

import (
	"fmt"
	"testing"
)

type question129 struct {
	para129
	ans129
}

// para 是参数
// one 代表第一个参数
type para129 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans129 struct {
	one int
}

func Test_Problem129(t *testing.T) {

	qs := []question129{

		question129{
			para129{[]int{}},
			ans129{0},
		},

		question129{
			para129{[]int{1, 2, 3}},
			ans129{25},
		},

		question129{
			para129{[]int{4, 9, 0, 5, 1}},
			ans129{1026},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 129------------------------\n")

	for _, q := range qs {
		_, p := q.ans129, q.para129
		fmt.Printf("【input】:%v      ", p)
		root := Ints2TreeNode(p.one)
		fmt.Printf("【output】:%v      \n", sumNumbers(root))
	}
	fmt.Printf("\n\n\n")
}
