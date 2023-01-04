package leetcode

func arrayPairSum(nums []int) int {
	array := [20001]int{}
	for i := 0; i < len(nums); i++ {
		array[nums[i]+10000]++
	}
	flag, sum := true, 0
	for i := 0; i < len(array); i++ {
		for array[i] > 0 {
			if flag {
				sum = sum + i - 10000
			}
			flag = !flag
			array[i]--
		}
	}
	return sum
}
