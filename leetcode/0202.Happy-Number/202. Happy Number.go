package leetcode

func isHappy(n int) bool {
	record := map[int]int{}
	for n != 1 {
		record[n] = n
		n = getSquareOfDigits(n)
		for _, previous := range record {
			if n == previous {
				return false
			}
		}
	}
	return true
}

func getSquareOfDigits(n int) int {
	squareOfDigits := 0
	temporary := n
	for temporary != 0 {
		remainder := temporary % 10
		squareOfDigits += remainder * remainder
		temporary /= 10
	}
	return squareOfDigits
}
