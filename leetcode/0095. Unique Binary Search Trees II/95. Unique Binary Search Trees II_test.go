package leetcode

import (
	"fmt"
	"testing"
)

type question95 struct {
	para95
	ans95
}

// para 是参数
// one 代表第一个参数
type para95 struct {
	one int
}

// ans 是答案
// one 代表第一个答案
type ans95 struct {
	one []*TreeNode
}

func Test_Problem95(t *testing.T) {

	qs := []question95{

		question95{
			para95{1},
			ans95{[]*TreeNode{&TreeNode{Val: 1, Left: nil, Right: nil}}},
		},

		question95{
			para95{3},
			ans95{[]*TreeNode{
				&TreeNode{Val: 1, Left: nil, Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 2, Left: nil, Right: nil}, Right: nil}},
				&TreeNode{Val: 1, Left: nil, Right: &TreeNode{Val: 2, Left: nil, Right: &TreeNode{Val: 3, Left: nil, Right: nil}}},
				&TreeNode{Val: 3, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 1, Left: nil, Right: nil}, Right: nil}, Right: nil},
				&TreeNode{Val: 3, Left: &TreeNode{Val: 1, Left: nil, Right: &TreeNode{Val: 2, Left: nil, Right: nil}}, Right: nil},
				&TreeNode{Val: 2, Left: &TreeNode{Val: 1, Left: nil, Right: nil}, Right: &TreeNode{Val: 3, Left: nil, Right: nil}},
			}}},
	}

	fmt.Printf("------------------------Leetcode Problem 95------------------------\n")

	for _, q := range qs {
		_, p := q.ans95, q.para95
		fmt.Printf("【input】:%v  \n", p)
		trees := generateTrees(p.one)
		for _, t := range trees {
			fmt.Printf("【output】:%v\n", Tree2Preorder(t))
		}
	}
	fmt.Printf("\n\n\n")
}
