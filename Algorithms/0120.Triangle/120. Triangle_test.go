package leetcode

import (
	"fmt"
	"testing"
)

type question120 struct {
	para120
	ans120
}

// para 是参数
// one 代表第一个参数
type para120 struct {
	one [][]int
}

// ans 是答案
// one 代表第一个答案
type ans120 struct {
	one int
}

func Test_Problem120(t *testing.T) {

	qs := []question120{
		question120{
			para120{[][]int{[]int{2}, []int{3, 4}, []int{6, 5, 7}, []int{4, 1, 8, 3}}},
			ans120{11},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 120------------------------\n")

	for _, q := range qs {
		_, p := q.ans120, q.para120
		fmt.Printf("【input】:%v       【output】:%v\n", p, minimumTotal(p.one))
	}
	fmt.Printf("\n\n\n")
}
