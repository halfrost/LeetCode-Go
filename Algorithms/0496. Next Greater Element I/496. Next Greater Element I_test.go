package leetcode

import (
	"fmt"
	"testing"
)

type question496 struct {
	para496
	ans496
}

// para 是参数
// one 代表第一个参数
type para496 struct {
	one     []int
	another []int
}

// ans 是答案
// one 代表第一个答案
type ans496 struct {
	one []int
}

func Test_Problem496(t *testing.T) {

	qs := []question496{

		question496{
			para496{[]int{4, 1, 2}, []int{1, 3, 4, 2}},
			ans496{[]int{-1, 3, -1}},
		},

		question496{
			para496{[]int{2, 4}, []int{1, 2, 3, 4}},
			ans496{[]int{3, -1}},
		},

		// 如需多个测试，可以复制上方元素。
	}

	fmt.Printf("------------------------Leetcode Problem 496------------------------\n")

	for _, q := range qs {
		_, p := q.ans496, q.para496
		fmt.Printf("【input】:%v       【output】:%v\n", p, nextGreaterElement(p.one, p.another))
	}
	fmt.Printf("\n\n\n")
}
