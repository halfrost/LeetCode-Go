package leetcode

import (
	"fmt"
	"testing"
)

type question1313 struct {
	para1313
	ans1313
}

// para 是参数
// one 代表第一个参数
type para1313 struct {
	nums []int
}

// ans 是答案
// one 代表第一个答案
type ans1313 struct {
	one []int
}

func Test_Problem1313(t *testing.T) {

	qs := []question1313{

		{
			para1313{[]int{1, 2, 3, 4}},
			ans1313{[]int{2, 4, 4, 4}},
		},

		{
			para1313{[]int{1, 1, 2, 3}},
			ans1313{[]int{1, 3, 3}},
		},

		{
			para1313{[]int{}},
			ans1313{[]int{}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1313------------------------\n")

	for _, q := range qs {
		_, p := q.ans1313, q.para1313
		fmt.Printf("【input】:%v      ", p)
		fmt.Printf("【output】:%v      \n", decompressRLElist(p.nums))
	}
	fmt.Printf("\n\n\n")
}
