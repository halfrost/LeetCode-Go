package leetcode

import (
	"fmt"
	"testing"
)

type question5618 struct {
	para5618
	ans5618
}

// para 是参数
// one 代表第一个参数
type para5618 struct {
	nums []int
	k    int
}

// ans 是答案
// one 代表第一个答案
type ans5618 struct {
	one int
}

func Test_Problem5618(t *testing.T) {

	qs := []question5618{

		{
			para5618{[]int{1, 2, 3, 4}, 5},
			ans5618{2},
		},

		{
			para5618{[]int{3, 1, 3, 4, 3}, 6},
			ans5618{1},
		},

		{
			para5618{[]int{2, 5, 4, 4, 1, 3, 4, 4, 1, 4, 4, 1, 2, 1, 2, 2, 3, 2, 4, 2}, 3},
			ans5618{4},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 5618------------------------\n")

	for _, q := range qs {
		_, p := q.ans5618, q.para5618
		fmt.Printf("【input】:%v       【output】:%v\n", p, maxOperations(p.nums, p.k))
	}
	fmt.Printf("\n\n\n")
}
