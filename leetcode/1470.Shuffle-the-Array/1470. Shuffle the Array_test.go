package leetcode

import (
	"fmt"
	"testing"
)

type question1470 struct {
	para1470
	ans1470
}

// para 是参数
// one 代表第一个参数
type para1470 struct {
	nums []int
	n    int
}

// ans 是答案
// one 代表第一个答案
type ans1470 struct {
	one []int
}

func Test_Problem1470(t *testing.T) {

	qs := []question1470{

		{
			para1470{[]int{2, 5, 1, 3, 4, 7}, 3},
			ans1470{[]int{2, 3, 5, 4, 1, 7}},
		},

		{
			para1470{[]int{1, 2, 3, 4, 4, 3, 2, 1}, 4},
			ans1470{[]int{1, 4, 2, 3, 3, 2, 4, 1}},
		},

		{
			para1470{[]int{1, 1, 2, 2}, 2},
			ans1470{[]int{1, 2, 1, 2}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1470------------------------\n")

	for _, q := range qs {
		_, p := q.ans1470, q.para1470
		fmt.Printf("【input】:%v      【output】:%v      \n", p, shuffle(p.nums, p.n))
	}
	fmt.Printf("\n\n\n")
}
