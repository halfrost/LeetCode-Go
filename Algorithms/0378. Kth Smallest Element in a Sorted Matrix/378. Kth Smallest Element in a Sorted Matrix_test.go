package leetcode

import (
	"fmt"
	"testing"
)

type question378 struct {
	para378
	ans378
}

// para 是参数
// one 代表第一个参数
type para378 struct {
	matrix [][]int
	k      int
}

// ans 是答案
// one 代表第一个答案
type ans378 struct {
	one int
}

func Test_Problem378(t *testing.T) {

	qs := []question378{

		question378{
			para378{[][]int{[]int{1, 5, 9}, []int{10, 11, 13}, []int{12, 13, 15}}, 8},
			ans378{13},
		},

		question378{
			para378{[][]int{[]int{1, 5, 7}, []int{11, 12, 13}, []int{12, 13, 15}}, 3},
			ans378{9},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 378------------------------\n")

	for _, q := range qs {
		_, p := q.ans378, q.para378
		fmt.Printf("【input】:%v       【output】:%v\n", p, kthSmallest378(p.matrix, p.k))
	}
	fmt.Printf("\n\n\n")
}
