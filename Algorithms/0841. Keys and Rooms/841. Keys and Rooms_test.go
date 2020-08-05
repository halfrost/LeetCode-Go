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

		question841{
			para841{[][]int{[]int{1}, []int{2}, []int{3}, []int{}}},
			ans841{true},
		},

		question841{
			para841{[][]int{[]int{1, 3}, []int{3, 0, 1}, []int{2}, []int{0}}},
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
