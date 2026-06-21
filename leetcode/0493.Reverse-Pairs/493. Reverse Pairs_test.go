package leetcode

import (
	"fmt"
	"testing"
)

type question493 struct {
	para493
	ans493
}

// para 是参数
// one 代表第一个参数
type para493 struct {
	nums []int
}

// ans 是答案
// one 代表第一个答案
type ans493 struct {
	one int
}

func clone493(nums []int) []int {
	c := make([]int, len(nums))
	copy(c, nums)
	return c
}

func Test_Problem493(t *testing.T) {

	qs := []question493{

		{
			para493{[]int{1, 3, 2, 3, 1}},
			ans493{2},
		},

		{
			para493{[]int{9, 8, 7, 4, 7, 2, 3, 8, 7, 0}},
			ans493{18},
		},

		{
			para493{[]int{2, 4, 3, 5, 1}},
			ans493{3},
		},

		{
			para493{[]int{-5, -5}},
			ans493{1},
		},

		{
			para493{[]int{2147483647, 2147483647, -2147483647, -2147483647, -2147483647, 2147483647}},
			ans493{9},
		},

		{
			para493{[]int{1}},
			ans493{0},
		},

		{
			para493{[]int{}},
			ans493{0},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 493------------------------\n")

	for _, q := range qs {
		a, p := q.ans493, q.para493
		got := reversePairs(clone493(p.nums))
		fmt.Printf("【input】:%v       【output】:%v\n", p, got)
		if got != a.one {
			t.Fatalf("reversePairs(%v) = %v, want %v", p.nums, got, a.one)
		}
		if got1 := reversePairs1(clone493(p.nums)); got1 != a.one {
			t.Fatalf("reversePairs1(%v) = %v, want %v", p.nums, got1, a.one)
		}
		if got2 := reversePairs2(clone493(p.nums)); got2 != a.one {
			t.Fatalf("reversePairs2(%v) = %v, want %v", p.nums, got2, a.one)
		}
	}
	fmt.Printf("\n\n\n")
}
