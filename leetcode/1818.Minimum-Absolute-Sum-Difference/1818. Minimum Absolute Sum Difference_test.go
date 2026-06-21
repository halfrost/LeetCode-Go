package leetcode

import (
	"fmt"
	"testing"
)

type question1818 struct {
	para1818
	ans1818
}

// para 是参数
// one 代表第一个参数
type para1818 struct {
	nums1 []int
	nums2 []int
}

// ans 是答案
// one 代表第一个答案
type ans1818 struct {
	one int
}

func Test_Problem1818(t *testing.T) {

	qs := []question1818{

		{
			para1818{[]int{1, 7, 5}, []int{2, 3, 5}},
			ans1818{3},
		},

		{
			para1818{[]int{2, 4, 6, 8, 10}, []int{2, 4, 6, 8, 10}},
			ans1818{0},
		},

		{
			para1818{[]int{1, 10, 4, 4, 2, 7}, []int{9, 3, 5, 1, 7, 4}},
			ans1818{20},
		},

		{
			para1818{[]int{1, 200000}, []int{200000, 1}},
			ans1818{199999},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1818------------------------\n")

	for _, q := range qs {
		a, p := q.ans1818, q.para1818
		got := minAbsoluteSumDiff(p.nums1, p.nums2)
		if got != a.one {
			t.Fatalf("input: %v, expected: %v, got: %v", p, a.one, got)
		}
		fmt.Printf("【input】:%v       【output】:%v\n", p, got)
	}
	fmt.Printf("\n\n\n")
}
