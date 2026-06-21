package leetcode

import (
	"fmt"
	"testing"
)

type question852 struct {
	para852
	ans852
}

// para 是参数
// one 代表第一个参数
type para852 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans852 struct {
	one int
}

func Test_Problem852(t *testing.T) {

	qs := []question852{

		{
			para852{[]int{0, 1, 0}},
			ans852{1},
		},

		{
			para852{[]int{0, 2, 1, 0}},
			ans852{1},
		},

		{
			para852{[]int{0, 5, 10, 9, 8, 7, 6, 5, 4}},
			ans852{2},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 852------------------------\n")

	for _, q := range qs {
		a, p := q.ans852, q.para852
		got := peakIndexInMountainArray(p.one)
		fmt.Printf("【input】:%v       【output】:%v\n", p, got)
		if got != a.one {
			t.Fatalf("peakIndexInMountainArray(%v) = %d, want %d", p.one, got, a.one)
		}
		if got1 := peakIndexInMountainArray1(p.one); got1 != a.one {
			t.Fatalf("peakIndexInMountainArray1(%v) = %d, want %d", p.one, got1, a.one)
		}
	}
	fmt.Printf("\n\n\n")
}
