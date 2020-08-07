package leetcode

import (
	"fmt"
	"testing"
)

type question212 struct {
	para212
	ans212
}

// para 是参数
// one 代表第一个参数
type para212 struct {
	b    [][]byte
	word []string
}

// ans 是答案
// one 代表第一个答案
type ans212 struct {
	one []string
}

func Test_Problem212(t *testing.T) {

	qs := []question212{

		question212{
			para212{[][]byte{
				[]byte{'A', 'B', 'C', 'E'},
				[]byte{'S', 'F', 'C', 'S'},
				[]byte{'A', 'D', 'E', 'E'},
			}, []string{"ABCCED", "SEE", "ABCB"}},
			ans212{[]string{"ABCCED", "SEE"}},
		},

		question212{
			para212{[][]byte{
				[]byte{'o', 'a', 'a', 'n'},
				[]byte{'e', 't', 'a', 'e'},
				[]byte{'i', 'h', 'k', 'r'},
				[]byte{'i', 'f', 'l', 'v'},
			}, []string{"oath", "pea", "eat", "rain"}},
			ans212{[]string{"oath", "eat"}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 212------------------------\n")

	for _, q := range qs {
		_, p := q.ans212, q.para212
		fmt.Printf("【input】:%v       【output】:%v\n\n\n", p, findWords(p.b, p.word))
	}
	fmt.Printf("\n\n\n")
}
