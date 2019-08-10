package leetcode

import (
	"fmt"
	"testing"
)

type question463 struct {
	para463
	ans463
}

// para 是参数
// one 代表第一个参数
type para463 struct {
	one [][]int
}

// ans 是答案
// one 代表第一个答案
type ans463 struct {
	one int
}

func Test_Problem463(t *testing.T) {

	qs := []question463{

		question463{
			para463{[][]int{[]int{0, 1, 0, 0}, []int{1, 1, 1, 0}, []int{0, 1, 0, 0}, []int{1, 1, 0, 0}}},
			ans463{16},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 463------------------------\n")

	for _, q := range qs {
		_, p := q.ans463, q.para463
		fmt.Printf("【input】:%v       【output】:%v\n", p, islandPerimeter(p.one))
	}
	fmt.Printf("\n\n\n")
}
