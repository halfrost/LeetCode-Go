package leetcode

import (
	"fmt"
	"testing"
)

type question102 struct {
	para102
	ans102
}

// para 是参数
// one 代表第一个参数
type para102 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans102 struct {
	one [][]int
}

func Test_Problem102(t *testing.T) {

	qs := []question102{

		question102{
			para102{[]int{}},
			ans102{[][]int{}},
		},

		question102{
			para102{[]int{1}},
			ans102{[][]int{[]int{1}}},
		},

		question102{
			para102{[]int{3, 9, 20, NULL, NULL, 15, 7}},
			ans102{[][]int{[]int{3}, []int{9, 20}, []int{15, 7}}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 102------------------------\n")

	for _, q := range qs {
		_, p := q.ans102, q.para102
		fmt.Printf("【input】:%v      ", p)
		root := Ints2TreeNode(p.one)
		fmt.Printf("【output】:%v      \n", levelOrder(root))
	}
	fmt.Printf("\n\n\n")
}
