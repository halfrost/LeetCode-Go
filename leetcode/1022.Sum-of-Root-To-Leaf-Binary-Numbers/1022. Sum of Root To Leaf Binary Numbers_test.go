package leetcode

import (
	"fmt"
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type question1022 struct {
	para1022
	ans1022
}

// para 是参数
// one 代表第一个参数
type para1022 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans1022 struct {
	one int
}

func Test_Problem1022(t *testing.T) {

	qs := []question1022{
		{
			para1022{[]int{1, 0, 1, 0, 1, 0, 1}},
			ans1022{22},
		},

		{
			para1022{[]int{0}},
			ans1022{0},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1022------------------------\n")

	for _, q := range qs {
		_, p := q.ans1022, q.para1022
		root := structures.Ints2TreeNode(p.one)
		fmt.Printf("【input】:%v       【output】:%v\n", p, sumRootToLeaf(root))
	}
	fmt.Printf("\n\n\n")
}
