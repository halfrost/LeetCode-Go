package leetcode

import (
	"fmt"
	"testing"
)

type question1436 struct {
	para1436
	ans1436
}

type para1436 struct {
	grid [][]int
}

type ans1436 struct {
	ans int
}

func Test_Problem1436(t *testing.T) {

	qs := []question1436{

		{
			para1436{[][]int{
				{3, 1, 1},
				{2, 5, 1},
				{1, 5, 5},
				{2, 1, 1},
			}},
			ans1436{24},
		},

		{
			para1436{[][]int{
				{1, 0, 0, 0, 0, 0, 1},
				{2, 0, 0, 0, 0, 3, 0},
				{2, 0, 9, 0, 0, 0, 0},
				{0, 3, 0, 5, 4, 0, 0},
				{1, 0, 2, 3, 0, 0, 6},
			}},
			ans1436{28},
		},
		{
			para1436{[][]int{
				{1, 0, 0, 3},
				{0, 0, 0, 3},
				{0, 0, 3, 3},
				{9, 0, 3, 3},
			}},
			ans1436{22},
		},
		{
			para1436{[][]int{
				{1, 1},
				{1, 1},
			}},
			ans1436{4},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1436------------------------\n")

	for _, q := range qs {
		_, p := q.ans1436, q.para1436
		fmt.Printf("【input】:%v      ", p)
		fmt.Printf("【output】:%v      \n", cherryPickup(p.grid))
	}
	fmt.Printf("\n\n\n")
}
