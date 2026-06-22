package leetcode

import (
	"fmt"
	"testing"
)

type question238 struct {
	para238
	ans238
}

// para is the parameter
// one represents the first parameter
type para238 struct {
	one []int
}

// ans is the answer
// one represents the first answer
type ans238 struct {
	one []int
}

func Test_Problem238(t *testing.T) {

	qs := []question238{

		{
			para238{[]int{1, 2, 3, 4}},
			ans238{[]int{24, 12, 8, 6}},
		},

		{
			para238{[]int{-1, 1, 0, -3, 3}},
			ans238{[]int{0, 0, 9, 0, 0}},
		},

		{
			para238{[]int{2, 3}},
			ans238{[]int{3, 2}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 238------------------------\n")

	for _, q := range qs {
		a, p := q.ans238, q.para238
		got := productExceptSelf(p.one)
		fmt.Printf("【input】:%v       【output】:%v\n", p, got)
		if !equal238(got, a.one) {
			t.Fatalf("input %v: expected %v, got %v", p.one, a.one, got)
		}
	}
	fmt.Printf("\n\n\n")
}

func equal238(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
