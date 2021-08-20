package leetcode

func clumsy(N int) int {
	res, count, tmp, flag := 0, 1, N, false
	for i := N - 1; i > 0; i-- {
		count = count % 4
		switch count {
		case 1:
			tmp = tmp * i
		case 2:
			tmp = tmp / i
		case 3:
			res = res + tmp
			flag = true
			tmp = -1
			res = res + i
		case 0:
			flag = false
			tmp = tmp * (i)
		}
		count++
	}
	if !flag {
		res = res + tmp
	}
	return res
}
