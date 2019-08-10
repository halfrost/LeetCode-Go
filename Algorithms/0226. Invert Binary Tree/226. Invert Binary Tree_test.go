package leetcode

import (
	"fmt"
	"testing"
)

type question226 struct {
	para226
	ans226
}

// para 是参数
// one 代表第一个参数
type para226 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans226 struct {
	one []int
}

func Test_Problem226(t *testing.T) {

	qs := []question226{

		question226{
			para226{[]int{}},
			ans226{[]int{}},
		},

		question226{
			para226{[]int{1}},
			ans226{[]int{1}},
		},

		question226{
			para226{[]int{4, 2, 7, 1, 3, 6, 9}},
			ans226{[]int{4, 7, 2, 9, 6, 3, 1}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 226------------------------\n")

	for _, q := range qs {
		_, p := q.ans226, q.para226
		fmt.Printf("【input】:%v      ", p)
		root := Ints2TreeNode(p.one)
		fmt.Printf("【output】:%v      \n", Tree2Preorder(invertTree(root)))
	}
	fmt.Printf("\n\n\n")
}
