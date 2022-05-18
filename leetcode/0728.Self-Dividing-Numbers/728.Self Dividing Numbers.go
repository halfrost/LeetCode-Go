package leetcode

func selfDividingNumbers(left int, right int) []int {
	var ans []int
	for num := left; num <= right; num++ {
		if selfDividingNum(num) {
			ans = append(ans, num)
		}
	}
	return ans
}

func selfDividingNum(num int) bool {
	for d := num; d > 0; d = d / 10 {
		reminder := d % 10
		if reminder == 0 {
			return false
		}
		if num%reminder != 0 {
			return false
		}
	}
	return true
}
