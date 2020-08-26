package leetcode

import (
	"fmt"
	"testing"
)

type question985 struct {
	para985
	ans985
}

// para 是参数
// one 代表第一个参数
type para985 struct {
	A       []int
	queries [][]int
}

// ans 是答案
// one 代表第一个答案
type ans985 struct {
	one []int
}

func Test_Problem985(t *testing.T) {

	qs := []question985{

		{
			para985{[]int{1, 2, 3, 4}, [][]int{{1, 0}, {-3, 1}, {-4, 0}, {2, 3}}},
			ans985{[]int{8, 6, 2, 4}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 985------------------------\n")

	for _, q := range qs {
		_, p := q.ans985, q.para985
		fmt.Printf("【input】:%v       【output】:%v\n\n\n", p, sumEvenAfterQueries(p.A, p.queries))
	}
	fmt.Printf("\n\n\n")
}
