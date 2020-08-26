package leetcode

import (
	"fmt"
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type question107 struct {
	para107
	ans107
}

// para 是参数
// one 代表第一个参数
type para107 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans107 struct {
	one [][]int
}

func Test_Problem107(t *testing.T) {

	qs := []question107{

		{
			para107{[]int{}},
			ans107{[][]int{}},
		},

		{
			para107{[]int{1}},
			ans107{[][]int{{1}}},
		},

		{
			para107{[]int{3, 9, 20, structures.NULL, structures.NULL, 15, 7}},
			ans107{[][]int{{15, 7}, {9, 20}, {3}}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 107------------------------\n")

	for _, q := range qs {
		_, p := q.ans107, q.para107
		fmt.Printf("【input】:%v      ", p)
		root := structures.Ints2TreeNode(p.one)
		fmt.Printf("【output】:%v      \n", levelOrderBottom(root))
	}
	fmt.Printf("\n\n\n")
}
