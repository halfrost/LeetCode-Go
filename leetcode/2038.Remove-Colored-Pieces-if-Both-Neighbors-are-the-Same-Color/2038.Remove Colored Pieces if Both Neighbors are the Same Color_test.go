package leetcode

import (
	"fmt"
	"testing"
)

type question2038 struct {
	para2038
	ans2038
}

// para 是参数
type para2038 struct {
	colors string
}

// ans 是答案
type ans2038 struct {
	ans bool
}

func Test_Problem2038(t *testing.T) {

	qs := []question2038{

		{
			para2038{"AAABABB"},
			ans2038{true},
		},

		{
			para2038{"AA"},
			ans2038{false},
		},

		{
			para2038{"ABBBBBBBAAA"},
			ans2038{false},
		},

		{
			// 回归用例：两段长度 3 的 B + 一段长度 4 的 A，旧实现会误判为 true
			para2038{"AAAABBBABBB"},
			ans2038{false},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 2038------------------------\n")

	for _, q := range qs {
		a, p := q.ans2038, q.para2038
		fmt.Printf("【input】:%v      ", p.colors)
		got := winnerOfGame(p.colors)
		fmt.Printf("【output】:%v      \n", got)
		if got != a.ans {
			t.Fatalf("winnerOfGame(%q) = %v, want %v", p.colors, got, a.ans)
		}
	}
	fmt.Printf("\n\n\n")
}
