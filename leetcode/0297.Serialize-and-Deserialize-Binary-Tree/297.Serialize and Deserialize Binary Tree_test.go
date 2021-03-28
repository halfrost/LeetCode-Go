package leetcode

import (
	"fmt"
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type question297 struct {
	para297
	ans297
}

// para 是参数
// one 代表第一个参数
type para297 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans297 struct {
	one []int
}

func Test_Problem297(t *testing.T) {
	qs := []question297{
		{
			para297{[]int{}},
			ans297{[]int{}},
		},
		{
			para297{[]int{1,2,3,-1,-1,4,5}},
			ans297{[]int{1,2,3,-1,-1,4,5}},
		},
		{
			para297{[]int{1,2}},
			ans297{[]int{1,2}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 297------------------------\n")

	for _, q := range qs {
		_, p := q.ans297, q.para297
		fmt.Printf("【input】:%v      ", p)
		root := structures.Ints2TreeNode(p.one)

		tree297 := Constructor()
		serialized := tree297.serialize(root)
		fmt.Printf("【output】:%v      \n", structures.Tree2Preorder(tree297.deserialize(serialized)))
	}
	fmt.Printf("\n\n\n")
}