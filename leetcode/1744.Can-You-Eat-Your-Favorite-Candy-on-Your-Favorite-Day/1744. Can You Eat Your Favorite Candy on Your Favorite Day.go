package leetcode

func canEat(candiesCount []int, queries [][]int) []bool {
	n := len(candiesCount)
	prefixSum := make([]int, n)
	prefixSum[0] = candiesCount[0]
	for i := 1; i < n; i++ {
		prefixSum[i] = prefixSum[i-1] + candiesCount[i]
	}
	res := make([]bool, len(queries))
	for i, q := range queries {
		favoriteType, favoriteDay, dailyCap := q[0], q[1], q[2]
		x1 := favoriteDay + 1
		y1 := (favoriteDay + 1) * dailyCap
		x2 := 1
		if favoriteType > 0 {
			x2 = prefixSum[favoriteType-1] + 1
		}
		y2 := prefixSum[favoriteType]
		res[i] = !(x1 > y2 || y1 < x2)
	}
	return res
}
