package leetcode

import (
	"fmt"
	"testing"
)

type question219 struct {
	para219
	ans219
}

// para 是参数
// one 代表第一个参数
type para219 struct {
	one []int
	k   int
}

// ans 是答案
// one 代表第一个答案
type ans219 struct {
	one bool
}

func Test_Problem219(t *testing.T) {

	qs := []question219{

		{
			para219{[]int{1, 2, 3, 1}, 3},
			ans219{true},
		},

		{
			para219{[]int{1, 0, 0, 1}, 1},
			ans219{true},
		},

		{
			para219{[]int{1, 2, 3, 1, 2, 3}, 2},
			ans219{false},
		},

		{
			para219{[]int{1}, 1},
			ans219{false},
		},

		{
			para219{[]int{1, 1}, 0},
			ans219{false},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 219------------------------\n")

	for _, q := range qs {
		a, p := q.ans219, q.para219
		got := containsNearbyDuplicate(p.one, p.k)
		if got != a.one {
			t.Fatalf("input: %v, expected: %v, got: %v", p, a.one, got)
		}
		fmt.Printf("【input】:%v       【output】:%v\n", p, got)
	}
	fmt.Printf("\n\n\n")
}
