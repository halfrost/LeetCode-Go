package leetcode

import (
	"fmt"
	"testing"
)

type question477 struct {
	para477
	ans477
}

// para 是参数
// one 代表第一个参数
type para477 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans477 struct {
	one int
}

func Test_Problem477(t *testing.T) {

	qs := []question477{

		question477{
			para477{[]int{4, 14, 2}},
			ans477{6},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 477------------------------\n")

	for _, q := range qs {
		_, p := q.ans477, q.para477
		fmt.Printf("【input】:%v       【output】:%v\n", p, totalHammingDistance(p.one))
	}
	fmt.Printf("\n\n\n")
}
