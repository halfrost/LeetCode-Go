package leetcode

import (
	"fmt"
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type question1290 struct {
	para1290
	ans1290
}

// para 是参数
// one 代表第一个参数
type para1290 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans1290 struct {
	one int
}

func Test_Problem1290(t *testing.T) {

	qs := []question1290{

		{
			para1290{[]int{1, 0, 1}},
			ans1290{5},
		},

		{
			para1290{[]int{0}},
			ans1290{0},
		},

		{
			para1290{[]int{1}},
			ans1290{1},
		},

		{
			para1290{[]int{0, 0}},
			ans1290{0},
		},

		{
			para1290{[]int{1, 0, 0, 1, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0}},
			ans1290{18880},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1290------------------------\n")

	for _, q := range qs {
		_, p := q.ans1290, q.para1290
		fmt.Printf("【input】:%v       【output】:%v\n", p, getDecimalValue(structures.Ints2List(p.one)))
	}
	fmt.Printf("\n\n\n")
}
