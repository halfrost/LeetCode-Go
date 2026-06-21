package leetcode

import (
	"fmt"
	"testing"
)

type question18 struct {
	para18
	ans18
}

// para 是参数
// one 代表第一个参数
type para18 struct {
	a []int
	t int
}

// ans 是答案
// one 代表第一个答案
type ans18 struct {
	one [][]int
}

func Test_Problem18(t *testing.T) {

	qs := []question18{

		{
			para18{[]int{1, 1, 1, 1}, 4},
			ans18{[][]int{{1, 1, 1, 1}}},
		},

		{
			para18{[]int{0, 1, 5, 0, 1, 5, 5, -4}, 11},
			ans18{[][]int{{-4, 5, 5, 5}, {0, 1, 5, 5}}},
		},

		{
			para18{[]int{1, 0, -1, 0, -2, 2}, 0},
			ans18{[][]int{{-1, 0, 0, 1}, {-2, -1, 1, 2}, {-2, 0, 0, 2}}},
		},

		{
			para18{[]int{1, 0, -1, 0, -2, 2, 0, 0, 0, 0}, 0},
			ans18{[][]int{{-1, 0, 0, 1}, {-2, -1, 1, 2}, {-2, 0, 0, 2}, {0, 0, 0, 0}}},
		},

		{
			para18{[]int{1, 0, -1, 0, -2, 2, 0, 0, 0, 0}, 1},
			ans18{[][]int{{-1, 0, 0, 2}, {-2, 0, 1, 2}, {0, 0, 0, 1}}},
		},

		{
			para18{[]int{1, 1, 3, 3, 2}, 8},
			ans18{[][]int{{1, 1, 3, 3}}},
		},

		{
			para18{[]int{1, 1, 2, 3, 3, 2}, 7},
			ans18{[][]int{{1, 1, 2, 3}}},
		},

		{
			para18{[]int{2, 2, 2, 2, 1}, 8},
			ans18{[][]int{{2, 2, 2, 2}}},
		},

		// 右端有重复值（5,5）且左指针不会吃掉它，覆盖右指针去重分支
		{
			para18{[]int{1, 1, 3, 4, 5, 5}, 10},
			ans18{[][]int{{1, 1, 3, 5}}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 18------------------------\n")

	for _, q := range qs {
		a, p := q.ans18, q.para18
		got := fourSum(append([]int{}, p.a...), p.t)
		fmt.Printf("【input】:%v       【output】:%v\n", p, got)
		fourSum1(append([]int{}, p.a...), p.t)
		fourSum2(append([]int{}, p.a...), p.t)
		if !sameQuads(got, a.one) {
			t.Fatalf("fourSum(%v, %d) = %v, want %v", p.a, p.t, got, a.one)
		}
	}
	fmt.Printf("\n\n\n")
}

// sameQuads compares two sets of quadruplets ignoring order.
func sameQuads(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	key := func(q []int) string {
		return fmt.Sprintf("%v", q)
	}
	count := map[string]int{}
	for _, q := range a {
		count[key(q)]++
	}
	for _, q := range b {
		count[key(q)]--
	}
	for _, v := range count {
		if v != 0 {
			return false
		}
	}
	return true
}
