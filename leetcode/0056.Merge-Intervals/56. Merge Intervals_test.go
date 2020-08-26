package leetcode

import (
	"fmt"
	"testing"
)

type question56 struct {
	para56
	ans56
}

// para 是参数
// one 代表第一个参数
type para56 struct {
	one []Interval
}

// ans 是答案
// one 代表第一个答案
type ans56 struct {
	one []Interval
}

func Test_Problem56(t *testing.T) {

	qs := []question56{

		{
			para56{[]Interval{}},
			ans56{[]Interval{}},
		},

		{
			para56{[]Interval{{Start: 1, End: 3}, {Start: 2, End: 6}, {Start: 8, End: 10}, {Start: 15, End: 18}}},
			ans56{[]Interval{{Start: 1, End: 6}, {Start: 8, End: 10}, {Start: 15, End: 18}}},
		},

		{
			para56{[]Interval{{Start: 1, End: 4}, {Start: 4, End: 5}}},
			ans56{[]Interval{{Start: 1, End: 5}}},
		},

		{
			para56{[]Interval{{Start: 1, End: 3}, {Start: 3, End: 6}, {Start: 5, End: 10}, {Start: 9, End: 18}}},
			ans56{[]Interval{{Start: 1, End: 18}}},
		},

		{
			para56{[]Interval{{Start: 15, End: 18}, {Start: 8, End: 10}, {Start: 2, End: 6}, {Start: 1, End: 3}}},
			ans56{[]Interval{{Start: 1, End: 6}, {Start: 8, End: 10}, {Start: 15, End: 18}}},
		},

		{
			para56{[]Interval{{Start: 1, End: 3}, {Start: 2, End: 6}, {Start: 8, End: 10}, {Start: 15, End: 18}}},
			ans56{[]Interval{{Start: 1, End: 6}, {Start: 8, End: 10}, {Start: 15, End: 18}}},
		},

		{
			para56{[]Interval{{Start: 1, End: 4}, {Start: 1, End: 5}}},
			ans56{[]Interval{{Start: 1, End: 5}}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 56------------------------\n")

	for _, q := range qs {
		_, p := q.ans56, q.para56
		fmt.Printf("【input】:%v       【output】:%v\n", p, merge56(p.one))
	}
	fmt.Printf("\n\n\n")
}
