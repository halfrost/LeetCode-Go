package leetcode

import (
	"fmt"
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type question1302 struct {
	para1302
	ans1302
}

// para 是参数
// one 代表第一个参数
type para1302 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans1302 struct {
	one int
}

func Test_Problem1302(t *testing.T) {

	qs := []question1302{

		{
			para1302{[]int{1, 2, 3, 4, 5, structures.NULL, 6, 7, structures.NULL, structures.NULL, structures.NULL, structures.NULL, 8}},
			ans1302{15},
		},

		{
			para1302{[]int{}},
			ans1302{0},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1302------------------------\n")

	for _, q := range qs {
		_, p := q.ans1302, q.para1302
		fmt.Printf("【input】:%v       【output】:%v\n", p, deepestLeavesSum(structures.Ints2TreeNode(p.one)))
	}
	fmt.Printf("\n\n\n")
}
