package leetcode

func removeStones(stones [][]int) int {
	if len(stones) <= 1 {
		return 0
	}
	uf, rowMap, colMap := UnionFind{}, map[int]int{}, map[int]int{}
	uf.init(len(stones))
	for i := 0; i < len(stones); i++ {
		if _, ok := rowMap[stones[i][0]]; ok {
			uf.union(rowMap[stones[i][0]], i)
		} else {
			rowMap[stones[i][0]] = i
		}
		if _, ok := colMap[stones[i][1]]; ok {
			uf.union(colMap[stones[i][1]], i)
		} else {
			colMap[stones[i][1]] = i
		}
	}
	return len(stones) - uf.totalCount()
}
