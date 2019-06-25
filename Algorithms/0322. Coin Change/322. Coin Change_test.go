package leetcode

import (
	"fmt"
	"testing"
)

type question322 struct {
	para322
	ans322
}

// para 是参数
// one 代表第一个参数
type para322 struct {
	one    []int
	amount int
}

// ans 是答案
// one 代表第一个答案
type ans322 struct {
	one int
}

func Test_Problem322(t *testing.T) {

	qs := []question322{

		question322{
			para322{[]int{186, 419, 83, 408}, 6249},
			ans322{20},
		},

		question322{
			para322{[]int{1, 2147483647}, 2},
			ans322{2},
		},

		question322{
			para322{[]int{1, 2, 5}, 11},
			ans322{3},
		},
		question322{
			para322{[]int{2}, 3},
			ans322{-1},
		},

		question322{
			para322{[]int{1}, 0},
			ans322{0},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 322------------------------\n")

	for _, q := range qs {
		_, p := q.ans322, q.para322
		fmt.Printf("【input】:%v       【output】:%v\n", p, coinChange(p.one, p.amount))
	}
	fmt.Printf("\n\n\n")
}
