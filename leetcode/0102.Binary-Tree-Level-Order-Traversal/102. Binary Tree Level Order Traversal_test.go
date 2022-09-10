package leetcode

import (
	"fmt"
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type question102 struct {
	para102
	ans102
}

// para 是参数
// one 代表第一个参数
type para102 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans102 struct {
	one [][]int
}

func Test_Problem102(t *testing.T) {

	qs := []question102{

		{
			para102{[]int{}},
			ans102{[][]int{}},
		},

		{
			para102{[]int{1}},
			ans102{[][]int{{1}}},
		},

		{
			para102{[]int{3, 9, 20, structures.NULL, structures.NULL, 15, 7}},
			ans102{[][]int{{3}, {9, 20}, {15, 7}}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 102------------------------\n")

	for _, q := range qs {
		_, p := q.ans102, q.para102
		fmt.Printf("【input】:%v      ", p)
		root := structures.Ints2TreeNode(p.one)
		fmt.Printf("【output】:%v      \n", levelOrder(root))
	}
	fmt.Printf("\n\n\n")
}
