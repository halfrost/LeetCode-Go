package leetcode

import (
	"fmt"
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type question1123 struct {
	para1123
	ans1123
}

// para 是参数
// one 代表第一个参数
type para1123 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans1123 struct {
	one []int
}

func Test_Problem1123(t *testing.T) {

	qs := []question1123{

		{
			para1123{[]int{}},
			ans1123{[]int{}},
		},

		{
			para1123{[]int{1}},
			ans1123{[]int{1}},
		},

		{
			para1123{[]int{1, 2, 3, 4}},
			ans1123{[]int{4}},
		},

		{
			para1123{[]int{1, 2, 3}},
			ans1123{[]int{1, 2, 3}},
		},

		{
			para1123{[]int{1, 2, 3, 4, 5}},
			ans1123{[]int{2, 4, 5}},
		},

		{
			para1123{[]int{1, 2, structures.NULL, 3, 4, structures.NULL, 6, structures.NULL, 5}},
			ans1123{[]int{2, 3, 4, structures.NULL, 6, structures.NULL, 5}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1123------------------------\n")

	for _, q := range qs {
		_, p := q.ans1123, q.para1123
		fmt.Printf("【input】:%v      ", p)
		root := structures.Ints2TreeNode(p.one)
		fmt.Printf("【output】:%v      \n", structures.Tree2ints(lcaDeepestLeaves(root)))
	}
	fmt.Printf("\n\n\n")
}
