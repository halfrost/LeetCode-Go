package leetcode

import (
	"github.com/halfrost/LeetCode-Go/template"
)

//解法一 线段树，sumRange 时间复杂度 O(1)

// NumArray define
type NumArray struct {
	st *template.SegmentTree
}

// Constructor303 define
func Constructor303(nums []int) NumArray {
	st := template.SegmentTree{}
	st.Init(nums, func(i, j int) int {
		return i + j
	})
	return NumArray{st: &st}
}

// SumRange define
func (ma *NumArray) SumRange(i int, j int) int {
	return ma.st.Query(i, j)
}

//解法二 prefixSum，sumRange 时间复杂度 O(1)

// // NumArray define
// type NumArray struct {
// 	prefixSum []int
// }

// // Constructor303 define
// func Constructor303(nums []int) NumArray {
// 	for i := 1; i < len(nums); i++ {
// 		nums[i] += nums[i-1]
// 	}
// 	return NumArray{prefixSum: nums}
// }

// // SumRange define
// func (this *NumArray) SumRange(i int, j int) int {
// 	if i > 0 {
// 		return this.prefixSum[j] - this.prefixSum[i-1]
// 	}
// 	return this.prefixSum[j]
// }

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */
