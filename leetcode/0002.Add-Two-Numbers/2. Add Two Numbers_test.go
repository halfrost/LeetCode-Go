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
	one     []int
	another []int
}

// ans 是答案
// one 代表第一个答案
type ans2 struct {
	one []int
}

func Test_Problem2(t *testing.T) {

	qs := []question2{

		{
			para2{[]int{}, []int{}},
			ans2{[]int{}},
		},

		{
			para2{[]int{1}, []int{1}},
			ans2{[]int{2}},
		},

		{
			para2{[]int{1, 2, 3, 4}, []int{1, 2, 3, 4}},
			ans2{[]int{2, 4, 6, 8}},
		},

		{
			para2{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
			ans2{[]int{2, 4, 6, 8, 0, 1}},
		},

		{
			para2{[]int{1}, []int{9, 9, 9, 9, 9}},
			ans2{[]int{0, 0, 0, 0, 0, 1}},
		},

		{
			para2{[]int{9, 9, 9, 9, 9}, []int{1}},
			ans2{[]int{0, 0, 0, 0, 0, 1}},
		},

		{
			para2{[]int{2, 4, 3}, []int{5, 6, 4}},
			ans2{[]int{7, 0, 8}},
		},

		{
			para2{[]int{1, 8, 3}, []int{7, 1}},
			ans2{[]int{8, 9, 3}},
		},
		// 如需多个测试，可以复制上方元素。
	}

	fmt.Printf("------------------------Leetcode Problem 2------------------------\n")

	for _, q := range qs {
		_, p := q.ans2, q.para2
		fmt.Printf("【input】:%v       【output】:%v\n", p, structures.List2Ints(addTwoNumbers(structures.Ints2List(p.one), structures.Ints2List(p.another))))
	}
	fmt.Printf("\n\n\n")
}
