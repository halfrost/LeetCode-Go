package leetcode

func maxRotateFunction(nums []int) int {
	n := len(nums)
	var sum, f int
	for i, num := range nums {
		sum += num
		f += i * num // F(0)
	}
	ans := f
	for i := 1; i < n; i++ {
		f += sum - n*nums[n-i] // F(i) = F(i-1) + sum - n*nums[n-i]
		if f > ans {
			ans = f
		}
	}
	return ans
}
