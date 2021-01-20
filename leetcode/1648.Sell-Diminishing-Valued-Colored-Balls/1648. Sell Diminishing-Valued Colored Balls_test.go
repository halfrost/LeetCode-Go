package leetcode

import (
	"fmt"
	"testing"
)

type question1648 struct {
	para1648
	ans1648
}

// para 是参数
// one 代表第一个参数
type para1648 struct {
	inventory []int
	orders    int
}

// ans 是答案
// one 代表第一个答案
type ans1648 struct {
	one int
}

func Test_Problem1648(t *testing.T) {

	qs := []question1648{

		{
			para1648{[]int{2, 3, 3, 4, 5}, 4},
			ans1648{16},
		},

		{
			para1648{[]int{2, 5}, 4},
			ans1648{14},
		},

		{
			para1648{[]int{3, 5}, 6},
			ans1648{19},
		},

		{
			para1648{[]int{2, 8, 4, 10, 6}, 20},
			ans1648{110},
		},

		{
			para1648{[]int{1000000000}, 1000000000},
			ans1648{21},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1648------------------------\n")

	for _, q := range qs {
		_, p := q.ans1648, q.para1648
		fmt.Printf("【input】:%v      【output】:%v      \n", p, maxProfit(p.inventory, p.orders))
	}
	fmt.Printf("\n\n\n")
}
