package leetcode

import (
	"fmt"
	"testing"
)

type question1160 struct {
	para1160
	ans1160
}

// para 是参数
// one 代表第一个参数
type para1160 struct {
	words []string
	chars string
}

// ans 是答案
// one 代表第一个答案
type ans1160 struct {
	one int
}

func Test_Problem1160(t *testing.T) {

	qs := []question1160{

		question1160{
			para1160{[]string{"cat", "bt", "hat", "tree"}, "atach"},
			ans1160{6},
		},

		question1160{
			para1160{[]string{"hello", "world", "leetcode"}, "welldonehoneyr"},
			ans1160{10},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1160------------------------\n")

	for _, q := range qs {
		_, p := q.ans1160, q.para1160
		fmt.Printf("【input】:%v       【output】:%v\n", p, countCharacters(p.words, p.chars))
	}
	fmt.Printf("\n\n\n")
}
