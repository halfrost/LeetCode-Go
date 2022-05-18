package leetcode

import (
	"fmt"
	"testing"
)

type question728 struct {
	para728
	ans728
}

// para 是参数
type para728 struct {
	left  int
	right int
}

// ans 是答案
type ans728 struct {
	ans []int
}

func Test_Problem728(t *testing.T) {

	qs := []question728{

		{
			para728{1, 22},
			ans728{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 15, 22}},
		},

		{
			para728{47, 85},
			ans728{[]int{48, 55, 66, 77}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 728------------------------\n")

	for _, q := range qs {
		_, p := q.ans728, q.para728
		fmt.Printf("【input】:%v      ", p)
		fmt.Printf("【output】:%v      \n", selfDividingNumbers(p.left, p.right))
	}
	fmt.Printf("\n\n\n")
}
