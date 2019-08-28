package leetcode

import (
	"fmt"
	"testing"
)

type question632 struct {
	para632
	ans632
}

// para 是参数
// one 代表第一个参数
type para632 struct {
	one [][]int
}

// ans 是答案
// one 代表第一个答案
type ans632 struct {
	one []int
}

func Test_Problem632(t *testing.T) {

	qs := []question632{

		question632{
			para632{[][]int{[]int{4, 10, 15, 24, 26}, []int{0, 9, 12, 20}, []int{5, 18, 22, 30}}},
			ans632{[]int{20, 24}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 632------------------------\n")

	for _, q := range qs {
		_, p := q.ans632, q.para632
		fmt.Printf("【input】:%v       【output】:%v\n", p, smallestRange(p.one))
	}
	fmt.Printf("\n\n\n")
}
