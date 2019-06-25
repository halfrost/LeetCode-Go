package leetcode

import (
	"fmt"
	"testing"
)

type question401 struct {
	para401
	ans401
}

// para 是参数
// one 代表第一个参数
type para401 struct {
	n int
}

// ans 是答案
// one 代表第一个答案
type ans401 struct {
	one []string
}

func Test_Problem401(t *testing.T) {

	qs := []question401{

		question401{
			para401{1},
			ans401{[]string{"1:00", "2:00", "4:00", "8:00", "0:01", "0:02", "0:04", "0:08", "0:16", "0:32"}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 401------------------------\n")

	for _, q := range qs {
		_, p := q.ans401, q.para401
		fmt.Printf("【input】:%v       【output】:%v\n", p, readBinaryWatch(p.n))
	}
	fmt.Printf("\n\n\n")
}
