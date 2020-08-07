package leetcode

import (
	"fmt"
	"testing"
)

type question386 struct {
	para386
	ans386
}

// para 是参数
// one 代表第一个参数
type para386 struct {
	n int
}

// ans 是答案
// one 代表第一个答案
type ans386 struct {
	one []int
}

func Test_Problem386(t *testing.T) {

	qs := []question386{

		question386{
			para386{13},
			ans386{[]int{1, 10, 11, 12, 13, 2, 3, 4, 5, 6, 7, 8, 9}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 386------------------------\n")

	for _, q := range qs {
		_, p := q.ans386, q.para386
		fmt.Printf("【input】:%v       【output】:%v\n", p, lexicalOrder(p.n))
	}
	fmt.Printf("\n\n\n")
}
