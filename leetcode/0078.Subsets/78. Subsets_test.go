package leetcode

import (
	"fmt"
	"testing"
)

type question78 struct {
	para78
	ans78
}

// para 是参数
// one 代表第一个参数
type para78 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans78 struct {
	one [][]int
}

func Test_Problem78(t *testing.T) {

	qs := []question78{

		question78{
			para78{[]int{}},
			ans78{[][]int{[]int{}}},
		},

		question78{
			para78{[]int{1, 2, 3}},
			ans78{[][]int{[]int{}, []int{1}, []int{2}, []int{3}, []int{1, 2}, []int{2, 3}, []int{1, 3}, []int{1, 2, 3}}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 78------------------------\n")

	for _, q := range qs {
		_, p := q.ans78, q.para78
		fmt.Printf("【input】:%v       【output】:%v\n", p, subsets(p.one))
	}
	fmt.Printf("\n\n\n")
}
