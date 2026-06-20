package leetcode

import (
	"fmt"
	"testing"
)

type question1203 struct {
	para1203
	ans1203
}

// para 是参数
// one 代表第一个参数
type para1203 struct {
	n           int
	m           int
	group       []int
	beforeItems [][]int
}

// ans 是答案
// one 代表第一个答案
type ans1203 struct {
	one []int
}

// clone1203 复制一份 group 切片，避免两个解法共享并就地改写同一底层数组。
func clone1203(a []int) []int {
	return append([]int(nil), a...)
}

func Test_Problem1203(t *testing.T) {

	qs := []question1203{

		{
			para1203{8, 2, []int{-1, -1, 1, 0, 0, 1, 0, -1}, [][]int{{}, {6}, {5}, {6}, {3, 6}, {}, {}, {}}},
			ans1203{[]int{6, 3, 4, 5, 2, 0, 7, 1}},
		},

		{
			para1203{8, 2, []int{-1, -1, 1, 0, 0, 1, 0, -1}, [][]int{{}, {6}, {5}, {6}, {3}, {}, {4}, {}}},
			ans1203{[]int{}},
		},

		// 祖先项目本身没有分组（group 为 -1），覆盖按 ancestor 自身作组的分支
		{
			para1203{3, 1, []int{-1, 0, 0}, [][]int{{}, {0}, {}}},
			ans1203{[]int{0, 1, 2}},
		},

		// 组与组之间形成环，无法完成组间拓扑排序，返回空
		{
			para1203{2, 2, []int{0, 1}, [][]int{{1}, {0}}},
			ans1203{[]int{}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1203------------------------\n")

	for _, q := range qs {
		a, p := q.ans1203, q.para1203
		// sortItems1 会就地把 group 里的 -1 改写成 m+i，必须给每个解法各传一份
		// group 的独立副本，否则第二个解法会读到被污染的 group 而数组越界 panic。
		got1 := sortItems1(p.n, p.m, clone1203(p.group), p.beforeItems)
		got := sortItems(p.n, p.m, clone1203(p.group), p.beforeItems)
		fmt.Printf("【input】:%v       【output】:%v\n", p, got1)
		// 拓扑排序的合法顺序不唯一，这里只校验长度：有解时长度为 n，无解时为空。
		if len(got) != len(a.one) || len(got1) != len(a.one) {
			t.Fatalf("input %v: expected len %v, sortItems got %v, sortItems1 got %v", p, len(a.one), got, got1)
		}
	}
	fmt.Printf("\n\n\n")
}
