package leetcode

import (
	"fmt"
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type question783 struct {
	para783
	ans783
}

// para 是参数
// one 代表第一个参数
type para783 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans783 struct {
	one int
}

func Test_Problem783(t *testing.T) {

	qs := []question783{

		{
			para783{[]int{4, 2, 6, 1, 3}},
			ans783{1},
		},

		{
			para783{[]int{1, 0, 48, structures.NULL, structures.NULL, 12, 49}},
			ans783{1},
		},

		{
			para783{[]int{90, 69, structures.NULL, 49, 89, structures.NULL, 52}},
			ans783{1},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 783------------------------\n")

	for _, q := range qs {
		_, p := q.ans783, q.para783
		fmt.Printf("【input】:%v      ", p)
		rootOne := structures.Ints2TreeNode(p.one)
		fmt.Printf("【output】:%v      \n", minDiffInBST(rootOne))
	}
	fmt.Printf("\n\n\n")
}
