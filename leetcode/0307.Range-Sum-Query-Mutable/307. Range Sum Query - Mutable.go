package leetcode

import "github.com/halfrost/LeetCode-Go/template"

// NumArray define
type NumArray struct {
	st *template.SegmentTree
}

// Constructor307 define
func Constructor307(nums []int) NumArray {
	st := template.SegmentTree{}
	st.Init(nums, func(i, j int) int {
		return i + j
	})
	return NumArray{st: &st}
}

// Update define
func (this *NumArray) Update(i int, val int) {
	this.st.Update(i, val)
}

// SumRange define
func (this *NumArray) SumRange(i int, j int) int {
	return this.st.Query(i, j)
}

//解法二 prefixSum，sumRange 时间复杂度 O(1)

// // NumArray define
// type NumArray307 struct {
// 	prefixSum []int
// 	data   []int
// }

// // Constructor307 define
// func Constructor307(nums []int) NumArray307 {
// 	data := make([]int, len(nums))
// 	for i := 0; i < len(nums); i++ {
// 		data[i] = nums[i]
// 	}
// 	for i := 1; i < len(nums); i++ {
// 		nums[i] += nums[i-1]
// 	}
// 	return NumArray307{prefixSum: nums, data: data}
// }

// // Update define
// func (this *NumArray307) Update(i int, val int) {
// 	this.data[i] = val
// 	this.prefixSum[0] = this.data[0]
// 	for i := 1; i < len(this.data); i++ {
// 		this.prefixSum[i] = this.prefixSum[i-1] + this.data[i]
// 	}
// }

// // SumRange define
// func (this *NumArray307) SumRange(i int, j int) int {
// 	if i > 0 {
// 		return this.prefixSum[j] - this.prefixSum[i-1]
// 	}
// 	return this.prefixSum[j]
// }

// 解法三 树状数组
// type NumArray struct {
// 	bit  template.BinaryIndexedTree
// 	data []int
// }

// // Constructor define
// func Constructor307(nums []int) NumArray {
// 	bit := template.BinaryIndexedTree{}
// 	bit.InitWithNums(nums)
// 	return NumArray{bit: bit, data: nums}
// }

// // Update define
// func (this *NumArray) Update(i int, val int) {
// 	this.bit.Add(i+1, val-this.data[i])
// 	this.data[i] = val
// }

// // SumRange define
// func (this *NumArray) SumRange(i int, j int) int {
// 	return this.bit.Query(j+1) - this.bit.Query(i)
// }

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(i,val);
 * param_2 := obj.SumRange(i,j);
 */
