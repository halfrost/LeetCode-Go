package leetcode

import (
	"fmt"
	"testing"
)

type question841 struct {
	para841
	ans841
}

// para 是参数
// one 代表第一个参数
type para841 struct {
	rooms [][]int
}

// ans 是答案
// one 代表第一个答案
type ans841 struct {
	one bool
}

func Test_Problem841(t *testing.T) {

	qs := []question841{

		{
			para841{[][]int{{1}, {2}, {3}, {}}},
			ans841{true},
		},

		{
			para841{[][]int{{1, 3}, {3, 0, 1}, {2}, {0}}},
			ans841{false},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 841------------------------\n")

	for _, q := range qs {
		_, p := q.ans841, q.para841
		fmt.Printf("【input】:%v       【output】:%v\n", p, canVisitAllRooms(p.rooms))
	}
	fmt.Printf("\n\n\n")
}
