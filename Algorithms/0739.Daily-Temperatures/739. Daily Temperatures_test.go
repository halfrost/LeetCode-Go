package leetcode

import (
	"fmt"
	"testing"
)

type question739 struct {
	para739
	ans739
}

// para 是参数
// one 代表第一个参数
type para739 struct {
	s []int
}

// ans 是答案
// one 代表第一个答案
type ans739 struct {
	one []int
}

func Test_Problem739(t *testing.T) {

	qs := []question739{

		question739{
			para739{[]int{73, 74, 75, 71, 69, 72, 76, 73}},
			ans739{[]int{1, 1, 4, 2, 1, 1, 0, 0}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 739------------------------\n")

	for _, q := range qs {
		_, p := q.ans739, q.para739
		fmt.Printf("【input】:%v       【output】:%v\n", p, dailyTemperatures(p.s))
	}
	fmt.Printf("\n\n\n")
}
