package leetcode

import (
	"fmt"
	"testing"
)

type question239 struct {
	para239
	ans239
}

// para 是参数
// one 代表第一个参数
type para239 struct {
	one []int
	k   int
}

// ans 是答案
// one 代表第一个答案
type ans239 struct {
	one []int
}

func Test_Problem239(t *testing.T) {

	qs := []question239{

		question239{
			para239{[]int{1, 3, -1, -3, 5, 3, 6, 7}, 3},
			ans239{[]int{3, 3, 5, 5, 6, 7}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 239------------------------\n")

	for _, q := range qs {
		_, p := q.ans239, q.para239
		fmt.Printf("【input】:%v       【output】:%v\n", p, maxSlidingWindow(p.one, p.k))
	}
	fmt.Printf("\n\n\n")
}
