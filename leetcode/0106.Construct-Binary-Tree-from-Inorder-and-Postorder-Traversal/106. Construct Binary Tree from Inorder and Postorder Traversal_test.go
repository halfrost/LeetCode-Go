package leetcode

import (
	"fmt"
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type question106 struct {
	para106
	ans106
}

// para 是参数
// one 代表第一个参数
type para106 struct {
	inorder   []int
	postorder []int
}

// ans 是答案
// one 代表第一个答案
type ans106 struct {
	one []int
}

func Test_Problem106(t *testing.T) {

	qs := []question106{

		{
			para106{[]int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3}},
			ans106{[]int{3, 9, 20, structures.NULL, structures.NULL, 15, 7}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 106------------------------\n")

	for _, q := range qs {
		_, p := q.ans106, q.para106
		fmt.Printf("【input】:%v      ", p)
		fmt.Printf("【output】:%v      \n", structures.Tree2ints(buildTree(p.inorder, p.postorder)))
	}
	fmt.Printf("\n\n\n")
}
