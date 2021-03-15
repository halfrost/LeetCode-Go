package leetcode

import (
	"fmt"
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type question2 struct {
	para2
	ans2
}

// para 是参数
// one 代表第一个参数
type para2 struct {
	head []int
	k    int
}

// ans 是答案
// one 代表第一个答案
type ans2 struct {
	one []int
}

func Test_Problem2(t *testing.T) {

	qs := []question2{

		{
			para2{[]int{1, 2, 3, 4, 5}, 2},
			ans2{[]int{1, 4, 3, 2, 5}},
		},

		{
			para2{[]int{7, 9, 6, 6, 7, 8, 3, 0, 9, 5}, 5},
			ans2{[]int{7, 9, 6, 6, 8, 7, 3, 0, 9, 5}},
		},

		{
			para2{[]int{1}, 1},
			ans2{[]int{1}},
		},

		{
			para2{[]int{1, 2}, 1},
			ans2{[]int{2, 1}},
		},

		{
			para2{[]int{1, 2, 3}, 2},
			ans2{[]int{1, 2, 3}},
		},
		// 如需多个测试，可以复制上方元素。
	}

	fmt.Printf("------------------------Leetcode Problem 2------------------------\n")

	for _, q := range qs {
		_, p := q.ans2, q.para2
		fmt.Printf("【input】:%v       【output】:%v\n", p, structures.List2Ints(swapNodes(structures.Ints2List(p.head), p.k)))
	}
	fmt.Printf("\n\n\n")
}
