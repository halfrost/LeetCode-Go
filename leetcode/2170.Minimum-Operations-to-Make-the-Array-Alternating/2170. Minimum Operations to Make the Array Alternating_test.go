package leetcode

import (
	"fmt"
	"testing"
)

type question2170 struct {
	para2170
	ans2170
}

// para 是参数
// one 代表第一个参数
type para2170 struct {
	nums []int
}

// ans 是答案
// one 代表第一个答案
type ans2170 struct {
	one int
}

func Test_Problem1(t *testing.T) {

	qs := []question2170{

		{
			para2170{[]int{1}},
			ans2170{0},
		},

		{
			para2170{[]int{3, 1, 3, 2, 4, 3}},
			ans2170{3},
		},

		{
			para2170{[]int{1, 2, 2, 2, 2}},
			ans2170{2},
		},

		{
			para2170{[]int{69, 91, 47, 74, 75, 94, 22, 100, 43, 50, 82, 47, 40, 51, 90, 27, 98, 85, 47, 14, 55, 82, 52, 9, 65, 90, 86, 45, 52, 52, 95, 40, 85, 3, 46, 77, 16, 59, 32, 22, 41, 87, 89, 78, 59, 78, 34, 26, 71, 9, 82, 68, 80, 74, 100, 6, 10, 53, 84, 80, 7, 87, 3, 82, 26, 26, 14, 37, 26, 58, 96, 73, 41, 2, 79, 43, 56, 74, 30, 71, 6, 100, 72, 93, 83, 40, 28, 79, 24}},
			ans2170{84},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 2170------------------------\n")

	for _, q := range qs {
		_, p := q.ans2170, q.para2170
		fmt.Printf("【input】:%v       【output】:%v\n", p, minimumOperations(p.nums))
	}
	fmt.Printf("\n\n\n")
}
