package leetcode

import (
	"fmt"
	"testing"
)

type question802 struct {
	para802
	ans802
}

// para 是参数
// one 代表第一个参数
type para802 struct {
	graph [][]int
}

// ans 是答案
// one 代表第一个答案
type ans802 struct {
	one []int
}

func Test_Problem802(t *testing.T) {

	qs := []question802{

		{
			para802{[][]int{{1, 2}, {2, 3}, {5}, {0}, {5}, {}, {}}},
			ans802{[]int{2, 4, 5, 6}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 802------------------------\n")

	for _, q := range qs {
		_, p := q.ans802, q.para802
		fmt.Printf("【input】:%v       【output】:%v\n", p, eventualSafeNodes(p.graph))
	}
	fmt.Printf("\n\n\n")
}
