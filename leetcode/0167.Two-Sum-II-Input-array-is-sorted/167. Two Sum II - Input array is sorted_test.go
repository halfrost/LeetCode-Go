package leetcode

import (
	"fmt"
	"testing"
)

type question167 struct {
	para167
	ans167
}

// para 是参数
// one 代表第一个参数
type para167 struct {
	one []int
	two int
}

// ans 是答案
// one 代表第一个答案
type ans167 struct {
	one []int
}

func Test_Problem167(t *testing.T) {

	qs := []question167{

		{
			para167{[]int{2, 7, 11, 15}, 9},
			ans167{[]int{1, 2}},
		},
		// 命中 j-- 分支：初始 numbers[i]+numbers[j] > target
		{
			para167{[]int{2, 3, 4}, 6},
			ans167{[]int{1, 3}},
		},
		// 无解：覆盖 return nil 分支
		{
			para167{[]int{1, 2, 3}, 100},
			ans167{nil},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 167------------------------\n")

	for _, q := range qs {
		a, p := q.ans167, q.para167
		got := twoSum167(p.one, p.two)
		fmt.Printf("【input】:%v    【output】:%v\n", p.one, got)
		if fmt.Sprintf("%v", got) != fmt.Sprintf("%v", a.one) {
			t.Fatalf("twoSum167(%v, %v) = %v, want %v", p.one, p.two, got, a.one)
		}
		got1 := twoSum167_1(p.one, p.two)
		if fmt.Sprintf("%v", got1) != fmt.Sprintf("%v", a.one) {
			t.Fatalf("twoSum167_1(%v, %v) = %v, want %v", p.one, p.two, got1, a.one)
		}
	}
	fmt.Printf("\n\n\n")
}
