package leetcode

import (
	"fmt"
	"testing"
)

type question820 struct {
	para820
	ans820
}

// para 是参数
// one 代表第一个参数
type para820 struct {
	words []string
}

// ans 是答案
// one 代表第一个答案
type ans820 struct {
	one int
}

func Test_Problem820(t *testing.T) {

	qs := []question820{

		{
			para820{[]string{"time", "me", "bell"}},
			ans820{10},
		},

		{
			para820{[]string{"t"}},
			ans820{2},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 820------------------------\n")

	for _, q := range qs {
		_, p := q.ans820, q.para820
		fmt.Printf("【input】:%v       【output】:%v\n", p, minimumLengthEncoding(p.words))
	}
	fmt.Printf("\n\n\n")
}
