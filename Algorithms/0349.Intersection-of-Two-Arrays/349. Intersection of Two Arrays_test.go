package leetcode

import (
	"fmt"
	"testing"
)

type question349 struct {
	para349
	ans349
}

// para 是参数
// one 代表第一个参数
type para349 struct {
	one     []int
	another []int
}

// ans 是答案
// one 代表第一个答案
type ans349 struct {
	one []int
}

func Test_Problem349(t *testing.T) {

	qs := []question349{

		question349{
			para349{[]int{}, []int{}},
			ans349{[]int{}},
		},

		question349{
			para349{[]int{1}, []int{1}},
			ans349{[]int{1}},
		},

		question349{
			para349{[]int{1, 2, 3, 4}, []int{1, 2, 3, 4}},
			ans349{[]int{1, 2, 3, 4}},
		},

		question349{
			para349{[]int{1, 2, 2, 1}, []int{2, 2}},
			ans349{[]int{2}},
		},

		question349{
			para349{[]int{1}, []int{9, 9, 9, 9, 9}},
			ans349{[]int{}},
		},

		question349{
			para349{[]int{4, 9, 5}, []int{9, 4, 9, 8, 4}},
			ans349{[]int{9, 4}},
		},
		// 如需多个测试，可以复制上方元素。
	}

	fmt.Printf("------------------------Leetcode Problem 349------------------------\n")

	for _, q := range qs {
		_, p := q.ans349, q.para349
		fmt.Printf("【input】:%v       【output】:%v\n", p, intersection(p.one, p.another))
	}
	fmt.Printf("\n\n\n")
}
