package leetcode

import (
	"fmt"
	"testing"
)

type question1705 struct {
	para1705
	ans1705
}

// para 是参数
type para1705 struct {
	apples []int
	days   []int
}

// ans 是答案
type ans1705 struct {
	ans int
}

func Test_Problem1705(t *testing.T) {

	qs := []question1705{

		{
			para1705{[]int{1, 2, 3, 5, 2}, []int{3, 2, 1, 4, 2}},
			ans1705{7},
		},

		{
			para1705{[]int{3, 0, 0, 0, 0, 2}, []int{3, 0, 0, 0, 0, 2}},
			ans1705{5},
		},

		{
			para1705{[]int{10}, []int{2}},
			ans1705{2},
		},

		{
			para1705{[]int{5, 1}, []int{2, 1}},
			ans1705{2},
		},

		{
			para1705{[]int{2}, []int{5}},
			ans1705{2},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1705------------------------\n")

	for _, q := range qs {
		a, p := q.ans1705, q.para1705
		got := eatenApples(p.apples, p.days)
		fmt.Printf("【input】:%v    【output】:%v\n", p, got)
		if got != a.ans {
			t.Fatalf("input %v expected %d got %d", p, a.ans, got)
		}
	}
	fmt.Printf("\n\n\n")
}
