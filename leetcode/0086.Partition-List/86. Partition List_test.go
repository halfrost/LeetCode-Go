package leetcode

import (
	"fmt"
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type question86 struct {
	para86
	ans86
}

// para 是参数
// one 代表第一个参数
type para86 struct {
	one []int
	x   int
}

// ans 是答案
// one 代表第一个答案
type ans86 struct {
	one []int
}

func Test_Problem86(t *testing.T) {

	qs := []question86{

		{
			para86{[]int{1, 4, 3, 2, 5, 2}, 3},
			ans86{[]int{1, 2, 2, 4, 3, 5}},
		},

		{
			para86{[]int{1, 1, 2, 2, 3, 3, 3}, 2},
			ans86{[]int{1, 1, 2, 2, 3, 3, 3}},
		},

		{
			para86{[]int{1, 4, 3, 2, 5, 2}, 0},
			ans86{[]int{1, 4, 3, 2, 5, 2}},
		},

		{
			para86{[]int{4, 3, 2, 5, 2}, 3},
			ans86{[]int{2, 2, 4, 3, 5}},
		},

		{
			para86{[]int{1, 1, 1, 1, 1, 1}, 1},
			ans86{[]int{1, 1, 1, 1, 1, 1}},
		},

		{
			para86{[]int{3, 1}, 2},
			ans86{[]int{1, 3}},
		},

		{
			para86{[]int{1, 2}, 3},
			ans86{[]int{1, 2}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 86------------------------\n")

	for _, q := range qs {
		_, p := q.ans86, q.para86
		fmt.Printf("【input】:%v       【output】:%v\n", p, structures.List2Ints(partition(structures.Ints2List(p.one), p.x)))
	}
	fmt.Printf("\n\n\n")
}
