package leetcode

import (
	"fmt"
	"testing"
)

type question445 struct {
	para445
	ans445
}

// para 是参数
// one 代表第一个参数
type para445 struct {
	one     []int
	another []int
}

// ans 是答案
// one 代表第一个答案
type ans445 struct {
	one []int
}

func Test_Problem445(t *testing.T) {

	qs := []question445{

		question445{
			para445{[]int{}, []int{}},
			ans445{[]int{}},
		},

		question445{
			para445{[]int{1}, []int{1}},
			ans445{[]int{2}},
		},

		question445{
			para445{[]int{1, 2, 3, 4}, []int{1, 2, 3, 4}},
			ans445{[]int{2, 4, 6, 8}},
		},

		question445{
			para445{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
			ans445{[]int{2, 4, 6, 9, 0}},
		},

		question445{
			para445{[]int{1}, []int{9, 9, 9, 9, 9}},
			ans445{[]int{1, 0, 0, 0, 0, 0}},
		},

		question445{
			para445{[]int{9, 9, 9, 9, 9}, []int{1}},
			ans445{[]int{1, 0, 0, 0, 0, 0}},
		},

		question445{
			para445{[]int{2, 4, 3}, []int{5, 6, 4}},
			ans445{[]int{8, 0, 7}},
		},

		question445{
			para445{[]int{1, 8, 3}, []int{7, 1}},
			ans445{[]int{2, 5, 4}},
		},
		// 如需多个测试，可以复制上方元素。
	}

	fmt.Printf("------------------------Leetcode Problem 445------------------------\n")

	for _, q := range qs {
		_, p := q.ans445, q.para445
		fmt.Printf("【input】:%v       【output】:%v\n", p, L2s(addTwoNumbers445(S2l(p.one), S2l(p.another))))
	}
	fmt.Printf("\n\n\n")
}
