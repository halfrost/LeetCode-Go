package leetcode

import (
	"fmt"
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type question623 struct {
	para623
	ans623
}

// para 是参数
// one 代表第一个参数
type para623 struct {
	one []int
	v   int
	d   int
}

// ans 是答案
// one 代表第一个答案
type ans623 struct {
	one []int
}

func Test_Problem623(t *testing.T) {

	qs := []question623{

		{
			para623{[]int{4, 2, 6, 3, 1, 5, structures.NULL}, 1, 2},
			ans623{[]int{4, 1, 1, 2, structures.NULL, structures.NULL, 6, 3, 1, 5, structures.NULL}},
		},

		{
			para623{[]int{4, 2, structures.NULL, 3, 1}, 1, 3},
			ans623{[]int{4, 2, structures.NULL, 1, 1, 3, structures.NULL, structures.NULL, 1}},
		},

		{
			para623{[]int{1, 2, 3, 4}, 5, 4},
			ans623{[]int{1, 2, 3, 4, structures.NULL, structures.NULL, structures.NULL, 5, 5}},
		},

		{
			para623{[]int{4, 2, 6, 3, 1, 5}, 1, 3},
			ans623{[]int{4, 2, 6, 1, 1, 1, 1, 3, structures.NULL, structures.NULL, 1, 5}},
		},

		{
			para623{[]int{4, 2, 6, 3, 1, 5}, 1, 1},
			ans623{[]int{1, 4, structures.NULL, 2, 6, 3, 1, 5}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 623------------------------\n")

	for _, q := range qs {
		_, p := q.ans623, q.para623
		fmt.Printf("【input】:%v      ", p)
		root := structures.Ints2TreeNode(p.one)
		fmt.Printf("【output】:%v      \n", structures.Tree2Preorder(addOneRow(root, p.v, p.d)))
	}
	fmt.Printf("\n\n\n")
}
