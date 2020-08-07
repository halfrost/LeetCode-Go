package leetcode

import (
	"fmt"
	"testing"
)

type question354 struct {
	para354
	ans354
}

// para 是参数
// one 代表第一个参数
type para354 struct {
	envelopes [][]int
}

// ans 是答案
// one 代表第一个答案
type ans354 struct {
	one int
}

func Test_Problem354(t *testing.T) {

	qs := []question354{

		question354{
			para354{[][]int{[]int{5, 4}, []int{6, 4}, []int{6, 7}, []int{2, 3}}},
			ans354{3},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 354------------------------\n")

	for _, q := range qs {
		_, p := q.ans354, q.para354
		fmt.Printf("【input】:%v       【output】:%v\n", p, maxEnvelopes(p.envelopes))
	}
	fmt.Printf("\n\n\n")
}
