package leetcode

import (
	"fmt"
	"testing"
)

type question387 struct {
	para387
	ans387
}

// para 是参数
// one 代表第一个参数
type para387 struct {
	n string
}

// ans 是答案
// one 代表第一个答案
type ans387 struct {
	one int
}

func Test_Problem387(t *testing.T) {

	qs := []question387{

		question387{
			para387{"leetcode"},
			ans387{0},
		},

		question387{
			para387{"loveleetcode"},
			ans387{2},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 387------------------------\n")

	for _, q := range qs {
		_, p := q.ans387, q.para387
		fmt.Printf("【input】:%v       【output】:%v\n", p, firstUniqChar(p.n))
	}
	fmt.Printf("\n\n\n")
}
