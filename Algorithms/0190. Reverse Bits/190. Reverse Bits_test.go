package leetcode

import (
	"fmt"
	"testing"
)

type question190 struct {
	para190
	ans190
}

// para 是参数
// one 代表第一个参数
type para190 struct {
	one uint32
}

// ans 是答案
// one 代表第一个答案
type ans190 struct {
	one uint32
}

func Test_Problem190(t *testing.T) {

	qs := []question190{

		question190{
			para190{43261596},
			ans190{964176192},
		},

		question190{
			para190{4294967293},
			ans190{3221225471},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 190------------------------\n")

	for _, q := range qs {
		_, p := q.ans190, q.para190
		fmt.Printf("【input】:%v       【output】:%v\n", p, reverseBits(p.one))
	}
	fmt.Printf("\n\n\n")
}
