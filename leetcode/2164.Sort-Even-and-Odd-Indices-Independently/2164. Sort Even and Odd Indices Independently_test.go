package leetcode

import (
	"fmt"
	"testing"
)

type question2164 struct {
	para2164
	ans2164
}

// para 是参数
// one 代表第一个参数
type para2164 struct {
	nums []int
}

// ans 是答案
// one 代表第一个答案
type ans2164 struct {
	one []int
}

func Test_Problem1(t *testing.T) {

	qs := []question2164{
		{
			para2164{[]int{4, 1, 2, 3}},
			ans2164{[]int{2, 3, 4, 1}},
		},

		{
			para2164{[]int{2, 1}},
			ans2164{[]int{2, 1}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 2164------------------------\n")

	for _, q := range qs {
		_, p := q.ans2164, q.para2164
		fmt.Printf("【input】:%v       【output】:%v\n", p, sortEvenOdd(p.nums))
	}
	fmt.Printf("\n\n\n")
}
