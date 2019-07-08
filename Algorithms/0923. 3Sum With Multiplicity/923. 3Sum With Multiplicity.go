package leetcode

import (
	"sort"
)

func threeSumMulti(A []int, target int) int {
	mod := 1000000007
	counter := map[int]int{}
	for _, value := range A {
		counter[value]++
	}

	uniqNums := []int{}
	for key := range counter {
		uniqNums = append(uniqNums, key)
	}
	sort.Ints(uniqNums)

	res := 0
	for i := 0; i < len(uniqNums); i++ {
		ni := counter[uniqNums[i]]
		if (uniqNums[i]*3 == target) && counter[uniqNums[i]] >= 3 {
			res += ni * (ni - 1) * (ni - 2) / 6
		}
		for j := i + 1; j < len(uniqNums); j++ {
			nj := counter[uniqNums[j]]
			if (uniqNums[i]*2+uniqNums[j] == target) && counter[uniqNums[i]] > 1 {
				res += ni * (ni - 1) / 2 * nj
			}
			if (uniqNums[j]*2+uniqNums[i] == target) && counter[uniqNums[j]] > 1 {
				res += nj * (nj - 1) / 2 * ni
			}
			c := target - uniqNums[i] - uniqNums[j]
			if c > uniqNums[j] && counter[c] > 0 {
				res += ni * nj * counter[c]
			}
		}
	}
	return res % mod
}
