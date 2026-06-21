package leetcode

import (
	"fmt"
	"sort"
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type question863 struct {
	para863
	ans863
}

// para 是参数
// one 代表第一个参数
type para863 struct {
	root   []int
	target int
	K      int
}

// ans 是答案
// one 代表第一个答案
type ans863 struct {
	one []int
}

// findTargetNode 在树中按值查找目标节点
func findTargetNode(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	}
	if node := findTargetNode(root.Left, val); node != nil {
		return node
	}
	return findTargetNode(root.Right, val)
}

func Test_Problem863(t *testing.T) {

	qs := []question863{

		{
			para863{[]int{3, 5, 1, 6, 2, 0, 8, structures.NULL, structures.NULL, 7, 4}, 5, 2},
			ans863{[]int{7, 4, 1}},
		},

		// target 在右子树，触发 rightDistance 分支
		{
			para863{[]int{3, 5, 1, 6, 2, 0, 8, structures.NULL, structures.NULL, 7, 4}, 1, 2},
			ans863{[]int{5}},
		},

		// K = 0，只返回 target 自身
		{
			para863{[]int{3, 5, 1, 6, 2, 0, 8, structures.NULL, structures.NULL, 7, 4}, 5, 0},
			ans863{[]int{5}},
		},

		// 单节点树
		{
			para863{[]int{1}, 1, 0},
			ans863{[]int{1}},
		},

		// 距离过大，无结果
		{
			para863{[]int{1, 2, 3}, 2, 5},
			ans863{[]int{}},
		},

		// target 为根节点
		{
			para863{[]int{3, 5, 1, 6, 2, 0, 8, structures.NULL, structures.NULL, 7, 4}, 3, 1},
			ans863{[]int{5, 1}},
		},

		// target 在左子树且距离父节点恰好为 K，触发 leftDistance == 0 分支
		{
			para863{[]int{3, 5, 1, 6, 2, 0, 8, structures.NULL, structures.NULL, 7, 4}, 5, 1},
			ans863{[]int{6, 2, 3}},
		},

		// target 在右子树且距离父节点恰好为 K，触发 rightDistance == 0 分支
		{
			para863{[]int{3, 5, 1, 6, 2, 0, 8, structures.NULL, structures.NULL, 7, 4}, 1, 1},
			ans863{[]int{0, 8, 3}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 863------------------------\n")

	for _, q := range qs {
		a, p := q.ans863, q.para863
		tree := structures.Ints2TreeNode(p.root)
		target := findTargetNode(tree, p.target)
		got := distanceK(tree, target, p.K)
		sort.Ints(got)
		want := append([]int{}, a.one...)
		sort.Ints(want)
		if fmt.Sprintf("%v", got) != fmt.Sprintf("%v", want) {
			t.Fatalf("distanceK(%v, %v, %v) = %v, want %v", p.root, p.target, p.K, got, a.one)
		}
		fmt.Printf("【input】:%v       【output】:%v\n", p, got)
	}
	fmt.Printf("\n\n\n")
}
