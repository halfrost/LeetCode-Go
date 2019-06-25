package leetcode

import (
	"fmt"
	"testing"
)

type question99 struct {
	para99
	ans99
}

// para 是参数
// one 代表第一个参数
type para99 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans99 struct {
	one []int
}

func Test_Problem99(t *testing.T) {

	qs := []question99{

		question99{
			para99{[]int{1, 3, NULL, NULL, 2}},
			ans99{[]int{3, 1, NULL, NULL, 2}},
		},

		question99{
			para99{[]int{3, 1, 4, NULL, NULL, 2}},
			ans99{[]int{2, 1, 4, NULL, NULL, 3}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 99------------------------\n")

	for _, q := range qs {
		_, p := q.ans99, q.para99
		fmt.Printf("【input】:%v      ", p)
		rootOne := Ints2TreeNode(p.one)
		recoverTree(rootOne)
		fmt.Printf("【output】:%v      \n", p)
	}
	fmt.Printf("\n\n\n")
}
