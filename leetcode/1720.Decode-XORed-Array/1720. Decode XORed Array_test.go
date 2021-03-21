package leetcode

import (
	"fmt"
	"testing"
)

type question1720 struct {
	para1720
	ans1720
}

// para 是参数
// one 代表第一个参数
type para1720 struct {
	encoded []int
	first   int
}

// ans 是答案
// one 代表第一个答案
type ans1720 struct {
	one []int
}

func Test_Problem1720(t *testing.T) {

	qs := []question1720{

		{
			para1720{[]int{1, 2, 3}, 1},
			ans1720{[]int{1, 0, 2, 1}},
		},

		{
			para1720{[]int{6, 2, 7, 3}, 4},
			ans1720{[]int{4, 2, 0, 7, 4}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1720------------------------\n")

	for _, q := range qs {
		_, p := q.ans1720, q.para1720
		fmt.Printf("【input】:%v       【output】:%v\n", p, decode(p.encoded, p.first))
	}
	fmt.Printf("\n\n\n")
}
