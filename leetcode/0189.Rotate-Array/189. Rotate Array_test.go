package leetcode

import (
	"fmt"
	"testing"
)

type question189 struct {
	para189
	ans189
}

// para 是参数
// one 代表第一个参数
type para189 struct {
	nums []int
	k    int
}

// ans 是答案
// one 代表第一个答案
type ans189 struct {
	one []int
}

func Test_Problem189(t *testing.T) {

	qs := []question189{

		{
			para189{[]int{1, 2, 3, 4, 5, 6, 7}, 3},
			ans189{[]int{5, 6, 7, 1, 2, 3, 4}},
		},

		{
			para189{[]int{-1, -100, 3, 99}, 2},
			ans189{[]int{3, 99, -1, -100}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 189------------------------\n")

	for _, q := range qs {
		_, p := q.ans189, q.para189
		fmt.Printf("【input】:%v       ", p)
		rotate(p.nums, p.k)
		fmt.Printf("【output】:%v\n", p.nums)
	}
	fmt.Printf("\n\n\n")
}
