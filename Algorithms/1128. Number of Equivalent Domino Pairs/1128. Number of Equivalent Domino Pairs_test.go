package leetcode

import (
	"fmt"
	"testing"
)

type question1128 struct {
	para1128
	ans1128
}

// para 是参数
// one 代表第一个参数
type para1128 struct {
	dominoes [][]int
}

// ans 是答案
// one 代表第一个答案
type ans1128 struct {
	one int
}

func Test_Problem1128(t *testing.T) {

	qs := []question1128{

		question1128{
			para1128{[][]int{[]int{1, 2}, []int{2, 1}, []int{3, 4}, []int{5, 6}}},
			ans1128{1},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1128------------------------\n")

	for _, q := range qs {
		_, p := q.ans1128, q.para1128
		fmt.Printf("【input】:%v       【output】:%v\n", p, numEquivDominoPairs(p.dominoes))
	}
	fmt.Printf("\n\n\n")
}
