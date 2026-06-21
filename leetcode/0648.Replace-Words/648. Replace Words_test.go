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

		{
			para648{[]string{"cat", "bat", "rat"}, "the cattle was rattled by the battery"},
			ans648{"the cat was rat by the bat"},
		},
		{
			para648{[]string{"cat", "cab", "ca"}, "category cable cattle"},
			ans648{"ca ca ca"},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 648------------------------\n")

	for _, q := range qs {
		a, p := q.ans648, q.para648
		got := replaceWords(p.one, p.s)
		fmt.Printf("【input】:%v       【output】:%v\n", p, got)
		if got != a.one {
			t.Fatalf("replaceWords(%v, %q) = %q, want %q", p.one, p.s, got, a.one)
		}
		got1 := replaceWords1(p.one, p.s)
		if got1 != a.one {
			t.Fatalf("replaceWords1(%v, %q) = %q, want %q", p.one, p.s, got1, a.one)
		}
	}
	fmt.Printf("\n\n\n")
}
