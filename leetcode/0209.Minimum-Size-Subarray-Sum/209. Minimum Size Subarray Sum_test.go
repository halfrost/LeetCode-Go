package leetcode

import (
	"fmt"
	"testing"
)

type question209 struct {
	para209
	ans209
}

// para 是参数
// one 代表第一个参数
type para209 struct {
	s   int
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans209 struct {
	one int
}

func Test_Problem209(t *testing.T) {

	qs := []question209{

		{
			para209{7, []int{2, 3, 1, 2, 4, 3}},
			ans209{2},
		},
		{
			para209{100, []int{1, 2, 3, 4}},
			ans209{0},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 209------------------------\n")

	for _, q := range qs {
		a, p := q.ans209, q.para209
		got := minSubArrayLen(p.s, p.one)
		if got != a.one {
			t.Fatalf("input: %v, expected: %v, got: %v", p, a.one, got)
		}
		fmt.Printf("【input】:%v       【output】:%v\n", p, got)
	}
	fmt.Printf("\n\n\n")
}
