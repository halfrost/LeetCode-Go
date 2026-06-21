package leetcode

import (
	"fmt"
	"testing"
)

type question4 struct {
	para4
	ans4
}

// para 是参数
// one 代表第一个参数
type para4 struct {
	nums1 []int
	nums2 []int
}

// ans 是答案
// one 代表第一个答案
type ans4 struct {
	one float64
}

func Test_Problem4(t *testing.T) {

	qs := []question4{

		{
			para4{[]int{1, 3}, []int{2}},
			ans4{2.0},
		},

		{
			para4{[]int{1, 2}, []int{3, 4}},
			ans4{2.5},
		},

		// nums1 长度大于 nums2，触发首行 swap 递归
		{
			para4{[]int{1, 2, 3, 4}, []int{5}},
			ans4{3.0},
		},

		// nums1Mid == 0 且为偶数长度，midLeft 取 nums2，midRight 取 min
		{
			para4{[]int{3, 4}, []int{1, 2}},
			ans4{2.5},
		},

		// nums2Mid == 0 分支：nums1 全部在左侧之前
		{
			para4{[]int{4, 5, 6}, []int{1, 2, 3}},
			ans4{3.5},
		},

		// nums1Mid == len(nums1) 分支，midRight 取 nums2
		{
			para4{[]int{1, 2}, []int{3, 4, 5, 6}},
			ans4{3.5},
		},

		// 相等元素，触发 max/min 的 a == b（非 a > b）路径
		{
			para4{[]int{2, 2}, []int{2, 2}},
			ans4{2.0},
		},

		// 奇数总长度
		{
			para4{[]int{1, 3, 5}, []int{2, 4}},
			ans4{3.0},
		},

		// 触发 min 的 a > b 分支（右侧取 nums2 的较小值）
		{
			para4{[]int{1, 4}, []int{2, 3}},
			ans4{2.5},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 4------------------------\n")

	for _, q := range qs {
		a, p := q.ans4, q.para4
		got := findMedianSortedArrays(p.nums1, p.nums2)
		if got != a.one {
			t.Fatalf("findMedianSortedArrays(%v, %v) = %v, want %v", p.nums1, p.nums2, got, a.one)
		}
		fmt.Printf("【input】:%v       【output】:%v\n", p, got)
	}
	fmt.Printf("\n\n\n")
}
