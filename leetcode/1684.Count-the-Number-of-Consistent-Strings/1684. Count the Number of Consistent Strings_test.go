package leetcode

import (
	"fmt"
	"testing"
)

type question1684 struct {
	para1684
	ans1684
}

// para 是参数
// one 代表第一个参数
type para1684 struct {
	allowed string
	words   []string
}

// ans 是答案
// one 代表第一个答案
type ans1684 struct {
	one int
}

func Test_Problem1684(t *testing.T) {

	qs := []question1684{

		{
			para1684{"ab", []string{"ad", "bd", "aaab", "baa", "badab"}},
			ans1684{2},
		},

		{
			para1684{"abc", []string{"a", "b", "c", "ab", "ac", "bc", "abc"}},
			ans1684{7},
		},

		{
			para1684{"cad", []string{"cc", "acd", "b", "ba", "bac", "bad", "ac", "d"}},
			ans1684{4},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1684------------------------\n")

	for _, q := range qs {
		_, p := q.ans1684, q.para1684
		fmt.Printf("【input】:%v       【output】:%v\n", p, countConsistentStrings(p.allowed, p.words))
	}
	fmt.Printf("\n\n\n")
}
