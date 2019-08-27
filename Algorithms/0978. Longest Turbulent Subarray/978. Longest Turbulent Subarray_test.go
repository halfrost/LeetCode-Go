package leetcode

import (
	"fmt"
	"testing"
)

type question978 struct {
	para978
	ans978
}

// para 是参数
// one 代表第一个参数
type para978 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans978 struct {
	one int
}

func Test_Problem978(t *testing.T) {

	qs := []question978{

		question978{
			para978{[]int{0, 1, 1, 0, 1, 0, 1, 1, 0, 0}},
			ans978{5},
		},

		question978{
			para978{[]int{9, 9}},
			ans978{1},
		},

		question978{
			para978{[]int{9, 4, 2, 10, 7, 8, 8, 1, 9}},
			ans978{5},
		},

		question978{
			para978{[]int{4, 8, 12, 16}},
			ans978{2},
		},

		question978{
			para978{[]int{100}},
			ans978{1},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 978------------------------\n")

	for _, q := range qs {
		_, p := q.ans978, q.para978
		fmt.Printf("【input】:%v       【output】:%v\n", p, maxTurbulenceSize(p.one))
	}
	fmt.Printf("\n\n\n")
}
