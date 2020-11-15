package leetcode

import (
	"fmt"
	"testing"
)

type question1649 struct {
	para1649
	ans1649
}

// para 是参数
// one 代表第一个参数
type para1649 struct {
	word1 string
	word2 string
}

// ans 是答案
// one 代表第一个答案
type ans1649 struct {
	one bool
}

func Test_Problem1649(t *testing.T) {

	qs := []question1649{

		{
			para1649{"abc", "bca"},
			ans1649{true},
		},

		{
			para1649{"a", "aa"},
			ans1649{false},
		},

		{
			para1649{"cabbba", "abbccc"},
			ans1649{true},
		},

		{
			para1649{"cabbba", "aabbss"},
			ans1649{false},
		},

		{
			para1649{"uau", "ssx"},
			ans1649{false},
		},

		{
			para1649{"uuukuuuukkuusuususuuuukuskuusuuusuusuuuuuuk", "kssskkskkskssskksskskksssssksskksskskksksuu"},
			ans1649{false},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1649------------------------\n")

	for _, q := range qs {
		_, p := q.ans1649, q.para1649
		fmt.Printf("【input】:%v      【output】:%v      \n", p, closeStrings(p.word1, p.word2))
	}
	fmt.Printf("\n\n\n")
}
