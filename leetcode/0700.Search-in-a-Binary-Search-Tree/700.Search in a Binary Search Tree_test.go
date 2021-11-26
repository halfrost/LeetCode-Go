package leetcode

import (
	"fmt"
	"testing"
)

type question700 struct {
	para700
	ans700
}

// para 是参数
type para700 struct {
	root *TreeNode
	val  int
}

// ans 是答案
type ans700 struct {
	ans *TreeNode
}

func Test_Problem700(t *testing.T) {

	qs := []question700{
		{
			para700{&TreeNode{Val: 4,
				Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 1, Left: nil, Right: nil}, Right: &TreeNode{Val: 3, Left: nil, Right: nil}},
				Right: &TreeNode{Val: 7, Left: nil, Right: nil}},
				2},
			ans700{&TreeNode{Val: 2, Left: &TreeNode{Val: 1, Left: nil, Right: nil}, Right: &TreeNode{Val: 3, Left: nil, Right: nil}}},
		},

		{
			para700{&TreeNode{Val: 4,
				Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 1, Left: nil, Right: nil}, Right: &TreeNode{Val: 3, Left: nil, Right: nil}},
				Right: &TreeNode{Val: 7, Left: nil, Right: nil}},
				5},
			ans700{nil},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 700------------------------\n")

	for _, q := range qs {
		_, p := q.ans700, q.para700
		fmt.Printf("【input】:%v    【output】:%v\n", p, searchBST(p.root, p.val))
	}
	fmt.Printf("\n\n\n")
}
