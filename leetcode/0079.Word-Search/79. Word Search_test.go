package leetcode

import (
	"fmt"
	"testing"
)

type question79 struct {
	para79
	ans79
}

// para 是参数
// one 代表第一个参数
type para79 struct {
	b    [][]byte
	word string
}

// ans 是答案
// one 代表第一个答案
type ans79 struct {
	one bool
}

func Test_Problem79(t *testing.T) {

	qs := []question79{

		{
			para79{[][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			}, "ABCCED"},
			ans79{true},
		},

		{
			para79{[][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			}, "SEE"},
			ans79{true},
		},

		{
			para79{[][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			}, "ABCB"},
			ans79{false},
		},

		{
			para79{[][]byte{
				{'o', 'a', 'a', 'n'},
				{'e', 't', 'a', 'e'},
				{'i', 'h', 'k', 'r'},
				{'i', 'f', 'l', 'v'},
			}, "oath"},
			ans79{true},
		},

		{
			para79{[][]byte{
				{'o', 'a', 'a', 'n'},
				{'e', 't', 'a', 'e'},
				{'i', 'h', 'k', 'r'},
				{'i', 'f', 'l', 'v'},
			}, "pea"},
			ans79{false},
		},

		{
			para79{[][]byte{
				{'o', 'a', 'a', 'n'},
				{'e', 't', 'a', 'e'},
				{'i', 'h', 'k', 'r'},
				{'i', 'f', 'l', 'v'},
			}, "eat"},
			ans79{true},
		},

		{
			para79{[][]byte{
				{'o', 'a', 'a', 'n'},
				{'e', 't', 'a', 'e'},
				{'i', 'h', 'k', 'r'},
				{'i', 'f', 'l', 'v'},
			}, "rain"},
			ans79{false},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 79------------------------\n")

	for _, q := range qs {
		_, p := q.ans79, q.para79
		fmt.Printf("【input】:%v       【output】:%v\n\n\n", p, exist(p.b, p.word))
	}
	fmt.Printf("\n\n\n")
}
