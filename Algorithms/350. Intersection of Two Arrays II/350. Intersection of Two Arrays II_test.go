package leetcode

import (
	"fmt"
	"testing"
)

type question350 struct {
	para350
	ans350
}

// para 是参数
// one 代表第一个参数
type para350 struct {
	one     []int
	another []int
}

// ans 是答案
// one 代表第一个答案
type ans350 struct {
	one []int
}

func Test_Problem350(t *testing.T) {

	qs := []question350{

		question350{
			para350{[]int{}, []int{}},
			ans350{[]int{}},
		},

		question350{
			para350{[]int{1}, []int{1}},
			ans350{[]int{1}},
		},

		question350{
			para350{[]int{1, 2, 3, 4}, []int{1, 2, 3, 4}},
			ans350{[]int{1, 2, 3, 4}},
		},

		question350{
			para350{[]int{1, 2, 2, 1}, []int{2, 2}},
			ans350{[]int{2, 2}},
		},

		question350{
			para350{[]int{1}, []int{9, 9, 9, 9, 9}},
			ans350{[]int{}},
		},

		question350{
			para350{[]int{4, 9, 5}, []int{9, 4, 9, 8, 4}},
			ans350{[]int{9, 4}},
		},

		question350{
			para350{[]int{4, 9, 5, 9, 4}, []int{9, 4, 9, 8, 4, 7, 9, 4, 4, 9}},
			ans350{[]int{9, 4, 9, 4}},
		},
		// 如需多个测试，可以复制上方元素。
	}

	fmt.Printf("------------------------Leetcode Problem 350------------------------\n")

	for _, q := range qs {
		_, p := q.ans350, q.para350
		fmt.Printf("【input】:%v       【output】:%v\n", p, intersect(p.one, p.another))
	}
	fmt.Printf("\n\n\n")
}
