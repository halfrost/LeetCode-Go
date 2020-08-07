package leetcode

import (
	"fmt"
	"testing"
)

type question1052 struct {
	para1052
	ans1052
}

// para 是参数
// one 代表第一个参数
type para1052 struct {
	customers []int
	grumpy    []int
	x         int
}

// ans 是答案
// one 代表第一个答案
type ans1052 struct {
	one int
}

func Test_Problem1052(t *testing.T) {

	qs := []question1052{

		question1052{
			para1052{[]int{4, 10, 10}, []int{1, 1, 0}, 2},
			ans1052{24},
		},

		question1052{
			para1052{[]int{1}, []int{0}, 1},
			ans1052{1},
		},

		question1052{
			para1052{[]int{1, 0, 1, 2, 1, 1, 7, 5}, []int{0, 1, 0, 1, 0, 1, 0, 1}, 3},
			ans1052{16},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1052------------------------\n")

	for _, q := range qs {
		_, p := q.ans1052, q.para1052
		fmt.Printf("【input】:%v       【output】:%v\n", p, maxSatisfied(p.customers, p.grumpy, p.x))
	}
	fmt.Printf("\n\n\n")
}
