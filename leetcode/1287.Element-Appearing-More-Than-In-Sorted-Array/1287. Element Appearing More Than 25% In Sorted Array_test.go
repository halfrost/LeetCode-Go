package leetcode

import (
	"fmt"
	"testing"
)

type question1287 struct {
	para1287
	ans1287
}

// para 是参数
// one 代表第一个参数
type para1287 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans1287 struct {
	one int
}

func Test_Problem1287(t *testing.T) {

	qs := []question1287{

		{
			para1287{[]int{1, 2, 2, 6, 6, 6, 6, 7, 10}},
			ans1287{6},
		},

		{
			para1287{[]int{1, 2, 3, 4, 5}},
			ans1287{-1},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1287------------------------\n")

	for _, q := range qs {
		a, p := q.ans1287, q.para1287
		out := findSpecialInteger(p.one)
		if out != a.one {
			t.Fatalf("input: %v, expected: %v, got: %v", p.one, a.one, out)
		}
		fmt.Printf("【input】:%v       【output】:%v\n", p, out)
	}
	fmt.Printf("\n\n\n")
}
