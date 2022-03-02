package leetcode

import (
	"fmt"
	"testing"
)

type question540 struct {
	para540
	ans540
}

// para 是参数
type para540 struct {
	nums []int
}

// ans 是答案
type ans540 struct {
	ans int
}

func Test_Problem540(t *testing.T) {

	qs := []question540{

		{
			para540{[]int{1, 1, 2, 3, 3, 4, 4, 8, 8}},
			ans540{2},
		},

		{
			para540{[]int{3, 3, 7, 7, 10, 11, 11}},
			ans540{10},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 540------------------------\n")

	for _, q := range qs {
		_, p := q.ans540, q.para540
		fmt.Printf("【input】:%v      ", p.nums)
		fmt.Printf("【output】:%v      \n", singleNonDuplicate(p.nums))
	}
	fmt.Printf("\n\n\n")
}
