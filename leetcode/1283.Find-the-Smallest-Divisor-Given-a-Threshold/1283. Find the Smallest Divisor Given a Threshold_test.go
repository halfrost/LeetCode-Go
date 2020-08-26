package leetcode

import (
	"fmt"
	"testing"
)

type question1283 struct {
	para1283
	ans1283
}

// para 是参数
// one 代表第一个参数
type para1283 struct {
	nums      []int
	threshold int
}

// ans 是答案
// one 代表第一个答案
type ans1283 struct {
	one int
}

func Test_Problem1283(t *testing.T) {

	qs := []question1283{

		{
			para1283{[]int{1, 2, 5, 9}, 6},
			ans1283{5},
		},

		{
			para1283{[]int{2, 3, 5, 7, 11}, 11},
			ans1283{3},
		},

		{
			para1283{[]int{19}, 5},
			ans1283{4},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1283------------------------\n")

	for _, q := range qs {
		_, p := q.ans1283, q.para1283
		fmt.Printf("【input】:%v       【output】:%v\n", p, smallestDivisor(p.nums, p.threshold))
	}
	fmt.Printf("\n\n\n")
}
