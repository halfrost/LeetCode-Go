package leetcode

import (
	"fmt"
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type question543 struct {
	para543
	ans543
}

// para 是参数
// one 代表第一个参数
type para543 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans543 struct {
	one int
}

func Test_Problem543(t *testing.T) {

	qs := []question543{

		{
			para543{[]int{1, 2, 3, 4, 5}},
			ans543{3},
		},

		{
			para543{[]int{1, 2}},
			ans543{1},
		},

		{
			para543{[]int{4, -7, -3, structures.NULL, structures.NULL, -9, -3, 9, -7, -4, structures.NULL, 6, structures.NULL, -6, -6, structures.NULL, structures.NULL, 0, 6, 5, structures.NULL, 9, structures.NULL, structures.NULL, -1, -4, structures.NULL, structures.NULL, structures.NULL, -2}},
			ans543{8},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 543------------------------\n")

	for _, q := range qs {
		_, p := q.ans543, q.para543
		fmt.Printf("【input】:%v      ", p)
		root := structures.Ints2TreeNode(p.one)
		fmt.Printf("【output】:%v      \n", diameterOfBinaryTree(root))
	}
	fmt.Printf("\n\n\n")
}
