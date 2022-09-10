package leetcode

import (
	"fmt"
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type question876 struct {
	para876
	ans876
}

// para 是参数
// one 代表第一个参数
type para876 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans876 struct {
	one int
}

func Test_Problem876(t *testing.T) {

	qs := []question876{

		{
			para876{[]int{1, 2, 3, 4, 5}},
			ans876{3},
		},
		{
			para876{[]int{1, 2, 3, 4}},
			ans876{3},
		},

		{
			para876{[]int{1}},
			ans876{1},
		},

		{
			para876{[]int{}},
			ans876{},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 876------------------------\n")

	for _, q := range qs {
		_, p := q.ans876, q.para876
		fmt.Printf("【input】:%v       【output】:%v\n", p, structures.List2Ints(middleNode(structures.Ints2List(p.one))))
	}
	fmt.Printf("\n\n\n")
}
