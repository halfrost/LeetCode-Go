package leetcode

import (
	"fmt"
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type question617 struct {
	para617
	ans617
}

// para 是参数
// one 代表第一个参数
type para617 struct {
	one     []int
	another []int
}

// ans 是答案
// one 代表第一个答案
type ans617 struct {
	one []int
}

func Test_Problem617(t *testing.T) {

	qs := []question617{

		{
			para617{[]int{}, []int{}},
			ans617{[]int{}},
		},

		{
			para617{[]int{}, []int{1}},
			ans617{[]int{1}},
		},

		{
			para617{[]int{1, 3, 2, 5}, []int{2, 1, 3, structures.NULL, 4, structures.NULL, 7}},
			ans617{[]int{3, 4, 5, 5, 4, structures.NULL, 7}},
		},

		{
			para617{[]int{1}, []int{1, 2}},
			ans617{[]int{2, 2}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 617------------------------\n")

	for _, q := range qs {
		_, p := q.ans617, q.para617
		fmt.Printf("【input】:%v      ", p)
		root1 := structures.Ints2TreeNode(p.one)
		root2 := structures.Ints2TreeNode(p.another)
		fmt.Printf("【output】:%v      \n", mergeTrees(root1, root2))
	}
	fmt.Printf("\n\n\n")
}
