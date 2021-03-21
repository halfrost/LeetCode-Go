package leetcode

import (
	"fmt"
	"testing"
)

type question1675 struct {
	para1675
	ans1675
}

// para 是参数
// one 代表第一个参数
type para1675 struct {
	nums []int
}

// ans 是答案
// one 代表第一个答案
type ans1675 struct {
	one int
}

func Test_Problem1675(t *testing.T) {

	qs := []question1675{

		{
			para1675{[]int{1, 2, 4, 3}},
			ans1675{1},
		},

		{
			para1675{[]int{4, 1, 5, 20, 3}},
			ans1675{3},
		},

		{
			para1675{[]int{2, 10, 8}},
			ans1675{3},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1675------------------------\n")

	for _, q := range qs {
		_, p := q.ans1675, q.para1675
		fmt.Printf("【input】:%v       【output】:%v\n", p, minimumDeviation(p.nums))
	}
	fmt.Printf("\n\n\n")
}
