package leetcode

import (
	"fmt"
	"testing"
)

type question1239 struct {
	para1239
	ans1239
}

// para 是参数
// one 代表第一个参数
type para1239 struct {
	arr []string
}

// ans 是答案
// one 代表第一个答案
type ans1239 struct {
	one int
}

func Test_Problem1239(t *testing.T) {

	qs := []question1239{

		{
			para1239{[]string{"un", "iq", "ue"}},
			ans1239{4},
		},

		{
			para1239{[]string{"cha", "r", "act", "ers"}},
			ans1239{6},
		},

		{
			para1239{[]string{"abcdefghijklmnopqrstuvwxyz"}},
			ans1239{26},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1239------------------------\n")

	for _, q := range qs {
		_, p := q.ans1239, q.para1239
		fmt.Printf("【input】:%v       【output】:%v\n", p, maxLength(p.arr))
	}
	fmt.Printf("\n\n\n")
}
