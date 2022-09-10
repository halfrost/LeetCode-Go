package leetcode

import (
	"fmt"
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type question530 struct {
	para530
	ans530
}

// para 是参数
// one 代表第一个参数
type para530 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans530 struct {
	one int
}

func Test_Problem530(t *testing.T) {

	qs := []question530{

		{
			para530{[]int{4, 2, 6, 1, 3}},
			ans530{1},
		},

		{
			para530{[]int{1, 0, 48, structures.NULL, structures.NULL, 12, 49}},
			ans530{1},
		},

		{
			para530{[]int{90, 69, structures.NULL, 49, 89, structures.NULL, 52}},
			ans530{1},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 530------------------------\n")

	for _, q := range qs {
		_, p := q.ans530, q.para530
		fmt.Printf("【input】:%v      ", p)
		rootOne := structures.Ints2TreeNode(p.one)
		fmt.Printf("【output】:%v      \n", getMinimumDifference(rootOne))
	}
	fmt.Printf("\n\n\n")
}
