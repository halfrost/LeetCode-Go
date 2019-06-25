package leetcode

import (
	"fmt"
	"testing"
)

type question54 struct {
	para54
	ans54
}

// para 是参数
// one 代表第一个参数
type para54 struct {
	one [][]int
}

// ans 是答案
// one 代表第一个答案
type ans54 struct {
	one []int
}

func Test_Problem54(t *testing.T) {

	qs := []question54{

		question54{
			para54{[][]int{[]int{3}, []int{2}}},
			ans54{[]int{3, 2}},
		},

		question54{
			para54{[][]int{[]int{2, 3}}},
			ans54{[]int{2, 3}},
		},

		question54{
			para54{[][]int{[]int{1}}},
			ans54{[]int{1}},
		},

		question54{
			para54{[][]int{[]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9}}},
			ans54{[]int{1, 2, 3, 6, 9, 8, 7, 4, 5}},
		},
		question54{
			para54{[][]int{[]int{1, 2, 3, 4}, []int{5, 6, 7, 8}, []int{9, 10, 11, 12}}},
			ans54{[]int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 54------------------------\n")

	for _, q := range qs {
		_, p := q.ans54, q.para54
		fmt.Printf("【input】:%v       【output】:%v\n", p, spiralOrder(p.one))
	}
	fmt.Printf("\n\n\n")
}
