package leetcode

import (
	"fmt"
	"testing"
)

type question748 struct {
	para748
	ans748
}

// para 是参数
// one 代表第一个参数
type para748 struct {
	c string
	w []string
}

// ans 是答案
// one 代表第一个答案
type ans748 struct {
	one string
}

func Test_Problem748(t *testing.T) {

	qs := []question748{

		question748{
			para748{"1s3 PSt", []string{"step", "steps", "stripe", "stepple"}},
			ans748{"steps"},
		},

		question748{
			para748{"1s3 456", []string{"looks", "pest", "stew", "show"}},
			ans748{"pest"},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 748------------------------\n")

	for _, q := range qs {
		_, p := q.ans748, q.para748
		fmt.Printf("【input】:%v       【output】:%v\n", p, shortestCompletingWord(p.c, p.w))
	}
	fmt.Printf("\n\n\n")
}
