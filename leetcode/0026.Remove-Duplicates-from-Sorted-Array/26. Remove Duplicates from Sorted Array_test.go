package leetcode

import (
	"fmt"
	"testing"
)

type question26 struct {
	para26
	ans26
}

// para 是参数
// one 代表第一个参数
type para26 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans26 struct {
	one int
}

func clone26(nums []int) []int {
	c := make([]int, len(nums))
	copy(c, nums)
	return c
}

func Test_Problem26(t *testing.T) {

	qs := []question26{

		{
			para26{[]int{}},
			ans26{0},
		},

		{
			para26{[]int{1, 1, 2}},
			ans26{2},
		},

		{
			para26{[]int{0, 0, 1, 1, 1, 1, 2, 3, 4, 4}},
			ans26{5},
		},

		{
			para26{[]int{0, 0, 0, 0, 0}},
			ans26{1},
		},

		{
			para26{[]int{1}},
			ans26{1},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 26------------------------\n")

	for _, q := range qs {
		a, p := q.ans26, q.para26
		got := removeDuplicates(clone26(p.one))
		if got != a.one {
			t.Fatalf("removeDuplicates(%v) = %v, want %v", p.one, got, a.one)
		}
		got1 := removeDuplicates1(clone26(p.one))
		if got1 != a.one {
			t.Fatalf("removeDuplicates1(%v) = %v, want %v", p.one, got1, a.one)
		}
		fmt.Printf("【input】:%v    【output】:%v\n", p.one, got)
	}

	// directly exercise removeElement1 helper, covering empty input,
	// the equal-index branch and the swap branch.
	if removeElement1([]int{}, 0, 0) != 0 {
		t.Fatalf("removeElement1 empty should return 0")
	}
	if removeElement1([]int{2, 1, 3, 1, 4}, 0, 1) != 3 {
		t.Fatalf("removeElement1 unexpected result")
	}

	fmt.Printf("\n\n\n")
}
