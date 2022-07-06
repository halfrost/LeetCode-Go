package leetcode

import (
	"fmt"
	"testing"
)

type question396 struct {
	para396
	ans396
}

// para 是参数
// one 代表第一个参数
type para396 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans396 struct {
	one int
}

func Test_Problem396(t *testing.T) {

	qs := []question396{
		{
			para396{[]int{4, 3, 2, 6}},
			ans396{26},
		},

		{
			para396{[]int{100}},
			ans396{0},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 396------------------------\n")

	for _, q := range qs {
		_, p := q.ans396, q.para396
		fmt.Printf("【input】:%v       【output】:%v\n", p, maxRotateFunction(p.one))
	}
	fmt.Printf("\n\n\n")
}
