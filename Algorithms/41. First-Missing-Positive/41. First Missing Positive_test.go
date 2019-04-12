package leetcode

import (
	"fmt"
	"testing"
)

type question41 struct {
	para41
	ans41
}

// para 是参数
// one 代表第一个参数
type para41 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans41 struct {
	one int
}

func Test_Problem41(t *testing.T) {

	qs := []question41{

		question41{
			para41{[]int{10, -1, 1, 2, 2, 3, 3, 3, 4, 4, 4, 4, 5, 5, 5, 5, 5, -3}},
			ans41{6},
		},

		question41{
			para41{[]int{10, -1, 8, 6, 7, 3, -2, 5, 4, 2, 1, -3}},
			ans41{9},
		},

		question41{
			para41{[]int{1}},
			ans41{2},
		},

		question41{
			para41{[]int{0, 2, 2, 1, 1}},
			ans41{3},
		},

		question41{
			para41{[]int{}},
			ans41{1},
		},

		question41{
			para41{[]int{1, 2, 0}},
			ans41{3},
		},

		question41{
			para41{[]int{3, 4, -1, 1}},
			ans41{2},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 41------------------------\n")
	for _, q := range qs {
		_, p := q.ans41, q.para41
		fmt.Printf("【input】:%v       【output】:%v\n", p, firstMissingPositive(p.one))
	}
	fmt.Printf("\n\n\n")
}
