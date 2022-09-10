package leetcode

import (
	"fmt"
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type question1026 struct {
	para1026
	ans1026
}

// para 是参数
// one 代表第一个参数
type para1026 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans1026 struct {
	one int
}

func Test_Problem1026(t *testing.T) {

	qs := []question1026{

		{
			para1026{[]int{}},
			ans1026{0},
		},

		{
			para1026{[]int{8, 3, 10, 1, 6, structures.NULL, 14, structures.NULL, structures.NULL, 4, 7, 13}},
			ans1026{7},
		},

		{
			para1026{[]int{7, 6, 4, 3, 1}},
			ans1026{6},
		},

		{
			para1026{[]int{1, 3, 2, 8, 4, 9}},
			ans1026{8},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1026------------------------\n")

	for _, q := range qs {
		_, p := q.ans1026, q.para1026
		tree := structures.Ints2TreeNode(p.one)
		fmt.Printf("【input】:%v       【output】:%v\n", p, maxAncestorDiff(tree))
	}
	fmt.Printf("\n\n\n")
}
