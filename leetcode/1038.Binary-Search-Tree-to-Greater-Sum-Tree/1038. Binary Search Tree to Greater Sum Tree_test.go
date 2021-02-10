package leetcode

import (
	"fmt"
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type question1038 struct {
	para1038
	ans1038
}

// para 是参数
// one 代表第一个参数
type para1038 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans1038 struct {
	one []int
}

func Test_Problem1038(t *testing.T) {

	qs := []question1038{

		{
			para1038{[]int{3, 1, structures.NULL, 0, structures.NULL, -4, structures.NULL, structures.NULL, -2}},
			ans1038{[]int{3, 4, structures.NULL, 4, structures.NULL, -2, structures.NULL, structures.NULL, 2}},
		},

		{
			para1038{[]int{2, 1}},
			ans1038{[]int{2, 3}},
		},

		{
			para1038{[]int{}},
			ans1038{[]int{}},
		},

		{
			para1038{[]int{4, 1, 6, 0, 2, 5, 7, structures.NULL, structures.NULL, structures.NULL, 3, structures.NULL, structures.NULL, structures.NULL, 8}},
			ans1038{[]int{30, 36, 21, 36, 35, 26, 15, structures.NULL, structures.NULL, structures.NULL, 33, structures.NULL, structures.NULL, structures.NULL, 8}},
		},

		{
			para1038{[]int{0, structures.NULL, 1}},
			ans1038{[]int{1, structures.NULL, 1}},
		},

		{
			para1038{[]int{1, 0, 2}},
			ans1038{[]int{3, 3, 2}},
		},

		{
			para1038{[]int{3, 2, 4, 1}},
			ans1038{[]int{7, 9, 4, 10}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1038------------------------\n")

	for _, q := range qs {
		_, p := q.ans1038, q.para1038
		fmt.Printf("【input】:%v      ", p)
		root := structures.Ints2TreeNode(p.one)
		fmt.Printf("【output】:%v      \n", structures.Tree2ints(bstToGst(root)))
	}
	fmt.Printf("\n\n\n")
}
