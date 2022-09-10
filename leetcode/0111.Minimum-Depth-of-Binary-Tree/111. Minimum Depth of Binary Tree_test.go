package leetcode

import (
	"fmt"
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type question111 struct {
	para111
	ans111
}

// para 是参数
// one 代表第一个参数
type para111 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans111 struct {
	one int
}

func Test_Problem111(t *testing.T) {

	qs := []question111{

		{
			para111{[]int{}},
			ans111{0},
		},

		{
			para111{[]int{1}},
			ans111{1},
		},

		{
			para111{[]int{3, 9, 20, structures.NULL, structures.NULL, 15, 7}},
			ans111{2},
		},

		{
			para111{[]int{1, 2}},
			ans111{2},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 111------------------------\n")

	for _, q := range qs {
		_, p := q.ans111, q.para111
		fmt.Printf("【input】:%v      ", p)
		root := structures.Ints2TreeNode(p.one)
		fmt.Printf("【output】:%v      \n", minDepth(root))
	}
	fmt.Printf("\n\n\n")
}
