package leetcode

import (
	"fmt"
	"testing"
)

type question720 struct {
	para720
	ans720
}

// para 是参数
// one 代表第一个参数
type para720 struct {
	w []string
}

// ans 是答案
// one 代表第一个答案
type ans720 struct {
	one string
}

func Test_Problem720(t *testing.T) {

	qs := []question720{

		question720{
			para720{[]string{"w", "wo", "wor", "worl", "world"}},
			ans720{"world"},
		},

		question720{
			para720{[]string{"a", "banana", "app", "appl", "ap", "apply", "apple"}},
			ans720{"apple"},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 720------------------------\n")

	for _, q := range qs {
		_, p := q.ans720, q.para720
		fmt.Printf("【input】:%v       【output】:%v\n", p, longestWord(p.w))
	}
	fmt.Printf("\n\n\n")
}
