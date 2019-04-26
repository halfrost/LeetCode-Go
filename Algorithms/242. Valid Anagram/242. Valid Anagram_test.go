package leetcode

import (
	"fmt"
	"testing"
)

type question242 struct {
	para242
	ans242
}

// para 是参数
// one 代表第一个参数
type para242 struct {
	one string
	two string
}

// ans 是答案
// one 代表第一个答案
type ans242 struct {
	one bool
}

func Test_Problem242(t *testing.T) {

	qs := []question242{

		question242{
			para242{"", ""},
			ans242{true},
		},
		question242{
			para242{"", "1"},
			ans242{false},
		},

		question242{
			para242{"anagram", "nagaram"},
			ans242{true},
		},

		question242{
			para242{"rat", "car"},
			ans242{false},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 242------------------------\n")

	for _, q := range qs {
		_, p := q.ans242, q.para242
		fmt.Printf("【input】:%v       【output】:%v\n", p, isAnagram(p.one, p.two))
	}
	fmt.Printf("\n\n\n")
}
