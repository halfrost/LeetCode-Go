package leetcode

import (
	"fmt"
	"testing"
)

type question763 struct {
	para763
	ans763
}

// para 是参数
// one 代表第一个参数
type para763 struct {
	one string
}

// ans 是答案
// one 代表第一个答案
type ans763 struct {
	one []int
}

func Test_Problem763(t *testing.T) {

	qs := []question763{

		question763{
			para763{"ababcbacadefegdehijhklij"},
			ans763{[]int{9, 7, 8}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 763------------------------\n")

	for _, q := range qs {
		_, p := q.ans763, q.para763
		fmt.Printf("【input】:%v       【output】:%v\n", p, partitionLabels(p.one))
	}
	fmt.Printf("\n\n\n")
}
