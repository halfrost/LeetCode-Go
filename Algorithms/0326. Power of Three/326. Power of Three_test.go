package leetcode

import (
	"fmt"
	"testing"
)

type question326 struct {
	para326
	ans326
}

// para 是参数
// one 代表第一个参数
type para326 struct {
	one int
}

// ans 是答案
// one 代表第一个答案
type ans326 struct {
	one bool
}

func Test_Problem326(t *testing.T) {

	qs := []question326{

		question326{
			para326{27},
			ans326{true},
		},

		question326{
			para326{0},
			ans326{false},
		},

		question326{
			para326{9},
			ans326{true},
		},

		question326{
			para326{45},
			ans326{false},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 326------------------------\n")

	for _, q := range qs {
		_, p := q.ans326, q.para326
		fmt.Printf("【input】:%v       【output】:%v\n", p, isPowerOfThree(p.one))
	}
	fmt.Printf("\n\n\n")
}
