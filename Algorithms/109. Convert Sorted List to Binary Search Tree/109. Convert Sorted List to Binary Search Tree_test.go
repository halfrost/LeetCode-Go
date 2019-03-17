package leetcode

import (
	"fmt"
	"testing"
)

type question109 struct {
	para109
	ans109
}

// para 是参数
// one 代表第一个参数
type para109 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans109 struct {
	one []int
}

func Test_Problem109(t *testing.T) {

	qs := []question109{

		question109{
			para109{[]int{-10, -3, 0, 5, 9}},
			ans109{[]int{0, -3, 9, -10, -9999, 5}},
		},

		question109{
			para109{[]int{-10}},
			ans109{[]int{-10}},
		},

		question109{
			para109{[]int{1, 2}},
			ans109{[]int{2, 1}},
		},

		question109{
			para109{[]int{1, 2, 3}},
			ans109{[]int{2, 1, 3}},
		},
		// question109{
		// 	para109{[]int{1, 2, 3, 4, 5}, 2, 2},
		// 	ans109{[]int{1, 2, 3, 4, 5}},
		// },

		// question109{
		// 	para109{[]int{1, 2, 3, 4, 5}, 1, 5},
		// 	ans109{[]int{5, 4, 3, 2, 1}},
		// },

		// question109{
		// 	para109{[]int{1, 2, 3, 4, 5, 6}, 3, 4},
		// 	ans109{[]int{1, 2, 4, 3, 5, 6}},
		// },

		// question109{
		// 	para109{[]int{3, 5}, 1, 2},
		// 	ans109{[]int{5, 3}},
		// },

		// question109{
		// 	para109{[]int{3}, 3, 5},
		// 	ans109{[]int{3}},
		// },
	}

	fmt.Printf("------------------------Leetcode Problem 109------------------------\n")

	for _, q := range qs {
		_, p := q.ans109, q.para109
		arr := []int{}
		T2s(sortedListToBST(S2l(p.one)), &arr)
		fmt.Printf("【input】:%v       【output】:%v\n", p, arr)
	}
	fmt.Printf("\n\n\n")
}
