package leetcode

import (
	"fmt"
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type question958 struct {
	para958
	ans958
}

// para 是参数
// one 代表第一个参数
type para958 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans958 struct {
	one bool
}

func Test_Problem958(t *testing.T) {

	qs := []question958{

		{
			para958{[]int{1, 2, 3, 4, 5, 6}},
			ans958{true},
		},

		{
			para958{[]int{1, 2, 3, 4, 5, structures.NULL, 7}},
			ans958{false},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 958------------------------\n")

	for _, q := range qs {
		_, p := q.ans958, q.para958
		fmt.Printf("【input】:%v      ", p)
		root := structures.Ints2TreeNode(p.one)
		fmt.Printf("【output】:%v      \n", isCompleteTree(root))
	}
	fmt.Printf("\n\n\n")
}
