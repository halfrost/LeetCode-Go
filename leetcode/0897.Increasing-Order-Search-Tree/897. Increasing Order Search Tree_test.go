package leetcode

import (
	"fmt"
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type question897 struct {
	para897
	ans897
}

// para 是参数
// one 代表第一个参数
type para897 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans897 struct {
	one []int
}

func Test_Problem897(t *testing.T) {

	qs := []question897{

		{
			para897{[]int{5, 3, 6, 2, 4, structures.NULL, 8, 1, structures.NULL, structures.NULL, structures.NULL, 7, 9}},
			ans897{[]int{1, structures.NULL, 2, structures.NULL, 3, structures.NULL, 4, structures.NULL, 5, structures.NULL, 6, structures.NULL, 7, structures.NULL, 8, structures.NULL, 9}},
		},

		{
			para897{[]int{3, 4, 4, 5, structures.NULL, structures.NULL, 5, 6, structures.NULL, structures.NULL, 6}},
			ans897{[]int{6, structures.NULL, 5, structures.NULL, 4, structures.NULL, 3, structures.NULL, 4, structures.NULL, 5, structures.NULL, 6}},
		},

		{
			para897{[]int{1, 2, 2, structures.NULL, 3, 3}},
			ans897{[]int{2, structures.NULL, 3, structures.NULL, 1, structures.NULL, 3, structures.NULL, 2}},
		},

		{
			para897{[]int{}},
			ans897{[]int{}},
		},

		{
			para897{[]int{1}},
			ans897{[]int{1}},
		},

		{
			para897{[]int{1, 2, 3}},
			ans897{[]int{2, structures.NULL, 1, structures.NULL, 3}},
		},

		{
			para897{[]int{1, 2, 2, 3, 4, 4, 3}},
			ans897{[]int{3, structures.NULL, 2, structures.NULL, 4, structures.NULL, 1, structures.NULL, 4, structures.NULL, 2, structures.NULL, 3}},
		},

		{
			para897{[]int{1, 2, 2, structures.NULL, 3, structures.NULL, 3}},
			ans897{[]int{2, structures.NULL, 3, structures.NULL, 1, structures.NULL, 2, structures.NULL, 3}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 897------------------------\n")

	for _, q := range qs {
		_, p := q.ans897, q.para897
		fmt.Printf("【input】:%v      ", p)
		rootOne := structures.Ints2TreeNode(p.one)
		fmt.Printf("【output】:%v      \n", structures.Tree2ints(increasingBST(rootOne)))
	}
	fmt.Printf("\n\n\n")
}
