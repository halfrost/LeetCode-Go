package leetcode

import (
	"fmt"
	"testing"
)

type question107 struct {
	para107
	ans107
}

// para 是参数
// one 代表第一个参数
type para107 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans107 struct {
	one [][]int
}

func Test_Problem107(t *testing.T) {

	qs := []question107{

		question107{
			para107{[]int{}},
			ans107{[][]int{}},
		},

		question107{
			para107{[]int{1}},
			ans107{[][]int{[]int{1}}},
		},

		question107{
			para107{[]int{3, 9, 20, NULL, NULL, 15, 7}},
			ans107{[][]int{[]int{15, 7}, []int{9, 20}, []int{3}}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 107------------------------\n")

	for _, q := range qs {
		_, p := q.ans107, q.para107
		fmt.Printf("【input】:%v      ", p)
		root := Ints2TreeNode(p.one)
		fmt.Printf("【output】:%v      \n", levelOrderBottom(root))
	}
	fmt.Printf("\n\n\n")
}
