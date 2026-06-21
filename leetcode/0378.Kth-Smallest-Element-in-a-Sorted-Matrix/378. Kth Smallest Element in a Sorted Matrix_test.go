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

		{
			para378{[][]int{{1, 5, 9}, {10, 11, 13}, {12, 13, 15}}, 8},
			ans378{13},
		},

		{
			para378{[][]int{{1, 5, 7}, {11, 12, 13}, {12, 13, 15}}, 3},
			ans378{9},
		},

		{
			para378{[][]int{{1, 3, 5}, {2, 4, 6}}, 2},
			ans378{2},
		},

		{
			para378{[][]int{}, 1},
			ans378{0},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 378------------------------\n")

	for _, q := range qs {
		_, p := q.ans378, q.para378
		got := kthSmallest3781(p.matrix, p.k)
		if len(p.matrix) != 0 && len(p.matrix[0]) != 0 {
			primary := kthSmallest378(p.matrix, p.k)
			if got != primary {
				t.Fatalf("kthSmallest3781(%v, %d) = %d, kthSmallest378 = %d", p.matrix, p.k, got, primary)
			}
			fmt.Printf("【input】:%v       【output】:%v\n", p, primary)
		} else {
			fmt.Printf("【input】:%v       【output】:%v\n", p, got)
		}
	}
	fmt.Printf("\n\n\n")
}
