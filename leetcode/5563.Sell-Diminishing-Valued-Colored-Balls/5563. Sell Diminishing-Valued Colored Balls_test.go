package leetcode

import (
	"fmt"
	"testing"
)

type question5563 struct {
	para5563
	ans5563
}

// para 是参数
// one 代表第一个参数
type para5563 struct {
	inventory []int
	orders    int
}

// ans 是答案
// one 代表第一个答案
type ans5563 struct {
	one int
}

func Test_Problem5563(t *testing.T) {

	qs := []question5563{

		{
			para5563{[]int{2, 5}, 4},
			ans5563{14},
		},

		{
			para5563{[]int{3, 5}, 6},
			ans5563{19},
		},

		{
			para5563{[]int{2, 8, 4, 10, 6}, 20},
			ans5563{110},
		},

		{
			para5563{[]int{1000000000}, 1000000000},
			ans5563{21},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 5563------------------------\n")

	for _, q := range qs {
		_, p := q.ans5563, q.para5563
		fmt.Printf("【input】:%v      【output】:%v      \n", p, maxProfit(p.inventory, p.orders))
	}
	fmt.Printf("\n\n\n")
}
