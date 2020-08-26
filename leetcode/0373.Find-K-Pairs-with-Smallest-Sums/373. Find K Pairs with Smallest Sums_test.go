package leetcode

import (
	"fmt"
	"testing"
)

type question373 struct {
	para373
	ans373
}

// para 是参数
// one 代表第一个参数
type para373 struct {
	nums1 []int
	nums2 []int
	k     int
}

// ans 是答案
// one 代表第一个答案
type ans373 struct {
	one [][]int
}

func Test_Problem373(t *testing.T) {

	qs := []question373{

		{
			para373{[]int{1, 7, 11}, []int{2, 4, 6}, 3},
			ans373{[][]int{{1, 2}, {1, 4}, {1, 6}}},
		},

		{
			para373{[]int{1, 1, 2}, []int{1, 2, 3}, 2},
			ans373{[][]int{{1, 1}, {1, 1}}},
		},

		{
			para373{[]int{1, 2}, []int{3}, 3},
			ans373{[][]int{{1, 3}, {2, 3}}},
		},

		{
			para373{[]int{1, 2, 4, 5, 6}, []int{3, 5, 7, 9}, 3},
			ans373{[][]int{{1, 3}, {2, 3}, {1, 5}}},
		},

		{
			para373{[]int{1, 1, 2}, []int{1, 2, 3}, 10},
			ans373{[][]int{{1, 1}, {1, 1}, {2, 1}, {1, 2}, {1, 2}, {2, 2}, {1, 3}, {1, 3}, {2, 3}}},
		},

		{
			para373{[]int{-10, -4, 0, 0, 6}, []int{3, 5, 6, 7, 8, 100}, 10},
			ans373{[][]int{{-10, 3}, {-10, 5}, {-10, 6}, {-10, 7}, {-10, 8}, {-4, 3}, {-4, 5}, {-4, 6}, {0, 3}, {0, 3}}},
		},

		{
			para373{[]int{0, 0, 0, 0, 0}, []int{-3, 22, 35, 56, 76}, 22},
			ans373{[][]int{{0, -3}, {0, -3}, {0, -3}, {0, -3}, {0, -3}, {0, 22}, {0, 22}, {0, 22}, {0, 22}, {0, 22}, {0, 35}, {0, 35}, {0, 35}, {0, 35}, {0, 35}, {0, 56}, {0, 56}, {0, 56}, {0, 56}, {0, 56}, {0, 76}, {0, 76}}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 373------------------------\n")

	for _, q := range qs {
		_, p := q.ans373, q.para373
		fmt.Printf("【input】:%v       【output】:%v\n", p, kSmallestPairs(p.nums1, p.nums2, p.k))
	}
	fmt.Printf("\n\n\n")
}
