package leetcode

import (
	"fmt"
	"testing"
)

type question509 struct {
	para509
	ans509
}

// para 是参数
// one 代表第一个参数
type para509 struct {
	one int
}

// ans 是答案
// one 代表第一个答案
type ans509 struct {
	one int
}

func Test_Problem509(t *testing.T) {

	qs := []question509{

		question509{
			para509{1},
			ans509{1},
		},

		question509{
			para509{2},
			ans509{1},
		},

		question509{
			para509{3},
			ans509{2},
		},
		// 如需多个测试，可以复制上方元素。
	}

	fmt.Printf("------------------------Leetcode Problem 509------------------------\n")

	for _, q := range qs {
		_, p := q.ans509, q.para509
		fmt.Printf("【input】:%v       【output】:%v\n", p, fib(p.one))
	}
	fmt.Printf("\n\n\n")
}
