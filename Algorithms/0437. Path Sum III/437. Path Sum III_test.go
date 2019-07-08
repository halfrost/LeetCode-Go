package leetcode

import (
	"fmt"
	"testing"
)

type question437 struct {
	para437
	ans437
}

// para 是参数
// one 代表第一个参数
type para437 struct {
	one []int
	sum int
}

// ans 是答案
// one 代表第一个答案
type ans437 struct {
	one int
}

func Test_Problem437(t *testing.T) {

	qs := []question437{
		question437{
			para437{[]int{}, 0},
			ans437{0},
		},

		question437{
			para437{[]int{10, 5, -3, 3, 2, NULL, 11, 3, -2, NULL, 1}, 8},
			ans437{3},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 437------------------------\n")

	for _, q := range qs {
		_, p := q.ans437, q.para437
		fmt.Printf("【input】:%v      ", p)
		root := Ints2TreeNode(p.one)
		fmt.Printf("【output】:%v      \n", pathSumIII(root, p.sum))
	}
	fmt.Printf("\n\n\n")
}
