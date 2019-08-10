package leetcode

import (
	"fmt"
	"testing"
)

type question94 struct {
	para94
	ans94
}

// para 是参数
// one 代表第一个参数
type para94 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans94 struct {
	one []int
}

func Test_Problem94(t *testing.T) {

	qs := []question94{

		question94{
			para94{[]int{}},
			ans94{[]int{}},
		},

		question94{
			para94{[]int{1}},
			ans94{[]int{1}},
		},

		question94{
			para94{[]int{1, NULL, 2, 3}},
			ans94{[]int{1, 2, 3}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 94------------------------\n")

	for _, q := range qs {
		_, p := q.ans94, q.para94
		fmt.Printf("【input】:%v      ", p)
		root := Ints2TreeNode(p.one)
		fmt.Printf("【output】:%v      \n", inorderTraversal(root))
	}
	fmt.Printf("\n\n\n")
}
