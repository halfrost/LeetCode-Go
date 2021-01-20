package leetcode

func decrypt(code []int, k int) []int {
	if k == 0 {
		for i := 0; i < len(code); i++ {
			code[i] = 0
		}
		return code
	}
	count, sum, res := k, 0, make([]int, len(code))
	if k > 0 {
		for i := 0; i < len(code); i++ {
			for j := i + 1; j < len(code); j++ {
				if count == 0 {
					break
				}
				sum += code[j]
				count--
			}
			if count > 0 {
				for j := 0; j < len(code); j++ {
					if count == 0 {
						break
					}
					sum += code[j]
					count--
				}
			}
			res[i] = sum
			sum, count = 0, k
		}
	}
	if k < 0 {
		for i := 0; i < len(code); i++ {
			for j := i - 1; j >= 0; j-- {
				if count == 0 {
					break
				}
				sum += code[j]
				count++
			}
			if count < 0 {
				for j := len(code) - 1; j >= 0; j-- {
					if count == 0 {
						break
					}
					sum += code[j]
					count++
				}
			}
			res[i] = sum
			sum, count = 0, k
		}
	}
	return res
}
