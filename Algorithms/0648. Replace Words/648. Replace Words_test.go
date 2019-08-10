package leetcode

import (
	"fmt"
	"testing"
)

type question648 struct {
	para648
	ans648
}

// para 是参数
// one 代表第一个参数
type para648 struct {
	one []string
	s   string
}

// ans 是答案
// one 代表第一个答案
type ans648 struct {
	one string
}

func Test_Problem648(t *testing.T) {

	qs := []question648{

		question648{
			para648{[]string{"cat", "bat", "rat"}, "the cattle was rattled by the battery"},
			ans648{"the cat was rat by the bat"},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 648------------------------\n")

	for _, q := range qs {
		_, p := q.ans648, q.para648
		fmt.Printf("【input】:%v       【output】:%v\n", p, replaceWords(p.one, p.s))
	}
	fmt.Printf("\n\n\n")
}
