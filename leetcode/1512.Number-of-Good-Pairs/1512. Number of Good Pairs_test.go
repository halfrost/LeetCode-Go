package leetcode

import (
	"fmt"
	"testing"
)

type question1512 struct {
	para1512
	ans1512
}

// para 是参数
// one 代表第一个参数
type para1512 struct {
	nums []int
}

// ans 是答案
// one 代表第一个答案
type ans1512 struct {
	one int
}

func Test_Problem1512(t *testing.T) {

	qs := []question1512{

		{
			para1512{[]int{1, 2, 3, 1, 1, 3}},
			ans1512{4},
		},

		{
			para1512{[]int{1, 1, 1, 1}},
			ans1512{6},
		},

		{
			para1512{[]int{1, 2, 3}},
			ans1512{0},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1512------------------------\n")

	for _, q := range qs {
		_, p := q.ans1512, q.para1512
		fmt.Printf("【input】:%v      【output】:%v      \n", p, numIdenticalPairs(p.nums))
	}
	fmt.Printf("\n\n\n")
}
