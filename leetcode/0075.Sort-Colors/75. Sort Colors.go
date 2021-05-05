package leetcode

// 在已知排序的数字个数时，可以使用刷油漆式的排序方式
// 第一个指针最大，第二指针次之，依次。。。
// [2 0 2 1 1 1 0]
// -> [2 2 2 2 2 2 2] 先全填上2
// -> [1 1 1 1 1 2 2] 统计下0和1的个数之和(作为数字1的右侧边界)，然后填上1
// -> [0 0 1 1 1 2 2] 统计下0的个数（作为数字0的右侧边界），然后填上0
func sortColors(nums []int) {
	one, zero := 0, 0
	for i, num := range nums {
		tmp := num
		nums[i] = 2
		if tmp <= 1 {
			nums[one] = 1
			one++
		}
		if tmp == 0 {
			nums[zero] = 0
			zero++
		}
	}
}
