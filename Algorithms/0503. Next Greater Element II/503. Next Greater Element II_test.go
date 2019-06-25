package leetcode

import (
	"fmt"
	"testing"
)

type question503 struct {
	para503
	ans503
}

// para 是参数
// one 代表第一个参数
type para503 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans503 struct {
	one []int
}

func Test_Problem503(t *testing.T) {

	qs := []question503{

		question503{
			para503{[]int{}},
			ans503{[]int{}},
		},

		question503{
			para503{[]int{1}},
			ans503{[]int{-1}},
		},

		question503{
			para503{[]int{1, 2, 1}},
			ans503{[]int{2, -1, 2}},
		},
		// 如需多个测试，可以复制上方元素。
	}

	fmt.Printf("------------------------Leetcode Problem 503------------------------\n")

	for _, q := range qs {
		_, p := q.ans503, q.para503
		fmt.Printf("【input】:%v       【output】:%v\n", p, nextGreaterElements(p.one))
	}
	fmt.Printf("\n\n\n")
}
