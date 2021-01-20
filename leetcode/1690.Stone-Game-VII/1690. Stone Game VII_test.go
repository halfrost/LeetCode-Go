package leetcode

import (
	"fmt"
	"testing"
)

type question1690 struct {
	para1690
	ans1690
}

// para 是参数
// one 代表第一个参数
type para1690 struct {
	stones []int
}

// ans 是答案
// one 代表第一个答案
type ans1690 struct {
	one int
}

func Test_Problem1690(t *testing.T) {

	qs := []question1690{

		{
			para1690{[]int{5, 3, 1, 4, 2}},
			ans1690{6},
		},

		{
			para1690{[]int{7, 90, 5, 1, 100, 10, 10, 2}},
			ans1690{122},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1690------------------------\n")

	for _, q := range qs {
		_, p := q.ans1690, q.para1690
		fmt.Printf("【input】:%v       【output】:%v\n", p, stoneGameVII(p.stones))
	}
	fmt.Printf("\n\n\n")
}
