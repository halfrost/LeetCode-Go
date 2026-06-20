package leetcode

import (
	"fmt"
	"testing"
)

type question215 struct {
	para215
	ans215
}

// para 是参数
// one 代表第一个参数
type para215 struct {
	one []int
	two int
}

// ans 是答案
// one 代表第一个答案
type ans215 struct {
	one int
}

// clone215 复制切片，避免多个就地改写的解法共享同一底层数组。
func clone215(a []int) []int {
	return append([]int(nil), a...)
}

func Test_Problem215(t *testing.T) {

	qs := []question215{

		{
			para215{[]int{3, 2, 1}, 2},
			ans215{2},
		},

		{
			para215{[]int{3, 2, 1, 5, 6, 4}, 2},
			ans215{5},
		},

		{
			para215{[]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4},
			ans215{4},
		},

		{
			para215{[]int{0, 0, 0, 0, 0}, 2},
			ans215{0},
		},

		{
			para215{[]int{1}, 1},
			ans215{1},
		},

		{
			para215{[]int{3, 2, 3, 1, 2, 4, 5, 5, 6, 7, 7, 8, 2, 3, 1, 1, 1, 10, 11, 5, 6, 2, 4, 7, 8, 5, 6}, 20},
			ans215{2},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 215------------------------\n")

	for _, q := range qs {
		a, p := q.ans215, q.para215
		// findKthLargest（快速选择）、findKthLargest1（排序）、getLeastNumbers 都会
		// 就地改写传入切片，所以各给一份独立副本，互不污染。
		got := findKthLargest(clone215(p.one), p.two)
		got1 := findKthLargest1(clone215(p.one), p.two)
		fmt.Printf("【input】:%v    【output】:%v\n", p.one, got)
		if got != a.one || got1 != a.one {
			t.Fatalf("input %v, k=%v: expected %v, got %v / %v", p.one, p.two, a.one, got, got1)
		}
		// getLeastNumbers 是扩展题（最小的 k 个数），返回前 k 小，校验长度即可。
		if least := getLeastNumbers(clone215(p.one), p.two); len(least) != p.two {
			t.Fatalf("getLeastNumbers(%v, %v) len = %v, want %v", p.one, p.two, len(least), p.two)
		}
	}
	fmt.Printf("\n\n\n")
}
