package leetcode

import (
	"fmt"
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type question1669 struct {
	para1669
	ans1669
}

// para 是参数
// one 代表第一个参数
type para1669 struct {
	one     []int
	a       int
	b       int
	another []int
}

// ans 是答案
// one 代表第一个答案
type ans1669 struct {
	one []int
}

func Test_Problem1669(t *testing.T) {

	qs := []question1669{

		{
			para1669{[]int{0, 1, 2, 3, 4, 5}, 3, 4, []int{1000000, 1000001, 1000002}},
			ans1669{[]int{0, 1, 2, 1000000, 1000001, 1000002, 5}},
		},

		{
			para1669{[]int{0, 1, 2, 3, 4, 5, 6}, 2, 5, []int{1000000, 1000001, 1000002, 1000003, 1000004}},
			ans1669{[]int{0, 1, 1000000, 1000001, 1000002, 1000003, 1000004, 6}},
		},

		{
			para1669{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 3, 5, []int{1000000, 1000001, 1000002, 1000003, 1000004, 1000005, 1000006}},
			ans1669{[]int{0, 1, 2, 1000000, 1000001, 1000002, 1000003, 1000004, 1000005, 1000006, 6, 7, 8, 9}},
		},

		{
			para1669{[]int{0, 1, 2}, 1, 1, []int{1000000, 1000001, 1000002, 1000003}},
			ans1669{[]int{0, 1000000, 1000001, 1000002, 1000003, 2}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1669------------------------\n")

	for _, q := range qs {
		_, p := q.ans1669, q.para1669
		fmt.Printf("【input】:%v       【output】:%v\n", p, structures.List2Ints(mergeInBetween(structures.Ints2List(p.one), p.a, p.b, structures.Ints2List(p.another))))
	}
	fmt.Printf("\n\n\n")
}
