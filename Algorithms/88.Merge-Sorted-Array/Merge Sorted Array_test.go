package leetcode

import (
	"fmt"
	"testing"
)

type question88 struct {
	para88
	ans88
}

// para 是参数
// one 代表第一个参数
type para88 struct {
	one []int
	m   int
	two []int
	n   int
}

// ans 是答案
// one 代表第一个答案
type ans88 struct {
	one []int
}

func Test_Problem88(t *testing.T) {

	qs := []question88{

		// question{
		// 	para{[]int{0}, 0, []int{1}, 1},
		// 	ans{[]int{1}},
		// },
		//
		// question{
		// 	para{[]int{1, 3, 5, 7}, 4, []int{2, 4}, 2},
		// 	ans{[]int{1, 2, 3, 4}},
		// },
		//
		// question{
		// 	para{[]int{1, 3, 5, 7}, 4, []int{2, 2}, 2},
		// 	ans{[]int{1, 2, 2, 3}},
		// },

		question88{
			para88{[]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3},
			ans88{[]int{1, 2, 2, 3, 5, 6}},
		},
	}
	fmt.Printf("------------------------Leetcode Problem 88------------------------\n")
	for _, q := range qs {
		_, p := q.ans88, q.para88
		fmt.Printf("【intput】:%v,%v,%v,%v       ", p.one, p.m, p.two, p.n)
		merge(p.one, p.m, p.two, p.n)
		fmt.Printf("【output】:%v\n", p)
	}
	fmt.Printf("\n\n\n")
}
