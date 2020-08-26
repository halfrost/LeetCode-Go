package leetcode

import (
	"fmt"
	"testing"
)

type question485 struct {
	para485
	ans485
}

// para 是参数
// one 代表第一个参数
type para485 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans485 struct {
	one int
}

func Test_Problem485(t *testing.T) {

	qs := []question485{

		{
			para485{[]int{1, 1, 0, 1, 1, 1}},
			ans485{3},
		},

		{
			para485{[]int{1, 1, 0, 0, 1, 0, 1, 0, 1, 1, 1, 0, 0, 1, 1, 1, 1}},
			ans485{4},
		},
		// 如需多个测试，可以复制上方元素。
	}

	fmt.Printf("------------------------Leetcode Problem 485------------------------\n")

	for _, q := range qs {
		_, p := q.ans485, q.para485
		fmt.Printf("【input】:%v       【output】:%v\n", p, findMaxConsecutiveOnes(p.one))
	}
	fmt.Printf("\n\n\n")
}
