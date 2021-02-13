package leetcode

import (
	"fmt"
	"testing"
)

type question1463 struct {
	para1463
	ans1463
}

type para1463 struct {
	grid [][]int
}

type ans1463 struct {
	ans int
}

func Test_Problem1463(t *testing.T) {

	qs := []question1463{

		{
			para1463{[][]int{
				{3, 1, 1},
				{2, 5, 1},
				{1, 5, 5},
				{2, 1, 1},
			}},
			ans1463{24},
		},

		{
			para1463{[][]int{
				{1, 0, 0, 0, 0, 0, 1},
				{2, 0, 0, 0, 0, 3, 0},
				{2, 0, 9, 0, 0, 0, 0},
				{0, 3, 0, 5, 4, 0, 0},
				{1, 0, 2, 3, 0, 0, 6},
			}},
			ans1463{28},
		},
		{
			para1463{[][]int{
				{1, 0, 0, 3},
				{0, 0, 0, 3},
				{0, 0, 3, 3},
				{9, 0, 3, 3},
			}},
			ans1463{22},
		},
		{
			para1463{[][]int{
				{1, 1},
				{1, 1},
			}},
			ans1463{4},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1463------------------------\n")

	for _, q := range qs {
		_, p := q.ans1463, q.para1463
		fmt.Printf("【input】:%v      ", p)
		fmt.Printf("【output】:%v      \n", cherryPickup(p.grid))
	}
	fmt.Printf("\n\n\n")
}
