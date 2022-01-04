package leetcocde

func numFriendRequests(ages []int) int {
	res, count := 0, [125]int{}
	for _, x := range ages {
		count[x]++
	}
	for i := 1; i <= 120; i++ {
		for j := 1; j <= 120; j++ {
			if j > i {
				continue
			}
			if (j-7)*2 <= i {
				continue
			}
			if j > 100 && i < 100 {
				continue
			}
			if i != j {
				res += count[i] * count[j]
			} else {
				res += count[i] * (count[j] - 1)
			}
		}
	}
	return res
}
