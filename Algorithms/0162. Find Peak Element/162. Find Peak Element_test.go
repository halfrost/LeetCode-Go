package leetcode

import (
	"fmt"
	"testing"
)

type question162 struct {
	para162
	ans162
}

// para 是参数
// one 代表第一个参数
type para162 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans162 struct {
	one int
}

func Test_Problem162(t *testing.T) {

	qs := []question162{

		question162{
			para162{[]int{2, 1, 2}},
			ans162{0},
		},

		question162{
			para162{[]int{3, 2, 1}},
			ans162{0},
		},

		question162{
			para162{[]int{1, 2}},
			ans162{1},
		},

		question162{
			para162{[]int{2, 1}},
			ans162{0},
		},

		question162{
			para162{[]int{1}},
			ans162{0},
		},

		question162{
			para162{[]int{1, 2, 3, 1}},
			ans162{2},
		},

		question162{
			para162{[]int{1, 2, 1, 3, 5, 6, 4}},
			ans162{5},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 162------------------------\n")

	for _, q := range qs {
		_, p := q.ans162, q.para162
		fmt.Printf("【input】:%v       【output】:%v\n", p, findPeakElement(p.one))
	}
	fmt.Printf("\n\n\n")
}
