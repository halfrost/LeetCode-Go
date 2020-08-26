package leetcode

import (
	"fmt"
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

type question206 struct {
	para206
	ans206
}

// para 是参数
// one 代表第一个参数
type para206 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans206 struct {
	one []int
}

func Test_Problem206(t *testing.T) {

	qs := []question206{

		{
			para206{[]int{1, 2, 3, 4, 5}},
			ans206{[]int{5, 4, 3, 2, 1}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 206------------------------\n")

	for _, q := range qs {
		_, p := q.ans206, q.para206
		fmt.Printf("【input】:%v       【output】:%v\n", p, structures.List2Ints(reverseList(structures.Ints2List(p.one))))
	}
}
