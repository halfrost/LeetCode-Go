package leetcode

import (
	"fmt"
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type question222 struct {
	para222
	ans222
}

// para 是参数
// one 代表第一个参数
type para222 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans222 struct {
	one int
}

func Test_Problem222(t *testing.T) {

	qs := []question222{

		{
			para222{[]int{}},
			ans222{0},
		},

		{
			para222{[]int{1}},
			ans222{1},
		},

		{
			para222{[]int{1, 2, 3}},
			ans222{3},
		},

		{
			para222{[]int{1, 2, 3, 4, 5, 6}},
			ans222{6},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 222------------------------\n")

	for _, q := range qs {
		_, p := q.ans222, q.para222
		fmt.Printf("【input】:%v      ", p)
		rootOne := structures.Ints2TreeNode(p.one)
		fmt.Printf("【output】:%v      \n", countNodes(rootOne))
	}
	fmt.Printf("\n\n\n")
}
