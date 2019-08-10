package leetcode

import (
	"fmt"
	"testing"
)

type question447 struct {
	para447
	ans447
}

// para 是参数
// one 代表第一个参数
type para447 struct {
	one [][]int
}

// ans 是答案
// one 代表第一个答案
type ans447 struct {
	one int
}

func Test_Problem447(t *testing.T) {

	qs := []question447{

		question447{
			para447{[][]int{[]int{0, 0}, []int{1, 0}, []int{2, 0}}},
			ans447{2},
		},

		// 如需多个测试，可以复制上方元素。
	}

	fmt.Printf("------------------------Leetcode Problem 447------------------------\n")

	for _, q := range qs {
		_, p := q.ans447, q.para447
		fmt.Printf("【input】:%v       【output】:%v\n", p, numberOfBoomerangs(p.one))
	}
	fmt.Printf("\n\n\n")
}
