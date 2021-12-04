package leetcode

import (
	"fmt"
	"testing"
)

type question383 struct {
	para383
	ans383
}

// para 是参数
type para383 struct {
	ransomNote string
	magazine   string
}

// ans 是答案
type ans383 struct {
	ans bool
}

func Test_Problem383(t *testing.T) {

	qs := []question383{

		{
			para383{"a", "b"},
			ans383{false},
		},

		{
			para383{"aa", "ab"},
			ans383{false},
		},

		{
			para383{"aa", "aab"},
			ans383{true},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 383------------------------\n")

	for _, q := range qs {
		_, p := q.ans383, q.para383
		fmt.Printf("【input】:%v    【output】:%v\n", p, canConstruct(p.ransomNote, p.magazine))
	}
	fmt.Printf("\n\n\n")
}
