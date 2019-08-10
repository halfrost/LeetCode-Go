package leetcode

import (
	"fmt"
	"testing"
)

type question389 struct {
	para389
	ans389
}

// para 是参数
// one 代表第一个参数
type para389 struct {
	s string
	t string
}

// ans 是答案
// one 代表第一个答案
type ans389 struct {
	one byte
}

func Test_Problem389(t *testing.T) {

	qs := []question389{

		question389{
			para389{"abcd", "abcde"},
			ans389{'e'},
		},

		// 如需多个测试，可以复制上方元素。
	}

	fmt.Printf("------------------------Leetcode Problem 389------------------------\n")

	for _, q := range qs {
		_, p := q.ans389, q.para389
		fmt.Printf("【input】:%v       【output】:%v\n", p, findTheDifference(p.s, p.t))
	}
	fmt.Printf("\n\n\n")
}
