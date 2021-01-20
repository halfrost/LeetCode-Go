package leetcode

func maximumWealth(accounts [][]int) int {
	res := 0
	for _, banks := range accounts {
		sAmount := 0
		for _, amount := range banks {
			sAmount += amount
		}
		if sAmount > res {
			res = sAmount
		}
	}
	return res
}
