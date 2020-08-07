package leetcode

import (
	"fmt"
	"testing"
)

type question1207 struct {
	para1207
	ans1207
}

// para 是参数
// one 代表第一个参数
type para1207 struct {
	arr []int
}

// ans 是答案
// one 代表第一个答案
type ans1207 struct {
	one bool
}

func Test_Problem1207(t *testing.T) {

	qs := []question1207{

		question1207{
			para1207{[]int{1, 2, 2, 1, 1, 3}},
			ans1207{true},
		},

		question1207{
			para1207{[]int{1, 2}},
			ans1207{false},
		},

		question1207{
			para1207{[]int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0}},
			ans1207{true},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1207------------------------\n")

	for _, q := range qs {
		_, p := q.ans1207, q.para1207
		fmt.Printf("【input】:%v       【output】:%v\n", p, uniqueOccurrences(p.arr))
	}
	fmt.Printf("\n\n\n")
}
