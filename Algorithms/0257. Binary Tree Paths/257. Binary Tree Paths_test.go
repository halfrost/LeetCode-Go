package leetcode

import (
	"fmt"
	"testing"
)

type question257 struct {
	para257
	ans257
}

// para 是参数
// one 代表第一个参数
type para257 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans257 struct {
	one []string
}

func Test_Problem257(t *testing.T) {

	qs := []question257{

		question257{
			para257{[]int{}},
			ans257{[]string{""}},
		},

		question257{
			para257{[]int{1, 2, 3, NULL, 5, NULL, NULL}},
			ans257{[]string{"1->2->5", "1->3"}},
		},

		question257{
			para257{[]int{1, 2, 3, 4, 5, 6, 7}},
			ans257{[]string{"1->2->4", "1->2->5", "1->2->4", "1->2->5", "1->3->6", "1->3->7"}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 257------------------------\n")

	for _, q := range qs {
		_, p := q.ans257, q.para257
		fmt.Printf("【input】:%v      ", p)
		root := Ints2TreeNode(p.one)
		fmt.Printf("【output】:%v      \n", binaryTreePaths(root))
	}
	fmt.Printf("\n\n\n")
}
