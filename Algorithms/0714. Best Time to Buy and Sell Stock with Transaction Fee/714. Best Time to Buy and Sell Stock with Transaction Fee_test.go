package leetcode

import (
	"fmt"
	"testing"
)

type question714 struct {
	para714
	ans714
}

// para 是参数
// one 代表第一个参数
type para714 struct {
	one []int
	f   int
}

// ans 是答案
// one 代表第一个答案
type ans714 struct {
	one int
}

func Test_Problem714(t *testing.T) {

	qs := []question714{

		question714{
			para714{[]int{}, 0},
			ans714{0},
		},

		question714{
			para714{[]int{7, 1, 5, 3, 6, 4}, 0},
			ans714{5},
		},

		question714{
			para714{[]int{7, 6, 4, 3, 1}, 0},
			ans714{0},
		},

		question714{
			para714{[]int{1, 3, 2, 8, 4, 9}, 2},
			ans714{8},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 714------------------------\n")

	for _, q := range qs {
		_, p := q.ans714, q.para714
		fmt.Printf("【input】:%v       【output】:%v\n", p, maxProfit714(p.one, p.f))
	}
	fmt.Printf("\n\n\n")
}
