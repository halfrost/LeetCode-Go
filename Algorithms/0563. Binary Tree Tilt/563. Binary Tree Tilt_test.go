package leetcode

import (
	"fmt"
	"testing"
)

type question563 struct {
	para563
	ans563
}

// para 是参数
// one 代表第一个参数
type para563 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans563 struct {
	one int
}

func Test_Problem563(t *testing.T) {

	qs := []question563{

		question563{
			para563{[]int{}},
			ans563{0},
		},

		question563{
			para563{[]int{1}},
			ans563{0},
		},

		question563{
			para563{[]int{3, 9, 20, NULL, NULL, 15, 7}},
			ans563{41},
		},

		question563{
			para563{[]int{1, 2, 3, 4, NULL, NULL, 5}},
			ans563{11},
		},

		question563{
			para563{[]int{1, 2, 3, 4, NULL, 5}},
			ans563{11},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 563------------------------\n")

	for _, q := range qs {
		_, p := q.ans563, q.para563
		fmt.Printf("【input】:%v      ", p)
		root := Ints2TreeNode(p.one)
		fmt.Printf("【output】:%v      \n", findTilt(root))
	}
	fmt.Printf("\n\n\n")
}
