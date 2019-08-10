package leetcode

import (
	"fmt"
	"testing"
)

type question993 struct {
	para993
	ans993
}

// para 是参数
// one 代表第一个参数
type para993 struct {
	one []int
	x   int
	y   int
}

// ans 是答案
// one 代表第一个答案
type ans993 struct {
	one bool
}

func Test_Problem993(t *testing.T) {

	qs := []question993{

		question993{
			para993{[]int{1, 2, 3, 4}, 4, 3},
			ans993{false},
		},

		question993{
			para993{[]int{1, 2, 3, NULL, 4, NULL, 5}, 5, 4},
			ans993{true},
		},

		question993{
			para993{[]int{1, 2, 3, NULL, 4}, 2, 3},
			ans993{false},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 993------------------------\n")

	for _, q := range qs {
		_, p := q.ans993, q.para993
		fmt.Printf("【input】:%v      ", p)
		root := Ints2TreeNode(p.one)
		fmt.Printf("【output】:%v      \n", isCousins(root, p.x, p.y))
	}
	fmt.Printf("\n\n\n")
}
