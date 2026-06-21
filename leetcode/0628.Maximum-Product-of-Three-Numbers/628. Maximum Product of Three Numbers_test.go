package leetcode

import (
	"fmt"
	"testing"
)

type question628 struct {
	para628
	ans628
}

// para 是参数
// one 代表第一个参数
type para628 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans628 struct {
	one int
}

func Test_Problem628(t *testing.T) {

	qs := []question628{

		{
			para628{[]int{3, -1, 4}},
			ans628{-12},
		},

		{
			para628{[]int{1, 2}},
			ans628{2},
		},

		{
			para628{[]int{1, 2, 3}},
			ans628{6},
		},

		{
			para628{[]int{1, 2, 3, 4}},
			ans628{24},
		},

		{
			para628{[]int{-2}},
			ans628{-2},
		},

		{
			para628{[]int{0}},
			ans628{0},
		},

		{
			para628{[]int{2, 3, -2, 4}},
			ans628{24},
		},

		{
			para628{[]int{-2, 0, -1}},
			ans628{0},
		},

		{
			para628{[]int{-2, 0, -1, 2, 3, 1, 10}},
			ans628{60},
		},

		{
			para628{[]int{}},
			ans628{0},
		},

		{
			para628{[]int{-4, -3, -2, -1}},
			ans628{0},
		},

		{
			para628{[]int{-10, -10, 1, 2, 3}},
			ans628{300},
		},

		{
			para628{[]int{5, 4, 4, 3}},
			ans628{80},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 628------------------------\n")

	for _, q := range qs {
		a, p := q.ans628, q.para628
		clone := make([]int, len(p.one))
		copy(clone, p.one)
		got := maximumProduct(p.one)
		fmt.Printf("【input】:%v       【output】:%v\n", p, got)
		if got != a.one {
			t.Fatalf("maximumProduct(%v) = %v, want %v", clone, got, a.one)
		}
		// maximumProduct1 is an O(n) alternative. It always returns the
		// true maximum product of three numbers (it has no "all values
		// <= 0 returns 0" short-circuit that maximumProduct uses), so it
		// only needs to agree with maximumProduct when the array has at
		// least three elements and contains a positive value.
		got1 := maximumProduct1(clone)
		if len(clone) >= 3 && got != 0 && got1 != a.one {
			t.Fatalf("maximumProduct1(%v) = %v, want %v", clone, got1, a.one)
		}
	}
	fmt.Printf("\n\n\n")
}
