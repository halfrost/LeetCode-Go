package leetcode

import (
	"fmt"
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type question971 struct {
	para971
	ans971
}

// para 是参数
// one 代表第一个参数
type para971 struct {
	one    []int
	voyage []int
}

// ans 是答案
// one 代表第一个答案
type ans971 struct {
	one []int
}

func Test_Problem971(t *testing.T) {

	qs := []question971{

		{
			para971{[]int{1, 2, structures.NULL}, []int{2, 1}},
			ans971{[]int{-1}},
		},

		{
			para971{[]int{1, 2, 3}, []int{1, 3, 2}},
			ans971{[]int{1}},
		},

		{
			para971{[]int{1, 2, 3}, []int{1, 2, 3}},
			ans971{[]int{}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 971------------------------\n")

	for _, q := range qs {
		_, p := q.ans971, q.para971
		fmt.Printf("【input】:%v      ", p)
		rootOne := structures.Ints2TreeNode(p.one)
		fmt.Printf("【output】:%v      \n", flipMatchVoyage(rootOne, p.voyage))
	}
	fmt.Printf("\n\n\n")
}
