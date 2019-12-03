package leetcode

import (
	"fmt"
	"testing"
)

type question786 struct {
	para786
	ans786
}

// para 是参数
// one 代表第一个参数
type para786 struct {
	A []int
	K int
}

// ans 是答案
// one 代表第一个答案
type ans786 struct {
	one []int
}

func Test_Problem786(t *testing.T) {

	qs := []question786{

		question786{
			para786{[]int{1, 2, 3, 5}, 3},
			ans786{[]int{2, 5}},
		},

		question786{
			para786{[]int{1, 7}, 1},
			ans786{[]int{1, 7}},
		},

		question786{
			para786{[]int{1, 2}, 1},
			ans786{[]int{1, 2}},
		},

		question786{
			para786{[]int{1, 2, 3, 5, 7}, 6},
			ans786{[]int{3, 7}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 786------------------------\n")

	for _, q := range qs {
		_, p := q.ans786, q.para786
		fmt.Printf("【input】:%v       【output】:%v\n", p, kthSmallestPrimeFraction(p.A, p.K))
	}
	fmt.Printf("\n\n\n")
}
