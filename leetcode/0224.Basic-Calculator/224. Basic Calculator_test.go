package leetcode

import (
	"fmt"
	"testing"
)

type question224 struct {
	para224
	ans224
}

// para 是参数
// one 代表第一个参数
type para224 struct {
	one string
}

// ans 是答案
// one 代表第一个答案
type ans224 struct {
	one int
}

func Test_Problem224(t *testing.T) {

	qs := []question224{

		{
			para224{"1 + 1"},
			ans224{2},
		},
		{
			para224{" 2-1 + 2 "},
			ans224{3},
		},

		{
			para224{"(1+(4+5+2)-3)+(6+8)"},
			ans224{23},
		},

		{
			para224{"2-(5-6)"},
			ans224{3},
		},

		{
			para224{"100 + 23 - 12"},
			ans224{111},
		},

		{
			para224{"2-(3-(4-1))"},
			ans224{3},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 224------------------------\n")

	for _, q := range qs {
		_, p := q.ans224, q.para224
		fmt.Printf("【input】:%v       【output】:%v\n", p, calculate(p.one))
		calculate1(p.one)
	}

	// 覆盖 calculateStr 中连续符号合并的分支：++、+-、-+、--
	signCases := []question224{
		{para224{"2++3"}, ans224{5}},
		{para224{"2+(0-3)"}, ans224{-1}},
		{para224{"2-+3"}, ans224{-1}},
		{para224{"2-(0-3)"}, ans224{5}},
	}
	for _, q := range signCases {
		if got := calculate1(q.para224.one); got != q.ans224.one {
			t.Fatalf("calculate1(%q) = %v, want %v", q.para224.one, got, q.ans224.one)
		}
	}

	fmt.Printf("\n\n\n")
}
