package leetcode

import (
	"fmt"
	"testing"
)

type question480 struct {
	para480
	ans480
}

// para 是参数
// one 代表第一个参数
type para480 struct {
	one []int
	k   int
}

// ans 是答案
// one 代表第一个答案
type ans480 struct {
	one []int
}

func Test_Problem480(t *testing.T) {

	qs := []question480{

		{
			para480{[]int{1, 3, -1, -3, 5, 3, 6, 7}, 3},
			ans480{[]int{1, -1, -1, 3, 5, 6}},
		},
		// 偶数窗口，覆盖 getMedian / medianSlidingWindow1 的 k%2==0 分支
		{
			para480{[]int{1, 3, -1, -3, 5, 3, 6, 7}, 4},
			ans480{[]int{0, 1, 1, 4, 5}},
		},
		// 覆盖 MaxHeapR.Top 中删除堆顶去重循环
		{
			para480{[]int{5, 2, 2, 7, 3, 7, 9, 0, 2, 3}, 3},
			ans480{[]int{2, 2, 3, 7, 7, 7, 2, 2}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 480------------------------\n")

	for _, q := range qs {
		_, p := q.ans480, q.para480
		got1 := medianSlidingWindow1(p.one, p.k)
		got := medianSlidingWindow(p.one, p.k)
		fmt.Printf("【input】:%v       【output】:%v\n", p, got1)
		if len(got) != len(got1) {
			t.Fatalf("two solutions length mismatch: %v vs %v", got, got1)
		}
		for i := range got {
			if got[i] != got1[i] {
				t.Fatalf("two solutions mismatch at %d: %v vs %v", i, got, got1)
			}
		}
	}

	// 覆盖 removeFromWindow 中元素不存在时的返回分支
	w := getWindowList([]int{1, 2, 3}, 3)
	if removeFromWindow(w, 99).Len() != 3 {
		t.Fatalf("removeFromWindow should keep list intact when value missing")
	}

	fmt.Printf("\n\n\n")
}
