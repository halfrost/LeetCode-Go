package leetcode

func lastRemaining(n int) int {
	start, dir, cnt, step := 1, true, n, 1
	for cnt > 1 {
		if dir { // 正向
			start += step
		} else { // 反向
			if cnt%2 == 1 {
				start += step
			}
		}
		dir = !dir
		cnt >>= 1
		step <<= 1
	}
	return start
}
