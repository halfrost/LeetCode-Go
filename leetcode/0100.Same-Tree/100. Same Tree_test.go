package leetcode

import (
	"fmt"
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type question100 struct {
	para100
	ans100
}

// para 是参数
// one 代表第一个参数
type para100 struct {
	one []int
	two []int
}

// ans 是答案
// one 代表第一个答案
type ans100 struct {
	one bool
}

func Test_Problem100(t *testing.T) {

	qs := []question100{

		{
			para100{[]int{}, []int{}},
			ans100{true},
		},

		{
			para100{[]int{}, []int{1}},
			ans100{false},
		},

		{
			para100{[]int{1}, []int{1}},
			ans100{true},
		},

		{
			para100{[]int{1, 2, 3}, []int{1, 2, 3}},
			ans100{true},
		},

		{
			para100{[]int{1, 2}, []int{1, structures.NULL, 2}},
			ans100{false},
		},

		{
			para100{[]int{1, 2, 1}, []int{1, 1, 2}},
			ans100{false},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 100------------------------\n")

	for _, q := range qs {
		_, p := q.ans100, q.para100
		fmt.Printf("【input】:%v      ", p)
		rootOne := structures.Ints2TreeNode(p.one)
		rootTwo := structures.Ints2TreeNode(p.two)
		fmt.Printf("【output】:%v      \n", isSameTree(rootOne, rootTwo))
	}
	fmt.Printf("\n\n\n")
}
