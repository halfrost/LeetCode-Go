package leetcode

import (
	"fmt"
	"testing"
)

type question145 struct {
	para145
	ans145
}

// para 是参数
// one 代表第一个参数
type para145 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans145 struct {
	one []int
}

func Test_Problem145(t *testing.T) {

	qs := []question145{

		question145{
			para145{[]int{}},
			ans145{[]int{}},
		},

		question145{
			para145{[]int{1}},
			ans145{[]int{1}},
		},

		question145{
			para145{[]int{1, NULL, 2, 3}},
			ans145{[]int{1, 2, 3}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 145------------------------\n")

	for _, q := range qs {
		_, p := q.ans145, q.para145
		fmt.Printf("【input】:%v      ", p)
		root := Ints2TreeNode(p.one)
		fmt.Printf("【output】:%v      \n", postorderTraversal(root))
	}
	fmt.Printf("\n\n\n")
}
