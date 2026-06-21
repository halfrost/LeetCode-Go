package leetcode

import (
	"fmt"
	"testing"
)

type question287 struct {
	para287
	ans287
}

// para 是参数
// one 代表第一个参数
type para287 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans287 struct {
	one int
}

func Test_Problem287(t *testing.T) {

	qs := []question287{

		{
			para287{[]int{1, 3, 4, 2, 2}},
			ans287{2},
		},

		{
			para287{[]int{3, 1, 3, 4, 2}},
			ans287{3},
		},

		{
			para287{[]int{2, 2, 2, 2, 2}},
			ans287{2},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 287------------------------\n")

	for _, q := range qs {
		a, p := q.ans287, q.para287
		fmt.Printf("【input】:%v       【output】:%v\n", p, findDuplicate(p.one))
		if got := findDuplicate1(p.one); got != a.one {
			t.Fatalf("findDuplicate1(%v) = %v, want %v", p.one, got, a.one)
		}
		if got := findDuplicate2(p.one); got != a.one {
			t.Fatalf("findDuplicate2(%v) = %v, want %v", p.one, got, a.one)
		}
	}

	// 覆盖 findDuplicate2 的边界分支
	if got := findDuplicate2([]int{}); got != 0 {
		t.Fatalf("findDuplicate2(empty) = %v, want 0", got)
	}
	if got := findDuplicate2([]int{0}); got != 0 {
		t.Fatalf("findDuplicate2([0]) = %v, want 0", got)
	}
	fmt.Printf("\n\n\n")
}
