package leetcode

import (
	"fmt"
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
	target []int
	K      int
}

// ans 是答案
// one 代表第一个答案
type ans863 struct {
	one []int
}

func Test_Problem863(t *testing.T) {

	qs := []question863{

		{
			para863{[]int{3, 5, 1, 6, 2, 0, 8, structures.NULL, structures.NULL, 7, 4}, []int{5}, 2},
			ans863{[]int{7, 4, 1}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 863------------------------\n")

	for _, q := range qs {
		_, p := q.ans863, q.para863
		tree, target := structures.Ints2TreeNode(p.root), structures.Ints2TreeNode(p.target)
		fmt.Printf("【input】:%v       【output】:%v\n", p, distanceK(tree, target, p.K))
	}
	fmt.Printf("\n\n\n")
}
