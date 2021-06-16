package leetcode

import (
	"fmt"
	"testing"
)

type question877 struct {
	para877
	ans877
}

// para 是参数
// one 代表第一个参数
type para877 struct {
	piles []int
}

// ans 是答案
// one 代表第一个答案
type ans877 struct {
	one bool
}

func Test_Problem877(t *testing.T) {

	qs := []question877{

		{
			para877{[]int{5, 3, 4, 5}},
			ans877{true},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 877------------------------\n")

	for _, q := range qs {
		_, p := q.ans877, q.para877
		fmt.Printf("【input】:%v       【output】:%v\n", p, stoneGame(p.piles))
	}
	fmt.Printf("\n\n\n")
}
