package leetcode

import (
	"fmt"
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type question938 struct {
	para938
	ans938
}

// para 是参数
// one 代表第一个参数
type para938 struct {
	one  []int
	low  int
	high int
}

// ans 是答案
// one 代表第一个答案
type ans938 struct {
	one int
}

func Test_Problem938(t *testing.T) {

	qs := []question938{

		{
			para938{[]int{10, 5, 15, 3, 7, structures.NULL, 18}, 7, 15},
			ans938{32},
		},

		{
			para938{[]int{10, 5, 15, 3, 7, 13, 18, 1, structures.NULL, 6}, 6, 10},
			ans938{23},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 938------------------------\n")

	for _, q := range qs {
		_, p := q.ans938, q.para938
		fmt.Printf("【input】:%v      ", p)
		rootOne := structures.Ints2TreeNode(p.one)
		fmt.Printf("【output】:%v      \n", rangeSumBST(rootOne, p.low, p.high))
	}
	fmt.Printf("\n\n\n")
}
