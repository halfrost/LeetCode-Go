package leetcode

import (
	"fmt"
	"testing"
)

type question966 struct {
	para966
	ans966
}

// para 是参数
// one 代表第一个参数
type para966 struct {
	wordlist []string
	queries  []string
}

// ans 是答案
// one 代表第一个答案
type ans966 struct {
	one []string
}

func Test_Problem966(t *testing.T) {

	qs := []question966{
		{
			para966{[]string{"KiTe", "kite", "hare", "Hare"}, []string{"kite", "Kite", "KiTe", "Hare", "HARE", "Hear", "hear", "keti", "keet", "keto"}},
			ans966{[]string{"kite", "KiTe", "KiTe", "Hare", "hare", "", "", "KiTe", "", "KiTe"}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 966------------------------\n")

	for _, q := range qs {
		_, p := q.ans966, q.para966
		fmt.Printf("【input】:%v       【output】:%#v\n", p, spellchecker(p.wordlist, p.queries))
	}
	fmt.Printf("\n\n\n")
}
