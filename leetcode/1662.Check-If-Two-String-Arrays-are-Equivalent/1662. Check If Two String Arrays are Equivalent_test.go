package leetcode

import (
	"fmt"
	"testing"
)

type question1662 struct {
	para1662
	ans1662
}

// para 是参数
// one 代表第一个参数
type para1662 struct {
	word1 []string
	word2 []string
}

// ans 是答案
// one 代表第一个答案
type ans1662 struct {
	one bool
}

func Test_Problem1662(t *testing.T) {

	qs := []question1662{

		{
			para1662{[]string{"ab", "c"}, []string{"a", "bc"}},
			ans1662{true},
		},

		{
			para1662{[]string{"a", "cb"}, []string{"ab", "c"}},
			ans1662{false},
		},

		{
			para1662{[]string{"abc", "d", "defg"}, []string{"abcddefg"}},
			ans1662{true},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1662------------------------\n")

	for _, q := range qs {
		_, p := q.ans1662, q.para1662
		fmt.Printf("【input】:%v      【output】:%v      \n", p, arrayStringsAreEqual(p.word1, p.word2))
	}
	fmt.Printf("\n\n\n")
}
