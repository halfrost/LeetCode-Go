package leetcode

import (
	"fmt"
	"testing"
)

type question1732 struct {
	para1732
	ans1732
}

// para 是参数
// one 代表第一个参数
type para1732 struct {
	gain []int
}

// ans 是答案
// one 代表第一个答案
type ans1732 struct {
	one int
}

func Test_Problem1732(t *testing.T) {

	qs := []question1732{

		{
			para1732{[]int{-5, 1, 5, 0, -7}},
			ans1732{1},
		},

		{
			para1732{[]int{-4, -3, -2, -1, 4, 3, 2}},
			ans1732{0},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1732------------------------\n")

	for _, q := range qs {
		_, p := q.ans1732, q.para1732
		fmt.Printf("【input】:%v       【output】:%v\n", p, largestAltitude(p.gain))
	}
	fmt.Printf("\n\n\n")
}
