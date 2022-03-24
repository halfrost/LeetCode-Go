package leetcode

import (
	"fmt"
	"testing"
)

type question2037 struct {
	para2037
	ans2037
}

// para 是参数
type para2037 struct {
	seats    []int
	students []int
}

// ans 是答案
type ans2037 struct {
	ans int
}

func Test_Problem2037(t *testing.T) {

	qs := []question2037{

		{
			para2037{[]int{3, 1, 5}, []int{2, 7, 4}},
			ans2037{4},
		},

		{
			para2037{[]int{4, 1, 5, 9}, []int{1, 3, 2, 6}},
			ans2037{7},
		},

		{
			para2037{[]int{2, 2, 6, 6}, []int{1, 3, 2, 6}},
			ans2037{4},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 2037------------------------\n")

	for _, q := range qs {
		_, p := q.ans2037, q.para2037
		fmt.Printf("【input】:%v      ", p)
		fmt.Printf("【output】:%v      \n", minMovesToSeat(p.seats, p.students))
	}
	fmt.Printf("\n\n\n")
}
