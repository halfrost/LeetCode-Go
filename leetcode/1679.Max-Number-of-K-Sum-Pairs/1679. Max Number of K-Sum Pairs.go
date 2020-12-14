package leetcode

// 解法一 优化版
func maxOperations(nums []int, k int) int {
	counter, res := make(map[int]int), 0
	for _, n := range nums {
		counter[n]++
	}
	if (k & 1) == 0 {
		res += counter[k>>1] >> 1
		// 能够由 2 个相同的数构成 k 的组合已经都排除出去了，剩下的一个单独的也不能组成 k 了
		// 所以这里要把它的频次置为 0 。如果这里不置为 0，下面代码判断逻辑还需要考虑重复使用数字的情况
		counter[k>>1] = 0
	}
	for num, freq := range counter {
		if num <= k/2 {
			remain := k - num
			if counter[remain] < freq {
				res += counter[remain]
			} else {
				res += freq
			}
		}
	}
	return res
}

// 解法二
func maxOperations_(nums []int, k int) int {
	counter, res := make(map[int]int), 0
	for _, num := range nums {
		counter[num]++
		remain := k - num
		if num == remain {
			if counter[num] >= 2 {
				res++
				counter[num] -= 2
			}
		} else {
			if counter[remain] > 0 {
				res++
				counter[remain]--
				counter[num]--
			}
		}
	}
	return res
}
