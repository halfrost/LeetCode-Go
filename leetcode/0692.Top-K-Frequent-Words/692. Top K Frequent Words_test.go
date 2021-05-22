package leetcode

import (
	"fmt"
	"testing"
)

type question692 struct {
	para692
	ans692
}

// para 是参数
// one 代表第一个参数
type para692 struct {
	words []string
	k     int
}

// ans 是答案
// one 代表第一个答案
type ans692 struct {
	one []string
}

func Test_Problem692(t *testing.T) {

	qs := []question692{

		{
			para692{[]string{"i", "love", "leetcode", "i", "love", "coding"}, 2},
			ans692{[]string{"i", "love"}},
		},

		{
			para692{[]string{"the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"}, 4},
			ans692{[]string{"the", "is", "sunny", "day"}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 692------------------------\n")

	for _, q := range qs {
		_, p := q.ans692, q.para692
		fmt.Printf("【input】:%v       【output】:%v\n", p, topKFrequent(p.words, p.k))
	}
	fmt.Printf("\n\n\n")
}
