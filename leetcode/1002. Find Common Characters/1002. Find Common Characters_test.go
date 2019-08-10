package leetcode

import (
	"fmt"
	"testing"
)

type question1002 struct {
	para1002
	ans1002
}

// para 是参数
// one 代表第一个参数
type para1002 struct {
	one []string
}

// ans 是答案
// one 代表第一个答案
type ans1002 struct {
	one []string
}

func Test_Problem1002(t *testing.T) {

	qs := []question1002{

		question1002{
			para1002{[]string{"bella", "label", "roller"}},
			ans1002{[]string{"e", "l", "l"}},
		},

		question1002{
			para1002{[]string{"cool", "lock", "cook"}},
			ans1002{[]string{"c", "o"}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1002------------------------\n")

	for _, q := range qs {
		_, p := q.ans1002, q.para1002
		fmt.Printf("【input】:%v       【output】:%v\n", p, commonChars(p.one))
	}
	fmt.Printf("\n\n\n")
}
