package leetcode

import (
	"fmt"
	"testing"
)

type question851 struct {
	para851
	ans851
}

// para 是参数
// one 代表第一个参数
type para851 struct {
	richer [][]int
	quiet  []int
}

// ans 是答案
// one 代表第一个答案
type ans851 struct {
	one []int
}

func Test_Problem851(t *testing.T) {

	qs := []question851{

		{
			para851{[][]int{{1, 0}, {2, 1}, {3, 1}, {3, 7}, {4, 3}, {5, 3}}, []int{3, 2, 5, 4, 6, 1, 7, 0}},
			ans851{[]int{5, 5, 2, 5, 4, 5, 6, 7}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 851------------------------\n")

	for _, q := range qs {
		_, p := q.ans851, q.para851
		fmt.Printf("【input】:%v       【output】:%v\n", p, loudAndRich(p.richer, p.quiet))
	}
	fmt.Printf("\n\n\n")
}
