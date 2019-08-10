package leetcode

import (
	"fmt"
	"testing"
)

type question454 struct {
	para454
	ans454
}

// para 是参数
// one 代表第一个参数
type para454 struct {
	a []int
	b []int
	c []int
	d []int
}

// ans 是答案
// one 代表第一个答案
type ans454 struct {
	one int
}

func Test_Problem454(t *testing.T) {

	qs := []question454{

		question454{
			para454{[]int{1, 2}, []int{-2, -1}, []int{-1, 2}, []int{0, 2}},
			ans454{2},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 454------------------------\n")

	for _, q := range qs {
		_, p := q.ans454, q.para454
		fmt.Printf("【input】:%v       【output】:%v\n", p, fourSumCount(p.a, p.b, p.c, p.d))
	}
	fmt.Printf("\n\n\n")
}
