package leetcode

import (
	"fmt"
	"testing"
)

type question456 struct {
	para456
	ans456
}

// para 是参数
// one 代表第一个参数
type para456 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans456 struct {
	one bool
}

func Test_Problem456(t *testing.T) {

	qs := []question456{

		question456{
			para456{[]int{}},
			ans456{false},
		},

		question456{
			para456{[]int{1, 2, 3, 4}},
			ans456{false},
		},

		question456{
			para456{[]int{3, 1, 4, 2}},
			ans456{true},
		},

		question456{
			para456{[]int{-1, 3, 2, 0}},
			ans456{true},
		},

		question456{
			para456{[]int{3, 5, 0, 3, 4}},
			ans456{true},
		},
		// 如需多个测试，可以复制上方元素。
	}

	fmt.Printf("------------------------Leetcode Problem 456------------------------\n")

	for _, q := range qs {
		_, p := q.ans456, q.para456
		fmt.Printf("【input】:%v       【output】:%v\n", p, find132pattern(p.one))
	}
	fmt.Printf("\n\n\n")
}
