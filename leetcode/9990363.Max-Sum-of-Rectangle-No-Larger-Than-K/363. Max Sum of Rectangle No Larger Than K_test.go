package leetcode

import (
	"fmt"
	"testing"
)

type question363 struct {
	para363
	ans363
}

// para 是参数
// one 代表第一个参数
type para363 struct {
	one [][]int
	k   int
}

// ans 是答案
// one 代表第一个答案
type ans363 struct {
	one int
}

// clone363 深拷贝矩阵。maxSumSubmatrix 和 maxSumSubmatrix1 都会就地把每行
// 改写成前缀和，所以两个解法必须各拿一份独立副本，否则第二次调用会在被改写过的
// 矩阵上再算一次前缀和，导致 maxSumSubmatrix1 永远找不到目标而死循环（卡住）。
func clone363(matrix [][]int) [][]int {
	res := make([][]int, len(matrix))
	for i := range matrix {
		res[i] = append([]int(nil), matrix[i]...)
	}
	return res
}

func Test_Problem363(t *testing.T) {

	qs := []question363{

		{
			para363{[][]int{{1, 0, 1}, {0, -2, 3}}, 2},
			ans363{2},
		},

		{
			para363{[][]int{{2, 2, -1}}, 0},
			ans363{-1},
		},

		{
			para363{[][]int{{5, -4, -3, 4}, {-3, -4, 4, 5}, {5, 1, 5, -4}}, 8},
			ans363{8},
		},

		{
			para363{[][]int{{5, -4, -3, 4}, {-3, -4, 4, 5}, {5, 1, 5, -4}}, 3},
			ans363{2},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 363------------------------\n")

	for _, q := range qs {
		a, p := q.ans363, q.para363
		// 两个解法各用一份独立副本，避免就地前缀和互相污染（详见 clone363 注释）。
		got := maxSumSubmatrix(clone363(p.one), p.k)
		got1 := maxSumSubmatrix1(clone363(p.one), p.k)
		fmt.Printf("【input】:%v       【output】:%v\n", p, got)
		if got != a.one || got1 != a.one {
			t.Fatalf("input %v, k=%v: expected %v, sol1 got %v, sol2 got %v", p.one, p.k, a.one, got, got1)
		}
	}
	fmt.Printf("\n\n\n")
}
