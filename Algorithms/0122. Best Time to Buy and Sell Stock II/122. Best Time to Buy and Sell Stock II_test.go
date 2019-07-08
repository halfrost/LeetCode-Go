package leetcode

import (
	"fmt"
	"testing"
)

type question122 struct {
	para122
	ans122
}

// para 是参数
// one 代表第一个参数
type para122 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans122 struct {
	one int
}

func Test_Problem122(t *testing.T) {

	qs := []question122{

		question122{
			para122{[]int{}},
			ans122{0},
		},

		question122{
			para122{[]int{7, 1, 5, 3, 6, 4}},
			ans122{7},
		},

		question122{
			para122{[]int{7, 6, 4, 3, 1}},
			ans122{0},
		},

		question122{
			para122{[]int{1, 2, 3, 4, 5}},
			ans122{4},
		},

		question122{
			para122{[]int{1, 2, 10, 11, 12, 15}},
			ans122{14},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 122------------------------\n")

	for _, q := range qs {
		_, p := q.ans122, q.para122
		fmt.Printf("【input】:%v       【output】:%v\n", p, maxProfit122(p.one))
	}
	fmt.Printf("\n\n\n")
}
