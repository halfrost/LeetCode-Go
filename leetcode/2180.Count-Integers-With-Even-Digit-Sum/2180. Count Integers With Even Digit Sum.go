package leetcode

func countEven(num int) int {
	count := 0
	for i := 1; i <= num; i++ {
		if addSum(i)%2 == 0 {
			count++
		}
	}
	return count
}

func addSum(num int) int {
	sum := 0
	tmp := num
	for tmp != 0 {
		sum += tmp % 10
		tmp = tmp / 10
	}
	return sum
}
