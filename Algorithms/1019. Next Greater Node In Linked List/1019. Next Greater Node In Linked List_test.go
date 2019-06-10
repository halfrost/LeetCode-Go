package leetcode

import (
	"fmt"
	"testing"
)

type question1019 struct {
	para1019
	ans1019
}

// para 是参数
// one 代表第一个参数
type para1019 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans1019 struct {
	one []int
}

func Test_Problem1019(t *testing.T) {

	qs := []question1019{

		question1019{
			para1019{[]int{2, 1, 5}},
			ans1019{[]int{5, 5, 0}},
		},

		question1019{
			para1019{[]int{2, 7, 4, 3, 5}},
			ans1019{[]int{7, 0, 5, 5, 0}},
		},

		question1019{
			para1019{[]int{1, 7, 5, 1, 9, 2, 5, 1}},
			ans1019{[]int{7, 9, 9, 9, 0, 5, 0, 0}},
		},

		question1019{
			para1019{[]int{1, 7, 5, 1, 9, 2, 5, 6, 7, 8, 1}},
			ans1019{[]int{7, 9, 9, 9, 0, 5, 6, 7, 8, 0, 0}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1019------------------------\n")

	for _, q := range qs {
		_, p := q.ans1019, q.para1019
		fmt.Printf("【input】:%v       【output】:%v\n", p, nextLargerNodes(S2l(p.one)))
	}
	fmt.Printf("\n\n\n")
}
