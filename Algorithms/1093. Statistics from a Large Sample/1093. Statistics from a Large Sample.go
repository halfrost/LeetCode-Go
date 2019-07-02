package leetcode

func sampleStats(count []int) []float64 {
	res := make([]float64, 5)
	res[0] = 255
	sum := 0
	for _, val := range count {
		sum += val
	}
	left, right := sum/2, sum/2
	if (sum % 2) == 0 {
		right++
	}
	pre, mode := 0, 0
	for i, val := range count {
		if val > 0 {
			if i < int(res[0]) {
				res[0] = float64(i)
			}
			res[1] = float64(i)
		}
		res[2] += float64(i*val) / float64(sum)
		if pre < left && pre+val >= left {
			res[3] += float64(i) / 2.0
		}
		if pre < right && pre+val >= right {
			res[3] += float64(i) / 2.0
		}
		pre += val

		if val > mode {
			mode = val
			res[4] = float64(i)
		}
	}
	return res
}
