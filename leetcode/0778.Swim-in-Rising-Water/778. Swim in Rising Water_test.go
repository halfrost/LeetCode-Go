package leetcode

import (
	"fmt"
	"testing"
)

type question778 struct {
	para778
	ans778
}

// para 是参数
// one 代表第一个参数
type para778 struct {
	grid [][]int
}

// ans 是答案
// one 代表第一个答案
type ans778 struct {
	one int
}

func Test_Problem778(t *testing.T) {

	qs := []question778{
		question778{
			para778{[][]int{[]int{0, 2}, []int{1, 3}}},
			ans778{3},
		},

		question778{
			para778{[][]int{[]int{0, 1, 2, 3, 4}, []int{24, 23, 22, 21, 5}, []int{12, 13, 14, 15, 16}, []int{11, 17, 18, 19, 20}, []int{10, 9, 8, 7, 6}}},
			ans778{16},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 778------------------------\n")

	for _, q := range qs {
		_, p := q.ans778, q.para778
		fmt.Printf("【input】:%v       【output】:%v\n", p, swimInWater(p.grid))
	}
	fmt.Printf("\n\n\n")
}
