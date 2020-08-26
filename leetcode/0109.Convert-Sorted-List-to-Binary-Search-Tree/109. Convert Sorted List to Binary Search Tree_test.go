package leetcode

import (
	"fmt"
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type question109 struct {
	para109
	ans109
}

// para 是参数
// one 代表第一个参数
type para109 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans109 struct {
	one []int
}

func Test_Problem109(t *testing.T) {

	qs := []question109{

		{
			para109{[]int{-10, -3, 0, 5, 9}},
			ans109{[]int{0, -10, 5, structures.NULL, -3, structures.NULL, 9}},
		},

		{
			para109{[]int{-10}},
			ans109{[]int{-10}},
		},

		{
			para109{[]int{1, 2}},
			ans109{[]int{1, 2}},
		},

		{
			para109{[]int{1, 2, 3}},
			ans109{[]int{2, 1, 3}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 109------------------------\n")

	for _, q := range qs {
		_, p := q.ans109, q.para109
		arr := []int{}
		structures.T2s(sortedListToBST(structures.Ints2List(p.one)), &arr)
		fmt.Printf("【input】:%v       【output】:%v\n", p, arr)
	}
	fmt.Printf("\n\n\n")
}
