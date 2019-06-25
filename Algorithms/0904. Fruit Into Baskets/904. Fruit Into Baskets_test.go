package leetcode

import (
	"fmt"
	"testing"
)

type question904 struct {
	para904
	ans904
}

// para 是参数
// one 代表第一个参数
type para904 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans904 struct {
	one int
}

func Test_Problem904(t *testing.T) {

	qs := []question904{

		question904{
			para904{[]int{1, 1}},
			ans904{2},
		},

		question904{
			para904{[]int{1, 2, 1}},
			ans904{3},
		},
		question904{
			para904{[]int{0, 1, 2, 2}},
			ans904{3},
		},

		question904{
			para904{[]int{1, 2, 3, 2, 2}},
			ans904{4},
		},

		question904{
			para904{[]int{3, 3, 3, 1, 2, 1, 1, 2, 3, 3, 4}},
			ans904{5},
		},

		question904{
			para904{[]int{4, 5}},
			ans904{2},
		},

		question904{
			para904{[]int{1}},
			ans904{1},
		},

		question904{
			para904{[]int{}},
			ans904{},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 904------------------------\n")

	for _, q := range qs {
		_, p := q.ans904, q.para904
		fmt.Printf("【input】:%v       【output】:%v\n", p, totalFruit(p.one))
	}
	fmt.Printf("\n\n\n")
}
