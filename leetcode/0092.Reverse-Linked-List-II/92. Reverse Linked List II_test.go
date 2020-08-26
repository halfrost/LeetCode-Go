package leetcode

import (
	"fmt"
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type question92 struct {
	para92
	ans92
}

// para 是参数
// one 代表第一个参数
type para92 struct {
	one  []int
	m, n int
}

// ans 是答案
// one 代表第一个答案
type ans92 struct {
	one []int
}

func Test_Problem92(t *testing.T) {

	qs := []question92{

		{
			para92{[]int{1, 2, 3, 4, 5}, 2, 4},
			ans92{[]int{1, 4, 3, 2, 5}},
		},

		{
			para92{[]int{1, 2, 3, 4, 5}, 2, 2},
			ans92{[]int{1, 2, 3, 4, 5}},
		},

		{
			para92{[]int{1, 2, 3, 4, 5}, 1, 5},
			ans92{[]int{5, 4, 3, 2, 1}},
		},

		{
			para92{[]int{1, 2, 3, 4, 5, 6}, 3, 4},
			ans92{[]int{1, 2, 4, 3, 5, 6}},
		},

		{
			para92{[]int{3, 5}, 1, 2},
			ans92{[]int{5, 3}},
		},

		{
			para92{[]int{3}, 3, 5},
			ans92{[]int{3}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 92------------------------\n")

	for _, q := range qs {
		_, p := q.ans92, q.para92
		fmt.Printf("【input】:%v       【output】:%v\n", p, structures.List2Ints(reverseBetween(structures.Ints2List(p.one), p.m, p.n)))
	}
	fmt.Printf("\n\n\n")
}
