package leetcode

import (
	"fmt"
	"testing"
)

type question357 struct {
	para357
	ans357
}

// para 是参数
// one 代表第一个参数
type para357 struct {
	one int
}

// ans 是答案
// one 代表第一个答案
type ans357 struct {
	one int
}

func Test_Problem357(t *testing.T) {

	qs := []question357{

		question357{
			para357{1},
			ans357{10},
		},

		question357{
			para357{2},
			ans357{91},
		},

		question357{
			para357{3},
			ans357{739},
		},

		question357{
			para357{4},
			ans357{5275},
		},

		question357{
			para357{5},
			ans357{32491},
		},

		question357{
			para357{6},
			ans357{168571},
		},
		// 如需多个测试，可以复制上方元素。
	}

	fmt.Printf("------------------------Leetcode Problem 357------------------------\n")

	for _, q := range qs {
		_, p := q.ans357, q.para357
		fmt.Printf("【input】:%v       【output】:%v\n", p, countNumbersWithUniqueDigits(p.one))
	}
	fmt.Printf("\n\n\n")
}
