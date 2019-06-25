package leetcode

import (
	"fmt"
	"testing"
)

type question628 struct {
	para628
	ans628
}

// para 是参数
// one 代表第一个参数
type para628 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans628 struct {
	one int
}

func Test_Problem628(t *testing.T) {

	qs := []question628{

		question628{
			para628{[]int{3, -1, 4}},
			ans628{-12},
		},

		question628{
			para628{[]int{1, 2}},
			ans628{2},
		},

		question628{
			para628{[]int{1, 2, 3}},
			ans628{6},
		},

		question628{
			para628{[]int{1, 2, 3, 4}},
			ans628{24},
		},

		question628{
			para628{[]int{-2}},
			ans628{-2},
		},

		question628{
			para628{[]int{0}},
			ans628{0},
		},

		question628{
			para628{[]int{2, 3, -2, 4}},
			ans628{-24},
		},

		question628{
			para628{[]int{-2, 0, -1}},
			ans628{0},
		},

		question628{
			para628{[]int{-2, 0, -1, 2, 3, 1, 10}},
			ans628{60},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 628------------------------\n")

	for _, q := range qs {
		_, p := q.ans628, q.para628
		fmt.Printf("【input】:%v       【output】:%v\n", p, maximumProduct(p.one))
	}
	fmt.Printf("\n\n\n")
}
