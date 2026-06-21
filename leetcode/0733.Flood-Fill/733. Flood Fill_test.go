package leetcode

import (
	"fmt"
	"testing"
)

type question733 struct {
	para733
	ans733
}

// para 是参数
// one 代表第一个参数
type para733 struct {
	one [][]int
	sr  int
	sc  int
	c   int
}

// ans 是答案
// one 代表第一个答案
type ans733 struct {
	one [][]int
}

func Test_Problem733(t *testing.T) {

	qs := []question733{

		{
			para733{[][]int{
				{1, 1, 1},
				{1, 1, 0},
				{1, 0, 1},
			}, 1, 1, 2},
			ans733{[][]int{
				{2, 2, 2},
				{2, 2, 0},
				{2, 0, 1},
			}},
		},

		// newColor == color, floodFill returns image unchanged
		{
			para733{[][]int{
				{0, 0, 0},
				{0, 1, 1},
			}, 1, 1, 1},
			ans733{[][]int{
				{0, 0, 0},
				{0, 1, 1},
			}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 733------------------------\n")

	for _, q := range qs {
		a, p := q.ans733, q.para733
		got := floodFill(p.one, p.sr, p.sc, p.c)
		if !equal733(got, a.one) {
			t.Fatalf("floodFill(%v, %d, %d, %d) = %v, want %v", p.one, p.sr, p.sc, p.c, got, a.one)
		}
		fmt.Printf("【input】:%v       【output】:%v\n", p, got)
	}

	// Cover the dfs733 early-return guard (image[x][y] == newColor) by
	// calling dfs733 directly on a cell that already has newColor.
	guard := [][]int{
		{5, 0},
		{0, 0},
	}
	dfs733(guard, 0, 0, 5)
	if guard[0][0] != 5 {
		t.Fatalf("dfs733 guard mutated cell: %v", guard)
	}

	fmt.Printf("\n\n\n")
}

func equal733(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if len(a[i]) != len(b[i]) {
			return false
		}
		for j := range a[i] {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}
