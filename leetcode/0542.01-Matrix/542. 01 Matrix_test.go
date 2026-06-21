package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

type question542 struct {
	para542
	ans542
}

// para 是参数
// one 代表第一个参数
type para542 struct {
	one [][]int
}

// ans 是答案
// one 代表第一个答案
type ans542 struct {
	one [][]int
}

func clone542(matrix [][]int) [][]int {
	res := make([][]int, len(matrix))
	for i, row := range matrix {
		res[i] = make([]int, len(row))
		copy(res[i], row)
	}
	return res
}

func Test_Problem542(t *testing.T) {

	qs := []question542{

		{
			para542{[][]int{}},
			ans542{[][]int{}},
		},

		{
			para542{[][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}},
			ans542{[][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}},
		},

		{
			para542{[][]int{{0, 0, 0}, {0, 1, 0}, {1, 1, 1}}},
			ans542{[][]int{{0, 0, 0}, {0, 1, 0}, {1, 2, 1}}},
		},

		{
			para542{[][]int{{1, 1, 1, 1, 1, 1}, {1, 1, 1, 1, 1, 1}, {1, 1, 0, 0, 1, 1}, {1, 1, 0, 0, 1, 1}, {1, 1, 1, 1, 1, 1}, {1, 1, 1, 1, 1, 1}}},
			ans542{[][]int{{4, 3, 2, 2, 3, 4}, {3, 2, 1, 1, 2, 3}, {2, 1, 0, 0, 1, 2}, {2, 1, 0, 0, 1, 2}, {3, 2, 1, 1, 2, 3}, {4, 3, 2, 2, 3, 4}}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 542------------------------\n")

	for _, q := range qs {
		a, p := q.ans542, q.para542
		dp := updateMatrixDP(clone542(p.one))
		fmt.Printf("【input】:%v       【output】:%v\n", p, dp)
		if !reflect.DeepEqual(dp, a.one) {
			t.Fatalf("updateMatrixDP(%v) = %v, want %v", p.one, dp, a.one)
		}
		bfs := updateMatrixBFS(clone542(p.one))
		if len(p.one) > 0 && !reflect.DeepEqual(bfs, a.one) {
			t.Fatalf("updateMatrixBFS(%v) = %v, want %v", p.one, bfs, a.one)
		}
		dfs := updateMatrixDFS(clone542(p.one))
		if len(p.one) > 0 && !reflect.DeepEqual(dfs, a.one) {
			t.Fatalf("updateMatrixDFS(%v) = %v, want %v", p.one, dfs, a.one)
		}
	}
	fmt.Printf("\n\n\n")
}
