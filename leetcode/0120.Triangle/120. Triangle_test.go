package leetcode

import (
	"fmt"
	"testing"
)

type question120 struct {
	para120
	ans120
}

// para 是参数
// one 代表第一个参数
type para120 struct {
	one [][]int
}

// ans 是答案
// one 代表第一个答案
type ans120 struct {
	one int
}

func clone120(triangle [][]int) [][]int {
	if triangle == nil {
		return nil
	}
	res := make([][]int, len(triangle))
	for i, row := range triangle {
		res[i] = append([]int(nil), row...)
	}
	return res
}

func Test_Problem120(t *testing.T) {

	qs := []question120{
		{
			para120{[][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}},
			ans120{11},
		},
		{
			para120{nil},
			ans120{0},
		},
		{
			para120{[][]int{{-10}}},
			ans120{-10},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 120------------------------\n")

	for _, q := range qs {
		a, p := q.ans120, q.para120
		got := minimumTotal(clone120(p.one))
		fmt.Printf("【input】:%v       【output】:%v\n", p, got)
		if got != a.one {
			t.Fatalf("minimumTotal(%v) = %v, want %v", p.one, got, a.one)
		}
		got1 := minimumTotal1(clone120(p.one))
		if got1 != a.one {
			t.Fatalf("minimumTotal1(%v) = %v, want %v", p.one, got1, a.one)
		}
	}
	fmt.Printf("\n\n\n")
}
