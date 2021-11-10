package leetcode

import (
	"fmt"
	"testing"
)

type question495 struct {
	para495
	ans495
}

// para 是参数
type para495 struct {
	timeSeries []int
	duration   int
}

// ans 是答案
type ans495 struct {
	ans int
}

func Test_Problem495(t *testing.T) {

	qs := []question495{

		{
			para495{[]int{1, 4}, 2},
			ans495{4},
		},

		{
			para495{[]int{1, 2}, 2},
			ans495{3},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 495------------------------\n")

	for _, q := range qs {
		_, p := q.ans495, q.para495
		fmt.Printf("【input】:%v       【output】:%v\n", p, findPoisonedDuration(p.timeSeries, p.duration))
	}
	fmt.Printf("\n\n\n")
}
