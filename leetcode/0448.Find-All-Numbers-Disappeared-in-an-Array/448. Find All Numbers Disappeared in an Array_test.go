package leetcode

import (
	"fmt"
	"testing"
)

type question448 struct {
	para448
	ans448
}

// para 是参数
// one 代表第一个参数
type para448 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans448 struct {
	one []int
}

func Test_Problem448(t *testing.T) {

	qs := []question448{

		{
			para448{[]int{4, 3, 2, 7, 8, 2, 3, 1}},
			ans448{[]int{5, 6}},
		},

		{
			para448{[]int{4, 3, 2, 10, 9, 2, 3, 1, 1, 1, 1}},
			ans448{[]int{5, 6, 7, 8, 11}},
		},
		// 如需多个测试，可以复制上方元素。
	}

	fmt.Printf("------------------------Leetcode Problem 448------------------------\n")

	for _, q := range qs {
		_, p := q.ans448, q.para448
		fmt.Printf("【input】:%v       【output】:%v\n", p, findDisappearedNumbers(p.one))
	}
	fmt.Printf("\n\n\n")
}
