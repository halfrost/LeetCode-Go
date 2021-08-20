package leetcode

import (
	"fmt"
	"testing"
)

type question611 struct {
	para611
	ans611
}

// para 是参数
// one 代表第一个参数
type para611 struct {
	nums []int
}

// ans 是答案
// one 代表第一个答案
type ans611 struct {
	one int
}

func Test_Problem611(t *testing.T) {

	qs := []question611{

		{
			para611{[]int{2, 2, 3, 4}},
			ans611{3},
		},

		{
			para611{[]int{4, 2, 3, 4}},
			ans611{4},
		},

		{
			para611{[]int{0, 0}},
			ans611{0},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 611------------------------\n")

	for _, q := range qs {
		_, p := q.ans611, q.para611
		fmt.Printf("【input】:%v       【output】:%v\n", p, triangleNumber(p.nums))
	}
	fmt.Printf("\n\n\n")
}
