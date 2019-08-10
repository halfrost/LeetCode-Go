package leetcode

import (
	"fmt"
	"testing"
)

type question1005 struct {
	para1005
	ans1005
}

// para 是参数
// one 代表第一个参数
type para1005 struct {
	one []int
	two int
}

// ans 是答案
// one 代表第一个答案
type ans1005 struct {
	one int
}

func Test_Problem1005(t *testing.T) {

	qs := []question1005{

		question1005{
			para1005{[]int{4, 2, 3}, 1},
			ans1005{5},
		},
		question1005{
			para1005{[]int{3, -1, 0, 2}, 3},
			ans1005{6},
		},

		question1005{
			para1005{[]int{2, -3, -1, 5, -4}, 2},
			ans1005{13},
		},

		question1005{
			para1005{[]int{-2, 9, 9, 8, 4}, 5},
			ans1005{32},
		},

		question1005{
			para1005{[]int{5, -7, -8, -3, 9, 5, -5, -7}, 7},
			ans1005{49},
		},

		question1005{
			para1005{[]int{5, 6, 9, -3, 3}, 2},
			ans1005{20},
		},

		question1005{
			para1005{[]int{8, -7, -3, -9, 1, 9, -6, -9, 3}, 8},
			ans1005{53},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1005------------------------\n")

	for _, q := range qs {
		_, p := q.ans1005, q.para1005
		fmt.Printf("【input】:%v       【output】:%v\n", p, largestSumAfterKNegations(p.one, p.two))
	}
	fmt.Printf("\n\n\n")
}
