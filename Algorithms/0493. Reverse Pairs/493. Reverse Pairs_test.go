package leetcode

import (
	"fmt"
	"testing"
)

type question493 struct {
	para493
	ans493
}

// para 是参数
// one 代表第一个参数
type para493 struct {
	nums []int
}

// ans 是答案
// one 代表第一个答案
type ans493 struct {
	one int
}

func Test_Problem493(t *testing.T) {

	qs := []question493{

		question493{
			para493{[]int{1, 3, 2, 3, 1}},
			ans493{2},
		},

		question493{
			para493{[]int{2, 4, 3, 5, 1}},
			ans493{3},
		},

		question493{
			para493{[]int{-5, -5}},
			ans493{1},
		},

		question493{
			para493{[]int{2147483647, 2147483647, -2147483647, -2147483647, -2147483647, 2147483647}},
			ans493{9},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 493------------------------\n")

	for _, q := range qs {
		_, p := q.ans493, q.para493
		fmt.Printf("【input】:%v       【output】:%v\n", p, reversePairs(p.nums))
	}
	fmt.Printf("\n\n\n")
}
