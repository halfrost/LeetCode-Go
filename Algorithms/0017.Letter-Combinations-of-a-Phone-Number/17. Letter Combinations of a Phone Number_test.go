package leetcode

import (
	"fmt"
	"testing"
)

type question17 struct {
	para17
	ans17
}

// para 是参数
// one 代表第一个参数
type para17 struct {
	s string
}

// ans 是答案
// one 代表第一个答案
type ans17 struct {
	one []string
}

func Test_Problem17(t *testing.T) {

	qs := []question17{

		question17{
			para17{"23"},
			ans17{[]string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 17------------------------\n")

	for _, q := range qs {
		_, p := q.ans17, q.para17
		fmt.Printf("【input】:%v       【output】:%v\n", p, letterCombinations(p.s))
	}
	fmt.Printf("\n\n\n")
}
