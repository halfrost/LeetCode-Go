package leetcode

import (
	"fmt"
	"testing"
)

type question1657 struct {
	para1657
	ans1657
}

// para 是参数
// one 代表第一个参数
type para1657 struct {
	word1 string
	word2 string
}

// ans 是答案
// one 代表第一个答案
type ans1657 struct {
	one bool
}

func Test_Problem1657(t *testing.T) {

	qs := []question1657{

		{
			para1657{"abc", "bca"},
			ans1657{true},
		},

		{
			para1657{"a", "aa"},
			ans1657{false},
		},

		{
			para1657{"cabbba", "abbccc"},
			ans1657{true},
		},

		{
			para1657{"cabbba", "aabbss"},
			ans1657{false},
		},

		{
			para1657{"uau", "ssx"},
			ans1657{false},
		},

		{
			para1657{"uuukuuuukkuusuususuuuukuskuusuuusuusuuuuuuk", "kssskkskkskssskksskskksssssksskksskskksksuu"},
			ans1657{false},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1657------------------------\n")

	for _, q := range qs {
		_, p := q.ans1657, q.para1657
		fmt.Printf("【input】:%v      【output】:%v      \n", p, closeStrings(p.word1, p.word2))
	}
	fmt.Printf("\n\n\n")
}
