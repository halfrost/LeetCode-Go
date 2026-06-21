package leetcode

import (
	"fmt"
	"testing"
)

type question399 struct {
	para399
	ans399
}

// para 是参数
// one 代表第一个参数
type para399 struct {
	e [][]string
	v []float64
	q [][]string
}

// ans 是答案
// one 代表第一个答案
type ans399 struct {
	one []float64
}

func Test_Problem399(t *testing.T) {

	qs := []question399{

		{
			para399{[][]string{{"a", "b"}, {"b", "c"}}, []float64{2.0, 3.0}, [][]string{{"a", "c"}, {"b", "a"}, {"a", "e"}, {"a", "a"}, {"x", "x"}}},
			ans399{[]float64{6.0, 0.5, -1.0, 1.0, -1.0}},
		},

		{
			// two disconnected components: a/b and c/d, query across them hits the
			// different-set branch (res = -1).
			para399{[][]string{{"a", "b"}, {"c", "d"}}, []float64{2.0, 3.0}, [][]string{{"a", "d"}, {"a", "b"}}},
			ans399{[]float64{-1.0, 2.0}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 399------------------------\n")

	for _, q := range qs {
		a, p := q.ans399, q.para399
		got := calcEquation(p.e, p.v, p.q)
		fmt.Printf("【input】:%v       【output】:%v\n", p, got)
		if len(got) != len(a.one) {
			t.Fatalf("length mismatch: got %v, want %v", got, a.one)
		}
		for i := range got {
			if got[i] != a.one[i] {
				t.Fatalf("mismatch at %d: got %v, want %v", i, got, a.one)
			}
		}
	}
	fmt.Printf("\n\n\n")
}
